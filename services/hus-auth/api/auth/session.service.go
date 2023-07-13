package auth

import (
	"hus-auth/common"
	"hus-auth/common/db"
	"hus-auth/common/dto"
	"hus-auth/common/helper"
	"hus-auth/common/hus"
	"hus-auth/common/service/session"
	"hus-auth/ent"
	"hus-auth/ent/connectedsession"
	"io/ioutil"

	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// HusSessionHandler godoc
// @Tags         auth
// @Router /hus [get]
// @Summary checks and issues the Hus session token
// @Description this endpoint can be used both for Cloudhus and subservices.
// @Description if the subservice redirects the client to this endpoint with service name, session id and redirect url, its session will be connected to Hus session.
// @Description and if fallback url is given, it will redirect to fallback url if it fails.
// @Description note that all urls must be url-encoded.
// @Param service query string true "subservice name"
// @Param redirect query string true "redirect url"
// @Param fallback query string false "fallback url"
// @Param sid query string true "subservice session id"
// @Success      303 "See Other, redirection"
// @Failure      303 "See Other, redirection"
func (ac authApiController) HusSessionHandler(c echo.Context) error {
	var err error

	// Query Parameters
	serviceName := c.QueryParam("service")  // name of the subservice that is requesting
	sessionID := c.QueryParam("sid")        // session ID of the subservice that is requesting
	redirectURL := c.QueryParam("redirect") // URL to be redirected after the request is processed
	fallbackURL := c.QueryParam("fallback") // URL to be redirected if the request fails
	if fallbackURL == "" {
		// if fallback URL is not given, it redirects to redirect URL
		fallbackURL = redirectURL
	}

	// if any of three parameters are not given, this request can't be handled.
	if serviceName == "" || sessionID == "" || redirectURL == "" {
		return c.Redirect(http.StatusSeeOther, common.Subservice["cloudhus"].Domain.URL+"/error")
	}

	// url decode
	redirectURL, err1 := url.QueryUnescape(redirectURL)
	fallbackURL, err2 := url.QueryUnescape(fallbackURL)
	if err1 != nil || err2 != nil {
		// invalid url
		return c.Redirect(http.StatusSeeOther, common.Subservice["cloudhus"].Domain.URL+"/error")
	}

	// service not registered, then halt.
	_, ok := common.Subservice[serviceName]
	if !ok {
		return c.Redirect(http.StatusSeeOther, fallbackURL)
	}
	// sessionID to UUID
	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, fallbackURL)
	}

	// get hus_st from cookie
	hus_st, err := c.Cookie("hus_st")
	if err != nil && err != http.ErrNoCookie {
		// there's an error while getting cookie, return error.
		return c.Redirect(http.StatusSeeOther, fallbackURL)
	}

	// validate Hus session
	hs, _, err := session.ValidateHusSession(c.Request().Context(), hus_st.Value)
	// if no valid Hus session found, establish new Hus session.
	// after redirection to this same endpoint, it will handle newly established Hus session.
	if err != nil {
		_, nhst, err := session.CreateHusSession(c.Request().Context())
		if err != nil {
			return c.Redirect(http.StatusSeeOther, fallbackURL)
		}

		nhstCookie := &http.Cookie{
			Name:     "hus_st",
			Value:    nhst,
			Path:     "/",
			Domain:   hus.AuthCookieDomain,
			HttpOnly: true,
			Secure:   hus.CookieSecure,
			SameSite: http.SameSiteLaxMode,
		}
		c.SetCookie(nhstCookie)

		// redirect to same endpoint here with same path and queries
		// this is to guarantee that the session is established between Cloudhus and the client
		tmpRedirect := common.Subservice["cloudhus"].Subdomains["auth"].URL +
			"/auth/hussession?service=" + serviceName +
			"&redirect=" + redirectURL +
			"&fallback=" + fallbackURL +
			"&sid=" + sessionID
		c.Redirect(http.StatusSeeOther, tmpRedirect)
	}

	// now handle the valid session
	err = session.ConnectSessions(c.Request().Context(), hs, serviceName, sessionUUID)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, fallbackURL)
	}

	// finally, rotate the Hus session
	nhst, err := session.RotateHusSession(c.Request().Context(), hs)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, fallbackURL)
	}

	nhstCookie := &http.Cookie{
		Name:     "hus_st",
		Value:    nhst,
		Path:     "/",
		Domain:   hus.AuthCookieDomain,
		HttpOnly: true,
		Secure:   hus.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	}
	if hs.Preserved {
		nhstCookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	}
	c.SetCookie(nhstCookie)

	return c.Redirect(http.StatusSeeOther, redirectURL)
}

