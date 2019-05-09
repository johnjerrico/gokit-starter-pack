package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type key int

const queryableKey key = 0

// NewContext ...
func NewContext(ctx context.Context, q Queryable) context.Context {
	return context.WithValue(ctx, queryableKey, q)
}

// QueryableFromContext ...
func QueryableFromContext(ctx context.Context) (Queryable, bool) {
	q, ok := ctx.Value(queryableKey).(Queryable)
	return q, ok
}

// Queryable ...
type Queryable interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// QueryableContext ...
type QueryableContext struct {
	q  Queryable
	db *sqlx.DB
	tx *sqlx.Tx
}

//NewQueryableContext ...
func NewQueryableContext(db *sqlx.DB) QueryableContext {
	return QueryableContext{
		q:  Queryable(db),
		db: db,
	}
}

//NewQueryableContextFromTx ...
func NewQueryableContextFromTx(tx *sqlx.Tx) QueryableContext {
	return QueryableContext{
		q:  Queryable(tx),
		tx: tx,
	}
}

//DB return db
func (qtx *QueryableContext) DB() *sqlx.DB {
	return qtx.db
}

//Tx return tx
func (qtx *QueryableContext) Tx() *sqlx.Tx {
	return qtx.tx
}

// BindNamed ...
func (qtx *QueryableContext) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return qtx.q.BindNamed(query, arg)
}

// DriverName ...
func (qtx *QueryableContext) DriverName() string {
	return qtx.q.DriverName()
}

// Get ...
func (qtx *QueryableContext) Get(dest interface{}, query string, args ...interface{}) error {
	return qtx.q.Get(dest, query, args...)
}

// GetContext ...
func (qtx *QueryableContext) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.GetContext(ctx, dest, query, args...)
	}
	return qtx.q.GetContext(ctx, dest, query, args...)
}

// MustExec ...
func (qtx *QueryableContext) MustExec(query string, args ...interface{}) sql.Result {
	return qtx.q.MustExec(query, args...)
}

// MustExecContext ...
func (qtx *QueryableContext) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.MustExecContext(ctx, query, args...)
	}
	return qtx.q.MustExecContext(ctx, query, args...)
}

// NamedExec ...
func (qtx *QueryableContext) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return qtx.q.NamedExec(query, arg)
}

// NamedExecContext ...
func (qtx *QueryableContext) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.NamedExecContext(ctx, query, arg)
	}
	return qtx.q.NamedExecContext(ctx, query, arg)
}

// NamedQuery ...
func (qtx *QueryableContext) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return qtx.q.NamedQuery(query, arg)
}

// PrepareNamed ...
func (qtx *QueryableContext) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return qtx.q.PrepareNamed(query)
}

// PrepareNamedContext ...
func (qtx *QueryableContext) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.PrepareNamedContext(ctx, query)
	}
	return qtx.q.PrepareNamedContext(ctx, query)
}

// Preparex ...
func (qtx *QueryableContext) Preparex(query string) (*sqlx.Stmt, error) {
	return qtx.q.Preparex(query)
}

// PreparexContext ...
func (qtx *QueryableContext) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.PreparexContext(ctx, query)
	}
	return qtx.q.PreparexContext(ctx, query)
}

// QueryRowx ...
func (qtx *QueryableContext) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return qtx.q.QueryRowx(query, args...)
}

// QueryRowxContext ...
func (qtx *QueryableContext) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.QueryRowxContext(ctx, query)
	}
	return qtx.q.QueryRowxContext(ctx, query)
}

// Queryx ...
func (qtx *QueryableContext) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return qtx.q.Queryx(query, args...)
}

// QueryxContext ...
func (qtx *QueryableContext) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		return queryable.QueryxContext(ctx, query, args...)
	}
	return qtx.q.QueryxContext(ctx, query, args...)
}

// Rebind ...
func (qtx *QueryableContext) Rebind(query string) string {
	return qtx.q.Rebind(query)
}

// Select ...
func (qtx *QueryableContext) Select(dest interface{}, query string, args ...interface{}) error {
	return qtx.q.Select(dest, query, args...)
}

// SelectContext ...
func (qtx *QueryableContext) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	queryable, ok := QueryableFromContext(ctx)
	if ok {
		queryable.SelectContext(ctx, dest, query, args...)
	}
	return qtx.q.SelectContext(ctx, dest, query)
}
