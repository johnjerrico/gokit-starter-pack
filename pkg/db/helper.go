package db

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

//Order model for order by clause
type Order struct {
	By        string
	Direction string
}

//Assign create assign operation without type formatting
func Assign(field string, value string) string {
	var buffer strings.Builder
	buffer.WriteString(`"`)
	buffer.WriteString(field)
	buffer.WriteString(`"=`)
	buffer.WriteString(value)
	return buffer.String()
}

//AssignString create assign operation with string type value
func AssignString(field string, value string) string {
	var buffer strings.Builder
	buffer.WriteString(`"`)
	buffer.WriteString(field)
	buffer.WriteString(`"='`)
	buffer.WriteString(value)
	buffer.WriteString(`'`)
	return buffer.String()
}

//AssignInt create assign operation with int type value
func AssignInt(field string, value int) string {
	var buffer strings.Builder
	buffer.WriteString(`"`)
	buffer.WriteString(field)
	buffer.WriteString(`"=`)
	buffer.WriteString(fmt.Sprintf("%d", value))
	buffer.WriteString(` `)
	return buffer.String()
}

//Extract convert array of fields into string of fields
func Extract(fields []string, prefix string) string {
	var builder strings.Builder
	for idx, _field := range fields {
		if idx > 0 {
			builder.WriteString(",")
		}
		if prefix != "" {

			builder.WriteString(prefix)
			builder.WriteString(".")
		}
		builder.WriteString(`"`)
		builder.WriteString(_field)
		builder.WriteString(`"`)
	}
	return builder.String()
}

//AddPrefix add prefix to all items
func AddPrefix(fields []string, prefix string) []string {
	var _fields []string
	for _, _field := range fields {
		_fields = append(_fields, prefix+_field)
	}
	return _fields
}

//IRunInTransaction interface for calling runInTransactionWith
type IRunInTransaction interface {
	UpdateQueryable(queryable QueryableContext)
}

//RunInTransactionWithDB runs function with the transaction context without applying the context
func RunInTransactionWithDB(db *sqlx.DB, f func(tx *sqlx.Tx) error) error {

	_tx := db.MustBegin()

	_err := f(_tx)
	if _err != nil {
		_tx.Rollback()
		return _err
	}

	_err = _tx.Commit()
	if _err != nil {
		return fmt.Errorf("error when committing transaction: %v", _err)
	}
	return nil
}

//RunInTransaction runs function with transaction context and auto apply the context into existing IRunTransaction object
func RunInTransaction(db *sqlx.DB, obj IRunInTransaction, f func(tx *sqlx.Tx) error) error {

	_tx := db.MustBegin()
	obj.UpdateQueryable(NewQueryableContextFromTx(_tx))
	_err := f(_tx)
	if _err != nil {
		_tx.Rollback()
		return _err
	}
	_err = _tx.Commit()
	if _err != nil {
		return fmt.Errorf("error when committing transaction: %v", _err)
	}
	return nil
}