// SessionConnectionHandler godoc
// @Tags         auth
// @Router /hus/connect/{token} [get]
// @Summary gets connection token from subservice and returns Hus session ID and user info
// @Description the token has properties pps, service and sid.
// @Param token path string true "pps, service name, session ID in signed token which expires only in 10 seconds"
// @Success      200 "Ok, session has been connected"
// @Failure      400 "Bad Request"
// @Failure      404 "Not Found, no such connected session"
func (ac authApiController) SessionConnectionHandler(c echo.Context) error {
	token := c.Param("token")
	claims, exp, err := helper.ParseJWTWithHMAC(token)
	if err != nil || exp {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	pps := claims["pps"].(string)
	if pps != "hus_connection" {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	service := claims["service"].(string)
	sid := claims["sid"].(string)
	suuid, err := uuid.Parse(sid)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid sid")
	}

	cs, err := db.Client.ConnectedSession.Query().Where(connectedsession.And(
		connectedsession.Service(service),
		connectedsession.Csid(suuid),
	)).WithHusSession(func(hsq *ent.HusSessionQuery) {
		hsq.WithUser()
	}).Only(c.Request().Context())
	if err != nil {
		return c.String(http.StatusNotFound, "no such session")
	}

	cu := cs.Edges.HusSession.Edges.User

	var hcu *dto.HusConnUser
	if cu != nil {
		hcu = &dto.HusConnUser{
			Uid:             cu.ID,
			ProfileImageURL: cu.ProfileImageURL,
			Email:           cu.Email,
			EmailVerified:   cu.EmailVerified,
			Name:            cu.Name,
			GivenName:       cu.GivenName,
			FamilyName:      cu.FamilyName,
		}
	}

	return c.JSON(http.StatusOK, dto.HusConnDto{
		Hsid: cs.Hsid.String(),
		User: hcu,
	})
}

// SignOutHandler godoc
// @Tags         auth
// @Router /hus/signout [patch]
// @Summary gets signout token from subservice and does signout process.
// @Description there are two types of signout process.
// @Description 1) sign out sessions related only to given hus session.
// @Description 2) sign out all related sessions to the user.
// @Param token path string true "sign out token"
// @Success      200 "Ok, session has been signed out"
// @Failure      400 "Bad Request"
// @Failure      500 "Internal Server Error"
func (ac authApiController) SignOutHandler(c echo.Context) error {
	// from request body get token string
	tokenBody := c.Request().Body
	defer tokenBody.Close()

	tokenByte, err := ioutil.ReadAll(tokenBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	token := string(tokenByte)

	claims, exp, err := helper.ParseJWTWithHMAC(token)
	if err != nil || exp {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	pps := claims["pps"].(string)
	if pps != "hus_signout" {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	hsid := claims["hsid"].(string)
	hsuuid, err := uuid.Parse(hsid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "parsing uuid failed")
	}

	soType := claims["type"].(string)
	switch soType {
	case "single":
		log.Println("not yet")
	case "total":
		err = session.SignOutTotal(c.Request().Context(), hsuuid)
		if err != nil {
			return c.String(http.StatusInternalServerError, "signout failed")
		}
	default:
		return c.String(http.StatusBadRequest, "invalid signout type")
	}

	return c.String(http.StatusOK, "signed out")
}
