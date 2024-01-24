// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server04/ent/likes"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Likes is the model entity for the Likes schema.
type Likes struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LikesQuery when eager-loading is set.
	Edges        LikesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LikesEdges holds the relations/edges for other nodes in the graph.
type LikesEdges struct {
	// Videos holds the value of the videos edge.
	Videos []*Videos `json:"videos,omitempty"`
	// User holds the value of the user edge.
	User []*UserSec `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VideosOrErr returns the Videos value or an error if the edge
// was not loaded in eager-loading.
func (e LikesEdges) VideosOrErr() ([]*Videos, error) {
	if e.loadedTypes[0] {
		return e.Videos, nil
	}
	return nil, &NotLoadedError{edge: "videos"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e LikesEdges) UserOrErr() ([]*UserSec, error) {
	if e.loadedTypes[1] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Likes) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case likes.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Likes fields.
func (l *Likes) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case likes.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Likes.
// This includes values selected through modifiers, order, etc.
func (l *Likes) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryVideos queries the "videos" edge of the Likes entity.
func (l *Likes) QueryVideos() *VideosQuery {
	return NewLikesClient(l.config).QueryVideos(l)
}

// QueryUser queries the "user" edge of the Likes entity.
func (l *Likes) QueryUser() *UserSecQuery {
	return NewLikesClient(l.config).QueryUser(l)
}

// Update returns a builder for updating this Likes.
// Note that you need to call Likes.Unwrap() before calling this method if this Likes
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Likes) Update() *LikesUpdateOne {
	return NewLikesClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Likes entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Likes) Unwrap() *Likes {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Likes is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Likes) String() string {
	var builder strings.Builder
	builder.WriteString("Likes(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteByte(')')
	return builder.String()
}

// LikesSlice is a parsable slice of Likes.
type LikesSlice []*Likes
