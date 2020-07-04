package dbcon

type RowScanner interface {
	Scan(dest ...interface{}) error
}
