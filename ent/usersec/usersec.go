// Code generated by ent, DO NOT EDIT.

package usersec

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usersec type in the database.
	Label = "user_sec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeVideoId holds the string denoting the videoid edge name in mutations.
	EdgeVideoId = "videoId"
	// EdgeCommentId holds the string denoting the commentid edge name in mutations.
	EdgeCommentId = "commentId"
	// EdgeLikeId holds the string denoting the likeid edge name in mutations.
	EdgeLikeId = "likeId"
	// Table holds the table name of the usersec in the database.
	Table = "user_secs"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "user_secs"
	// ProfileInverseTable is the table name for the UserProfile entity.
	// It exists in this package in order to avoid circular dependency with the "userprofile" package.
	ProfileInverseTable = "user_profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "user_profile_user_secure"
	// VideoIdTable is the table that holds the videoId relation/edge. The primary key declared below.
	VideoIdTable = "user_sec_videoId"
	// VideoIdInverseTable is the table name for the Videos entity.
	// It exists in this package in order to avoid circular dependency with the "videos" package.
	VideoIdInverseTable = "videos"
	// CommentIdTable is the table that holds the commentId relation/edge. The primary key declared below.
	CommentIdTable = "user_sec_commentId"
	// CommentIdInverseTable is the table name for the Comments entity.
	// It exists in this package in order to avoid circular dependency with the "comments" package.
	CommentIdInverseTable = "comments"
	// LikeIdTable is the table that holds the likeId relation/edge. The primary key declared below.
	LikeIdTable = "user_sec_likeId"
	// LikeIdInverseTable is the table name for the Likes entity.
	// It exists in this package in order to avoid circular dependency with the "likes" package.
	LikeIdInverseTable = "likes"
)

// Columns holds all SQL columns for usersec fields.
var Columns = []string{
	FieldID,
	FieldPassword,
	FieldEmail,
	FieldAddress,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_secs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_profile_user_secure",
}

var (
	// VideoIdPrimaryKey and VideoIdColumn2 are the table columns denoting the
	// primary key for the videoId relation (M2M).
	VideoIdPrimaryKey = []string{"user_sec_id", "videos_id"}
	// CommentIdPrimaryKey and CommentIdColumn2 are the table columns denoting the
	// primary key for the commentId relation (M2M).
	CommentIdPrimaryKey = []string{"user_sec_id", "comments_id"}
	// LikeIdPrimaryKey and LikeIdColumn2 are the table columns denoting the
	// primary key for the likeId relation (M2M).
	LikeIdPrimaryKey = []string{"user_sec_id", "likes_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
)

// OrderOption defines the ordering options for the UserSec queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByVideoIdCount orders the results by videoId count.
func ByVideoIdCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoIdStep(), opts...)
	}
}

// ByVideoId orders the results by videoId terms.
func ByVideoId(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoIdStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentIdCount orders the results by commentId count.
func ByCommentIdCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentIdStep(), opts...)
	}
}

// ByCommentId orders the results by commentId terms.
func ByCommentId(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentIdStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikeIdCount orders the results by likeId count.
func ByLikeIdCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikeIdStep(), opts...)
	}
}

// ByLikeId orders the results by likeId terms.
func ByLikeId(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikeIdStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}
func newVideoIdStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoIdInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VideoIdTable, VideoIdPrimaryKey...),
	)
}
func newCommentIdStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentIdInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CommentIdTable, CommentIdPrimaryKey...),
	)
}
func newLikeIdStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikeIdInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LikeIdTable, LikeIdPrimaryKey...),
	)
}