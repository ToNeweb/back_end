// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server04/ent/comments"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comments is the model entity for the Comments schema.
type Comments struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CommentStr holds the value of the "commentStr" field.
	CommentStr string `json:"commentStr,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentsQuery when eager-loading is set.
	Edges        CommentsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentsEdges holds the relations/edges for other nodes in the graph.
type CommentsEdges struct {
	// VideoId holds the value of the videoId edge.
	VideoId []*Videos `json:"videoId,omitempty"`
	// UserId holds the value of the userId edge.
	UserId []*UserSec `json:"userId,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VideoIdOrErr returns the VideoId value or an error if the edge
// was not loaded in eager-loading.
func (e CommentsEdges) VideoIdOrErr() ([]*Videos, error) {
	if e.loadedTypes[0] {
		return e.VideoId, nil
	}
	return nil, &NotLoadedError{edge: "videoId"}
}

// UserIdOrErr returns the UserId value or an error if the edge
// was not loaded in eager-loading.
func (e CommentsEdges) UserIdOrErr() ([]*UserSec, error) {
	if e.loadedTypes[1] {
		return e.UserId, nil
	}
	return nil, &NotLoadedError{edge: "userId"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comments) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comments.FieldID:
			values[i] = new(sql.NullInt64)
		case comments.FieldCommentStr:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comments fields.
func (c *Comments) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comments.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comments.FieldCommentStr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commentStr", values[i])
			} else if value.Valid {
				c.CommentStr = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comments.
// This includes values selected through modifiers, order, etc.
func (c *Comments) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryVideoId queries the "videoId" edge of the Comments entity.
func (c *Comments) QueryVideoId() *VideosQuery {
	return NewCommentsClient(c.config).QueryVideoId(c)
}

// QueryUserId queries the "userId" edge of the Comments entity.
func (c *Comments) QueryUserId() *UserSecQuery {
	return NewCommentsClient(c.config).QueryUserId(c)
}

// Update returns a builder for updating this Comments.
// Note that you need to call Comments.Unwrap() before calling this method if this Comments
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comments) Update() *CommentsUpdateOne {
	return NewCommentsClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comments entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comments) Unwrap() *Comments {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comments is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comments) String() string {
	var builder strings.Builder
	builder.WriteString("Comments(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("commentStr=")
	builder.WriteString(c.CommentStr)
	builder.WriteByte(')')
	return builder.String()
}

// CommentsSlice is a parsable slice of Comments.
type CommentsSlice []*Comments
