// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"server04/ent/migrate"

	"server04/ent/comments"
	"server04/ent/likes"
	"server04/ent/userprofile"
	"server04/ent/usersec"
	"server04/ent/videos"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Comments is the client for interacting with the Comments builders.
	Comments *CommentsClient
	// Likes is the client for interacting with the Likes builders.
	Likes *LikesClient
	// UserProfile is the client for interacting with the UserProfile builders.
	UserProfile *UserProfileClient
	// UserSec is the client for interacting with the UserSec builders.
	UserSec *UserSecClient
	// Videos is the client for interacting with the Videos builders.
	Videos *VideosClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Comments = NewCommentsClient(c.config)
	c.Likes = NewLikesClient(c.config)
	c.UserProfile = NewUserProfileClient(c.config)
	c.UserSec = NewUserSecClient(c.config)
	c.Videos = NewVideosClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Comments:    NewCommentsClient(cfg),
		Likes:       NewLikesClient(cfg),
		UserProfile: NewUserProfileClient(cfg),
		UserSec:     NewUserSecClient(cfg),
		Videos:      NewVideosClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Comments:    NewCommentsClient(cfg),
		Likes:       NewLikesClient(cfg),
		UserProfile: NewUserProfileClient(cfg),
		UserSec:     NewUserSecClient(cfg),
		Videos:      NewVideosClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Comments.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Comments.Use(hooks...)
	c.Likes.Use(hooks...)
	c.UserProfile.Use(hooks...)
	c.UserSec.Use(hooks...)
	c.Videos.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Comments.Intercept(interceptors...)
	c.Likes.Intercept(interceptors...)
	c.UserProfile.Intercept(interceptors...)
	c.UserSec.Intercept(interceptors...)
	c.Videos.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CommentsMutation:
		return c.Comments.mutate(ctx, m)
	case *LikesMutation:
		return c.Likes.mutate(ctx, m)
	case *UserProfileMutation:
		return c.UserProfile.mutate(ctx, m)
	case *UserSecMutation:
		return c.UserSec.mutate(ctx, m)
	case *VideosMutation:
		return c.Videos.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CommentsClient is a client for the Comments schema.
type CommentsClient struct {
	config
}

