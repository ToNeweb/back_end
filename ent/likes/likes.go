// Code generated by ent, DO NOT EDIT.

package likes

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the likes type in the database.
	Label = "likes"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the likes in the database.
	Table = "likes"
	// VideosTable is the table that holds the videos relation/edge. The primary key declared below.
	VideosTable = "videos_likeId"
	// VideosInverseTable is the table name for the Videos entity.
	// It exists in this package in order to avoid circular dependency with the "videos" package.
	VideosInverseTable = "videos"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_sec_likeId"
	// UserInverseTable is the table name for the UserSec entity.
	// It exists in this package in order to avoid circular dependency with the "usersec" package.
	UserInverseTable = "user_secs"
)

// Columns holds all SQL columns for likes fields.
var Columns = []string{
	FieldID,
}

var (
	// VideosPrimaryKey and VideosColumn2 are the table columns denoting the
	// primary key for the videos relation (M2M).
	VideosPrimaryKey = []string{"videos_id", "likes_id"}
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_sec_id", "likes_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Likes queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVideosCount orders the results by videos count.
func ByVideosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideosStep(), opts...)
	}
}

// ByVideos orders the results by videos terms.
func ByVideos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, VideosTable, VideosPrimaryKey...),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
