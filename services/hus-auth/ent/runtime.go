// Code generated by ent, DO NOT EDIT.

package ent

import (
	"hus-auth/ent/hussession"
	"hus-auth/ent/schema"
	"hus-auth/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	hussessionFields := schema.HusSession{}.Fields()
	_ = hussessionFields
	// hussessionDescTid is the schema descriptor for tid field.
	hussessionDescTid := hussessionFields[1].Descriptor()
	// hussession.DefaultTid holds the default value on creation for the tid field.
	hussession.DefaultTid = hussessionDescTid.Default.(func() uuid.UUID)
	// hussessionDescIat is the schema descriptor for iat field.
	hussessionDescIat := hussessionFields[2].Descriptor()
	// hussession.DefaultIat holds the default value on creation for the iat field.
	hussession.DefaultIat = hussessionDescIat.Default.(func() time.Time)
	// hussessionDescPreserved is the schema descriptor for preserved field.
	hussessionDescPreserved := hussessionFields[3].Descriptor()
	// hussession.DefaultPreserved holds the default value on creation for the preserved field.
	hussession.DefaultPreserved = hussessionDescPreserved.Default.(bool)
	// hussessionDescUpdatedAt is the schema descriptor for updated_at field.
	hussessionDescUpdatedAt := hussessionFields[6].Descriptor()
	// hussession.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	hussession.DefaultUpdatedAt = hussessionDescUpdatedAt.Default.(func() time.Time)
	// hussession.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	hussession.UpdateDefaultUpdatedAt = hussessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// hussessionDescID is the schema descriptor for id field.
	hussessionDescID := hussessionFields[0].Descriptor()
	// hussession.DefaultID holds the default value on creation for the id field.
	hussession.DefaultID = hussessionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[10].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[11].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