// NewCommentsClient returns a client for the Comments from the given config.
func NewCommentsClient(c config) *CommentsClient {
	return &CommentsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comments.Hooks(f(g(h())))`.
func (c *CommentsClient) Use(hooks ...Hook) {
	c.hooks.Comments = append(c.hooks.Comments, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `comments.Intercept(f(g(h())))`.
func (c *CommentsClient) Intercept(interceptors ...Interceptor) {
	c.inters.Comments = append(c.inters.Comments, interceptors...)
}

// Create returns a builder for creating a Comments entity.
func (c *CommentsClient) Create() *CommentsCreate {
	mutation := newCommentsMutation(c.config, OpCreate)
	return &CommentsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comments entities.
func (c *CommentsClient) CreateBulk(builders ...*CommentsCreate) *CommentsCreateBulk {
	return &CommentsCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CommentsClient) MapCreateBulk(slice any, setFunc func(*CommentsCreate, int)) *CommentsCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CommentsCreateBulk{err: fmt.Errorf("calling to CommentsClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CommentsCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CommentsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comments.
func (c *CommentsClient) Update() *CommentsUpdate {
	mutation := newCommentsMutation(c.config, OpUpdate)
	return &CommentsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentsClient) UpdateOne(co *Comments) *CommentsUpdateOne {
	mutation := newCommentsMutation(c.config, OpUpdateOne, withComments(co))
	return &CommentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentsClient) UpdateOneID(id int) *CommentsUpdateOne {
	mutation := newCommentsMutation(c.config, OpUpdateOne, withCommentsID(id))
	return &CommentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comments.
func (c *CommentsClient) Delete() *CommentsDelete {
	mutation := newCommentsMutation(c.config, OpDelete)
	return &CommentsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentsClient) DeleteOne(co *Comments) *CommentsDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentsClient) DeleteOneID(id int) *CommentsDeleteOne {
	builder := c.Delete().Where(comments.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentsDeleteOne{builder}
}

// Query returns a query builder for Comments.
func (c *CommentsClient) Query() *CommentsQuery {
	return &CommentsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeComments},
		inters: c.Interceptors(),
	}
}

// Get returns a Comments entity by its id.
func (c *CommentsClient) Get(ctx context.Context, id int) (*Comments, error) {
	return c.Query().Where(comments.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentsClient) GetX(ctx context.Context, id int) *Comments {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVideoId queries the videoId edge of a Comments.
func (c *CommentsClient) QueryVideoId(co *Comments) *VideosQuery {
	query := (&VideosClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comments.Table, comments.FieldID, id),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comments.VideoIdTable, comments.VideoIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserId queries the userId edge of a Comments.
func (c *CommentsClient) QueryUserId(co *Comments) *UserSecQuery {
	query := (&UserSecClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comments.Table, comments.FieldID, id),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comments.UserIdTable, comments.UserIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentsClient) Hooks() []Hook {
	return c.hooks.Comments
}

// Interceptors returns the client interceptors.
func (c *CommentsClient) Interceptors() []Interceptor {
	return c.inters.Comments
}

func (c *CommentsClient) mutate(ctx context.Context, m *CommentsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommentsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommentsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommentsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Comments mutation op: %q", m.Op())
	}
}

// LikesClient is a client for the Likes schema.
type LikesClient struct {
	config
}

// NewLikesClient returns a client for the Likes from the given config.
func NewLikesClient(c config) *LikesClient {
	return &LikesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `likes.Hooks(f(g(h())))`.
func (c *LikesClient) Use(hooks ...Hook) {
	c.hooks.Likes = append(c.hooks.Likes, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `likes.Intercept(f(g(h())))`.
func (c *LikesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Likes = append(c.inters.Likes, interceptors...)
}

// Create returns a builder for creating a Likes entity.
func (c *LikesClient) Create() *LikesCreate {
	mutation := newLikesMutation(c.config, OpCreate)
	return &LikesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Likes entities.
func (c *LikesClient) CreateBulk(builders ...*LikesCreate) *LikesCreateBulk {
	return &LikesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LikesClient) MapCreateBulk(slice any, setFunc func(*LikesCreate, int)) *LikesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LikesCreateBulk{err: fmt.Errorf("calling to LikesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LikesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LikesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Likes.
func (c *LikesClient) Update() *LikesUpdate {
	mutation := newLikesMutation(c.config, OpUpdate)
	return &LikesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LikesClient) UpdateOne(l *Likes) *LikesUpdateOne {
	mutation := newLikesMutation(c.config, OpUpdateOne, withLikes(l))
	return &LikesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LikesClient) UpdateOneID(id int) *LikesUpdateOne {
	mutation := newLikesMutation(c.config, OpUpdateOne, withLikesID(id))
	return &LikesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Likes.
func (c *LikesClient) Delete() *LikesDelete {
	mutation := newLikesMutation(c.config, OpDelete)
	return &LikesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LikesClient) DeleteOne(l *Likes) *LikesDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LikesClient) DeleteOneID(id int) *LikesDeleteOne {
	builder := c.Delete().Where(likes.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LikesDeleteOne{builder}
}

// Query returns a query builder for Likes.
func (c *LikesClient) Query() *LikesQuery {
	return &LikesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLikes},
		inters: c.Interceptors(),
	}
}

// Get returns a Likes entity by its id.
func (c *LikesClient) Get(ctx context.Context, id int) (*Likes, error) {
	return c.Query().Where(likes.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LikesClient) GetX(ctx context.Context, id int) *Likes {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVideos queries the videos edge of a Likes.
func (c *LikesClient) QueryVideos(l *Likes) *VideosQuery {
	query := (&VideosClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likes.Table, likes.FieldID, id),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, likes.VideosTable, likes.VideosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Likes.
func (c *LikesClient) QueryUser(l *Likes) *UserSecQuery {
	query := (&UserSecClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likes.Table, likes.FieldID, id),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, likes.UserTable, likes.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LikesClient) Hooks() []Hook {
	return c.hooks.Likes
}

// Interceptors returns the client interceptors.
func (c *LikesClient) Interceptors() []Interceptor {
	return c.inters.Likes
}

func (c *LikesClient) mutate(ctx context.Context, m *LikesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LikesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LikesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LikesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LikesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Likes mutation op: %q", m.Op())
	}
}

// UserProfileClient is a client for the UserProfile schema.
type UserProfileClient struct {
	config
}

// NewUserProfileClient returns a client for the UserProfile from the given config.
func NewUserProfileClient(c config) *UserProfileClient {
	return &UserProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userprofile.Hooks(f(g(h())))`.
func (c *UserProfileClient) Use(hooks ...Hook) {
	c.hooks.UserProfile = append(c.hooks.UserProfile, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `userprofile.Intercept(f(g(h())))`.
func (c *UserProfileClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserProfile = append(c.inters.UserProfile, interceptors...)
}

// Create returns a builder for creating a UserProfile entity.
func (c *UserProfileClient) Create() *UserProfileCreate {
	mutation := newUserProfileMutation(c.config, OpCreate)
	return &UserProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserProfile entities.
func (c *UserProfileClient) CreateBulk(builders ...*UserProfileCreate) *UserProfileCreateBulk {
	return &UserProfileCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserProfileClient) MapCreateBulk(slice any, setFunc func(*UserProfileCreate, int)) *UserProfileCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserProfileCreateBulk{err: fmt.Errorf("calling to UserProfileClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserProfileCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserProfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserProfile.
func (c *UserProfileClient) Update() *UserProfileUpdate {
	mutation := newUserProfileMutation(c.config, OpUpdate)
	return &UserProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserProfileClient) UpdateOne(up *UserProfile) *UserProfileUpdateOne {
	mutation := newUserProfileMutation(c.config, OpUpdateOne, withUserProfile(up))
	return &UserProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserProfileClient) UpdateOneID(id int) *UserProfileUpdateOne {
	mutation := newUserProfileMutation(c.config, OpUpdateOne, withUserProfileID(id))
	return &UserProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserProfile.
func (c *UserProfileClient) Delete() *UserProfileDelete {
	mutation := newUserProfileMutation(c.config, OpDelete)
	return &UserProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserProfileClient) DeleteOne(up *UserProfile) *UserProfileDeleteOne {
	return c.DeleteOneID(up.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserProfileClient) DeleteOneID(id int) *UserProfileDeleteOne {
	builder := c.Delete().Where(userprofile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserProfileDeleteOne{builder}
}

// Query returns a query builder for UserProfile.
func (c *UserProfileClient) Query() *UserProfileQuery {
	return &UserProfileQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserProfile},
		inters: c.Interceptors(),
	}
}

// Get returns a UserProfile entity by its id.
func (c *UserProfileClient) Get(ctx context.Context, id int) (*UserProfile, error) {
	return c.Query().Where(userprofile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserProfileClient) GetX(ctx context.Context, id int) *UserProfile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserSecure queries the UserSecure edge of a UserProfile.
func (c *UserProfileClient) QueryUserSecure(up *UserProfile) *UserSecQuery {
	query := (&UserSecClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := up.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userprofile.Table, userprofile.FieldID, id),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userprofile.UserSecureTable, userprofile.UserSecureColumn),
		)
		fromV = sqlgraph.Neighbors(up.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserProfileClient) Hooks() []Hook {
	return c.hooks.UserProfile
}

// Interceptors returns the client interceptors.
func (c *UserProfileClient) Interceptors() []Interceptor {
	return c.inters.UserProfile
}

func (c *UserProfileClient) mutate(ctx context.Context, m *UserProfileMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserProfileCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserProfileDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserProfile mutation op: %q", m.Op())
	}
}

// UserSecClient is a client for the UserSec schema.
type UserSecClient struct {
	config
}

// NewUserSecClient returns a client for the UserSec from the given config.
func NewUserSecClient(c config) *UserSecClient {
	return &UserSecClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usersec.Hooks(f(g(h())))`.
func (c *UserSecClient) Use(hooks ...Hook) {
	c.hooks.UserSec = append(c.hooks.UserSec, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `usersec.Intercept(f(g(h())))`.
func (c *UserSecClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserSec = append(c.inters.UserSec, interceptors...)
}

// Create returns a builder for creating a UserSec entity.
func (c *UserSecClient) Create() *UserSecCreate {
	mutation := newUserSecMutation(c.config, OpCreate)
	return &UserSecCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserSec entities.
func (c *UserSecClient) CreateBulk(builders ...*UserSecCreate) *UserSecCreateBulk {
	return &UserSecCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserSecClient) MapCreateBulk(slice any, setFunc func(*UserSecCreate, int)) *UserSecCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserSecCreateBulk{err: fmt.Errorf("calling to UserSecClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserSecCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserSecCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserSec.
func (c *UserSecClient) Update() *UserSecUpdate {
	mutation := newUserSecMutation(c.config, OpUpdate)
	return &UserSecUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserSecClient) UpdateOne(us *UserSec) *UserSecUpdateOne {
	mutation := newUserSecMutation(c.config, OpUpdateOne, withUserSec(us))
	return &UserSecUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserSecClient) UpdateOneID(id int) *UserSecUpdateOne {
	mutation := newUserSecMutation(c.config, OpUpdateOne, withUserSecID(id))
	return &UserSecUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserSec.
func (c *UserSecClient) Delete() *UserSecDelete {
	mutation := newUserSecMutation(c.config, OpDelete)
	return &UserSecDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserSecClient) DeleteOne(us *UserSec) *UserSecDeleteOne {
	return c.DeleteOneID(us.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserSecClient) DeleteOneID(id int) *UserSecDeleteOne {
	builder := c.Delete().Where(usersec.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserSecDeleteOne{builder}
}

// Query returns a query builder for UserSec.
func (c *UserSecClient) Query() *UserSecQuery {
	return &UserSecQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserSec},
		inters: c.Interceptors(),
	}
}

// Get returns a UserSec entity by its id.
func (c *UserSecClient) Get(ctx context.Context, id int) (*UserSec, error) {
	return c.Query().Where(usersec.ID(id)).Only(ctx)
}

///sep
func (c *UserSecClient) GetByEmail(ctx context.Context, email string) (*UserSec, error) {
	return c.Query().Where(usersec.Email(email)).Only(ctx)
}


// GetX is like Get, but panics if an error occurs.
func (c *UserSecClient) GetX(ctx context.Context, id int) *UserSec {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProfile queries the profile edge of a UserSec.
func (c *UserSecClient) QueryProfile(us *UserSec) *UserProfileQuery {
	query := (&UserProfileClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := us.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersec.Table, usersec.FieldID, id),
			sqlgraph.To(userprofile.Table, userprofile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usersec.ProfileTable, usersec.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(us.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVideoId queries the videoId edge of a UserSec.
func (c *UserSecClient) QueryVideoId(us *UserSec) *VideosQuery {
	query := (&VideosClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := us.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersec.Table, usersec.FieldID, id),
			sqlgraph.To(videos.Table, videos.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usersec.VideoIdTable, usersec.VideoIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(us.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCommentId queries the commentId edge of a UserSec.
func (c *UserSecClient) QueryCommentId(us *UserSec) *CommentsQuery {
	query := (&CommentsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := us.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersec.Table, usersec.FieldID, id),
			sqlgraph.To(comments.Table, comments.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usersec.CommentIdTable, usersec.CommentIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(us.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikeId queries the likeId edge of a UserSec.
func (c *UserSecClient) QueryLikeId(us *UserSec) *LikesQuery {
	query := (&LikesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := us.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersec.Table, usersec.FieldID, id),
			sqlgraph.To(likes.Table, likes.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usersec.LikeIdTable, usersec.LikeIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(us.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserSecClient) Hooks() []Hook {
	return c.hooks.UserSec
}

// Interceptors returns the client interceptors.
func (c *UserSecClient) Interceptors() []Interceptor {
	return c.inters.UserSec
}

func (c *UserSecClient) mutate(ctx context.Context, m *UserSecMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserSecCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserSecUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserSecUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserSecDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserSec mutation op: %q", m.Op())
	}
}

// VideosClient is a client for the Videos schema.
type VideosClient struct {
	config
}

// NewVideosClient returns a client for the Videos from the given config.
func NewVideosClient(c config) *VideosClient {
	return &VideosClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `videos.Hooks(f(g(h())))`.
func (c *VideosClient) Use(hooks ...Hook) {
	c.hooks.Videos = append(c.hooks.Videos, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `videos.Intercept(f(g(h())))`.
func (c *VideosClient) Intercept(interceptors ...Interceptor) {
	c.inters.Videos = append(c.inters.Videos, interceptors...)
}

// Create returns a builder for creating a Videos entity.
func (c *VideosClient) Create() *VideosCreate {
	mutation := newVideosMutation(c.config, OpCreate)
	return &VideosCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Videos entities.
func (c *VideosClient) CreateBulk(builders ...*VideosCreate) *VideosCreateBulk {
	return &VideosCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *VideosClient) MapCreateBulk(slice any, setFunc func(*VideosCreate, int)) *VideosCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &VideosCreateBulk{err: fmt.Errorf("calling to VideosClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*VideosCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &VideosCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Videos.
func (c *VideosClient) Update() *VideosUpdate {
	mutation := newVideosMutation(c.config, OpUpdate)
	return &VideosUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideosClient) UpdateOne(v *Videos) *VideosUpdateOne {
	mutation := newVideosMutation(c.config, OpUpdateOne, withVideos(v))
	return &VideosUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideosClient) UpdateOneID(id int) *VideosUpdateOne {
	mutation := newVideosMutation(c.config, OpUpdateOne, withVideosID(id))
	return &VideosUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Videos.
func (c *VideosClient) Delete() *VideosDelete {
	mutation := newVideosMutation(c.config, OpDelete)
	return &VideosDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VideosClient) DeleteOne(v *Videos) *VideosDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VideosClient) DeleteOneID(id int) *VideosDeleteOne {
	builder := c.Delete().Where(videos.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideosDeleteOne{builder}
}

// Query returns a query builder for Videos.
func (c *VideosClient) Query() *VideosQuery {
	return &VideosQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVideos},
		inters: c.Interceptors(),
	}
}

// Get returns a Videos entity by its id.
func (c *VideosClient) Get(ctx context.Context, id int) (*Videos, error) {
	return c.Query().Where(videos.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideosClient) GetX(ctx context.Context, id int) *Videos {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Videos.
func (c *VideosClient) QueryUser(v *Videos) *UserSecQuery {
	query := (&UserSecClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, id),
			sqlgraph.To(usersec.Table, usersec.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, videos.UserTable, videos.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikeId queries the likeId edge of a Videos.
func (c *VideosClient) QueryLikeId(v *Videos) *LikesQuery {
	query := (&LikesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, id),
			sqlgraph.To(likes.Table, likes.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, videos.LikeIdTable, videos.LikeIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCommentId queries the commentId edge of a Videos.
func (c *VideosClient) QueryCommentId(v *Videos) *CommentsQuery {
	query := (&CommentsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(videos.Table, videos.FieldID, id),
			sqlgraph.To(comments.Table, comments.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, videos.CommentIdTable, videos.CommentIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideosClient) Hooks() []Hook {
	return c.hooks.Videos
}

// Interceptors returns the client interceptors.
func (c *VideosClient) Interceptors() []Interceptor {
	return c.inters.Videos
}

func (c *VideosClient) mutate(ctx context.Context, m *VideosMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VideosCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VideosUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VideosUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VideosDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Videos mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Comments, Likes, UserProfile, UserSec, Videos []ent.Hook
	}
	inters struct {
		Comments, Likes, UserProfile, UserSec, Videos []ent.Interceptor
	}
)