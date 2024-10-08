// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sekalahita/epirus/internal/ent/gen/googleauth"
	"github.com/sekalahita/epirus/internal/ent/gen/predicate"
	"github.com/sekalahita/epirus/internal/ent/gen/user"
)

// GoogleAuthQuery is the builder for querying GoogleAuth entities.
type GoogleAuthQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.GoogleAuth
	withUser   *UserQuery
	loadTotal  []func(context.Context, []*GoogleAuth) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoogleAuthQuery builder.
func (gaq *GoogleAuthQuery) Where(ps ...predicate.GoogleAuth) *GoogleAuthQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit the number of records to be returned by this query.
func (gaq *GoogleAuthQuery) Limit(limit int) *GoogleAuthQuery {
	gaq.ctx.Limit = &limit
	return gaq
}

// Offset to start from.
func (gaq *GoogleAuthQuery) Offset(offset int) *GoogleAuthQuery {
	gaq.ctx.Offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GoogleAuthQuery) Unique(unique bool) *GoogleAuthQuery {
	gaq.ctx.Unique = &unique
	return gaq
}

// Order specifies how the records should be ordered.
func (gaq *GoogleAuthQuery) Order(o ...OrderFunc) *GoogleAuthQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// QueryUser chains the current query on the "user" edge.
func (gaq *GoogleAuthQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(googleauth.Table, googleauth.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, googleauth.UserTable, googleauth.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GoogleAuth entity from the query.
// Returns a *NotFoundError when no GoogleAuth was found.
func (gaq *GoogleAuthQuery) First(ctx context.Context) (*GoogleAuth, error) {
	nodes, err := gaq.Limit(1).All(setContextOp(ctx, gaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{googleauth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GoogleAuthQuery) FirstX(ctx context.Context) *GoogleAuth {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoogleAuth ID from the query.
// Returns a *NotFoundError when no GoogleAuth ID was found.
func (gaq *GoogleAuthQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gaq.Limit(1).IDs(setContextOp(ctx, gaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{googleauth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GoogleAuthQuery) FirstIDX(ctx context.Context) string {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoogleAuth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoogleAuth entity is found.
// Returns a *NotFoundError when no GoogleAuth entities are found.
func (gaq *GoogleAuthQuery) Only(ctx context.Context) (*GoogleAuth, error) {
	nodes, err := gaq.Limit(2).All(setContextOp(ctx, gaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{googleauth.Label}
	default:
		return nil, &NotSingularError{googleauth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GoogleAuthQuery) OnlyX(ctx context.Context) *GoogleAuth {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoogleAuth ID in the query.
// Returns a *NotSingularError when more than one GoogleAuth ID is found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GoogleAuthQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gaq.Limit(2).IDs(setContextOp(ctx, gaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{googleauth.Label}
	default:
		err = &NotSingularError{googleauth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GoogleAuthQuery) OnlyIDX(ctx context.Context) string {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoogleAuths.
func (gaq *GoogleAuthQuery) All(ctx context.Context) ([]*GoogleAuth, error) {
	ctx = setContextOp(ctx, gaq.ctx, "All")
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GoogleAuth, *GoogleAuthQuery]()
	return withInterceptors[[]*GoogleAuth](ctx, gaq, qr, gaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GoogleAuthQuery) AllX(ctx context.Context) []*GoogleAuth {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoogleAuth IDs.
func (gaq *GoogleAuthQuery) IDs(ctx context.Context) (ids []string, err error) {
	if gaq.ctx.Unique == nil && gaq.path != nil {
		gaq.Unique(true)
	}
	ctx = setContextOp(ctx, gaq.ctx, "IDs")
	if err = gaq.Select(googleauth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GoogleAuthQuery) IDsX(ctx context.Context) []string {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gaq *GoogleAuthQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Count")
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gaq, querierCount[*GoogleAuthQuery](), gaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GoogleAuthQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GoogleAuthQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Exist")
	switch _, err := gaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GoogleAuthQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoogleAuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GoogleAuthQuery) Clone() *GoogleAuthQuery {
	if gaq == nil {
		return nil
	}
	return &GoogleAuthQuery{
		config:     gaq.config,
		ctx:        gaq.ctx.Clone(),
		order:      append([]OrderFunc{}, gaq.order...),
		inters:     append([]Interceptor{}, gaq.inters...),
		predicates: append([]predicate.GoogleAuth{}, gaq.predicates...),
		withUser:   gaq.withUser.Clone(),
		// clone intermediate query.
		sql:  gaq.sql.Clone(),
		path: gaq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GoogleAuthQuery) WithUser(opts ...func(*UserQuery)) *GoogleAuthQuery {
	query := (&UserClient{config: gaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gaq.withUser = query
	return gaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoogleAuth.Query().
//		GroupBy(googleauth.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (gaq *GoogleAuthQuery) GroupBy(field string, fields ...string) *GoogleAuthGroupBy {
	gaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoogleAuthGroupBy{build: gaq}
	grbuild.flds = &gaq.ctx.Fields
	grbuild.label = googleauth.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GoogleAuth.Query().
//		Select(googleauth.FieldCreatedAt).
//		Scan(ctx, &v)
func (gaq *GoogleAuthQuery) Select(fields ...string) *GoogleAuthSelect {
	gaq.ctx.Fields = append(gaq.ctx.Fields, fields...)
	sbuild := &GoogleAuthSelect{GoogleAuthQuery: gaq}
	sbuild.label = googleauth.Label
	sbuild.flds, sbuild.scan = &gaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoogleAuthSelect configured with the given aggregations.
func (gaq *GoogleAuthQuery) Aggregate(fns ...AggregateFunc) *GoogleAuthSelect {
	return gaq.Select().Aggregate(fns...)
}

func (gaq *GoogleAuthQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gaq.inters {
		if inter == nil {
			return fmt.Errorf("gen: uninitialized interceptor (forgotten import gen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gaq); err != nil {
				return err
			}
		}
	}
	for _, f := range gaq.ctx.Fields {
		if !googleauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}
		gaq.sql = prev
	}
	return nil
}

func (gaq *GoogleAuthQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoogleAuth, error) {
	var (
		nodes       = []*GoogleAuth{}
		_spec       = gaq.querySpec()
		loadedTypes = [1]bool{
			gaq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GoogleAuth).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GoogleAuth{config: gaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gaq.withUser; query != nil {
		if err := gaq.loadUser(ctx, query, nodes, nil,
			func(n *GoogleAuth, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range gaq.loadTotal {
		if err := gaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gaq *GoogleAuthQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GoogleAuth, init func(*GoogleAuth), assign func(*GoogleAuth, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*GoogleAuth)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gaq *GoogleAuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	_spec.Node.Columns = gaq.ctx.Fields
	if len(gaq.ctx.Fields) > 0 {
		_spec.Unique = gaq.ctx.Unique != nil && *gaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GoogleAuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(googleauth.Table, googleauth.Columns, sqlgraph.NewFieldSpec(googleauth.FieldID, field.TypeString))
	_spec.From = gaq.sql
	if unique := gaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gaq.path != nil {
		_spec.Unique = true
	}
	if fields := gaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, googleauth.FieldID)
		for i := range fields {
			if fields[i] != googleauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gaq *GoogleAuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(googleauth.Table)
	columns := gaq.ctx.Fields
	if len(columns) == 0 {
		columns = googleauth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gaq.ctx.Unique != nil && *gaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gaq.modifiers {
		m(selector)
	}
	for _, p := range gaq.predicates {
		p(selector)
	}
	for _, p := range gaq.order {
		p(selector)
	}
	if offset := gaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gaq *GoogleAuthQuery) Modify(modifiers ...func(s *sql.Selector)) *GoogleAuthSelect {
	gaq.modifiers = append(gaq.modifiers, modifiers...)
	return gaq.Select()
}

// GoogleAuthGroupBy is the group-by builder for GoogleAuth entities.
type GoogleAuthGroupBy struct {
	selector
	build *GoogleAuthQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GoogleAuthGroupBy) Aggregate(fns ...AggregateFunc) *GoogleAuthGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the selector query and scans the result into the given value.
func (gagb *GoogleAuthGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gagb.build.ctx, "GroupBy")
	if err := gagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoogleAuthQuery, *GoogleAuthGroupBy](ctx, gagb.build, gagb, gagb.build.inters, v)
}

func (gagb *GoogleAuthGroupBy) sqlScan(ctx context.Context, root *GoogleAuthQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gagb.fns))
	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gagb.flds)+len(gagb.fns))
		for _, f := range *gagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoogleAuthSelect is the builder for selecting fields of GoogleAuth entities.
type GoogleAuthSelect struct {
	*GoogleAuthQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gas *GoogleAuthSelect) Aggregate(fns ...AggregateFunc) *GoogleAuthSelect {
	gas.fns = append(gas.fns, fns...)
	return gas
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GoogleAuthSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gas.ctx, "Select")
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoogleAuthQuery, *GoogleAuthSelect](ctx, gas.GoogleAuthQuery, gas, gas.inters, v)
}

func (gas *GoogleAuthSelect) sqlScan(ctx context.Context, root *GoogleAuthQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gas.fns))
	for _, fn := range gas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gas *GoogleAuthSelect) Modify(modifiers ...func(s *sql.Selector)) *GoogleAuthSelect {
	gas.modifiers = append(gas.modifiers, modifiers...)
	return gas
}
