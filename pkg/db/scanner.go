package db

//Scanner helper for combinining sql.row and sql.rows
type Scanner interface {
	Scan(dest ...interface{}) error
}
