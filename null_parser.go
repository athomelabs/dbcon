package dbcon

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

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
