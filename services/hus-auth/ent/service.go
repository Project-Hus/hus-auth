// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hus-auth/ent/service"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Protocol holds the value of the "protocol" field.
	Protocol string `json:"protocol,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceQuery when eager-loading is set.
	Edges ServiceEdges `json:"edges"`
}

// ServiceEdges holds the relations/edges for other nodes in the graph.
type ServiceEdges struct {
	// Subdomains holds the value of the subdomains edge.
	Subdomains []*Subdomain `json:"subdomains,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubdomainsOrErr returns the Subdomains value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) SubdomainsOrErr() ([]*Subdomain, error) {
	if e.loadedTypes[0] {
		return e.Subdomains, nil
	}
	return nil, &NotLoadedError{edge: "subdomains"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			values[i] = new(sql.NullInt64)
		case service.FieldName, service.FieldProtocol, service.FieldDomain:
			values[i] = new(sql.NullString)
		case service.FieldCreatedAt, service.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Service", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case service.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case service.FieldProtocol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field protocol", values[i])
			} else if value.Valid {
				s.Protocol = value.String
			}
		case service.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		case service.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case service.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QuerySubdomains queries the "subdomains" edge of the Service entity.
func (s *Service) QuerySubdomains() *SubdomainQuery {
	return NewServiceClient(s.config).QuerySubdomains(s)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("protocol=")
	builder.WriteString(s.Protocol)
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(s.Domain)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service