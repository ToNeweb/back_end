// Code generated by ent, DO NOT EDIT.

package ent

import (
	"server04/ent/schema"
	"server04/ent/usersec"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersecFields := schema.UserSec{}.Fields()
	_ = usersecFields
	// usersecDescAddress is the schema descriptor for address field.
	usersecDescAddress := usersecFields[2].Descriptor()
	// usersec.DefaultAddress holds the default value on creation for the address field.
	usersec.DefaultAddress = usersecDescAddress.Default.(string)
}
