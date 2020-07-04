package dbcon

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// ErrExpectedOneRow is used for indicate that expected 1 row
var ErrExpectedOneRow = errors.New("expected 1 row")

func ExecAffectingOneRow(stmt *sql.Stmt, args ...interface{}) error {
	r, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return ErrExpectedOneRow
	}

	return nil
}

// TimeToNull returns nil if date has zero value
func TimeToNull(t time.Time) pq.NullTime {
	r := pq.NullTime{}
	r.Time = t

	if !t.IsZero() {
		r.Valid = true
	}

	return r
}

// ParseDateToTime returns nil if date has zero value
func ParseDateToTime(s string) pq.NullTime {
	format := "15:04:05"
	t, _ := time.Parse(format, s)

	return TimeToNull(t)
}

// IntToNull returns nil if date has zero value
func IntToNull(i int64) sql.NullInt64 {
	r := sql.NullInt64{}
	r.Int64 = i

	if i > 0 {
		r.Valid = true
	}

	return r
}

// StringToNull returns nil if date has zero value
func StringToNull(s string) sql.NullString {
	r := sql.NullString{}
	r.String = s

	if s != "" {
		r.Valid = true
	}

	return r
}

// FloatToNull returns nil if date has zero value
func FloatToNull(i float64) sql.NullFloat64 {
	r := sql.NullFloat64{}
	r.Float64 = i

	if i > 0 {
		r.Valid = true
	}

	return r
}
