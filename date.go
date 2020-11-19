package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Date specifies a date without time (only year, month and day).
// Internally it represents a time.Time instant with zero UTC time parts.
// The zero value of the struct is represented as "0001-01-01".
type Date struct {
	t time.Time
}

// NewDate creates a new date instance from the provided year, month and day.
func NewDate(year int, month time.Month, day int) Date {
	return Date{
		t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

// Parse parses a formatted string and returns the date value it represents.
// The string must be in in ISO 8601 extended format (e.g. "2006-01-02").
func Parse(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)

	return Date{t: t}, err
}

// String returns a string representing the date instant in ISO 8601
// extended format (e.g. "2006-01-02").
func (d Date) String() string {
	return d.t.Format("2006-01-02")
}

// IsZero reports whether d represents the zero date instant,
// year 1, January 1 ("0001-01-01").
func (d Date) IsZero() bool {
	return d.t.IsZero()
}

// Day returns the day of the month specified by d.
func (d Date) Day() int {
	return d.t.Day()
}

// Month returns the month of the year specified by t.
func (d Date) Month() time.Month {
	return d.t.Month()
}

// Year returns the year in which d occurs.
func (d Date) Year() int {
	return d.t.Year()
}

// Before reports whether the date instant d is before d2.
func (d Date) Before(d2 Date) bool {
	return d.t.Before(d2.t)
}

// After reports whether the date instant d is after d2.
func (d Date) After(d2 Date) bool {
	return d.t.After(d2.t)
}

// Equal reports whether the date instant d is equal to d2.
func (d Date) Equal(d2 Date) bool {
	d1Year, d1Month, d1Day := d.t.Date()
	d2Year, d2Month, d2Day := d2.t.Date()

	return d1Year == d2Year && d1Month == d2Month && d1Day == d2Day
}

// Sub returns the duration d-d2 in days.
func (d Date) Sub(d2 Date) float64 {
	return d.t.Sub(d2.t).Hours() / 24
}

// Scan parses a value (usually from db). It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner.
func (d *Date) Scan(v interface{}) (err error) {
	err = nil

	switch v := v.(type) {
	case nil:
		d.t = time.Time{}
	case []byte:
		if len(v) > 0 {
			d.t, err = time.Parse("2006-01-02", string(v))
		} else {
			d.t = time.Time{}
		}
	case string:
		if v != "" {
			d.t, err = time.Parse("2006-01-02", v)
		} else {
			d.t = time.Time{}
		}
	case time.Time:
		year, month, day := v.Date()
		d.t = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful date", v, v)
		d.t = time.Time{}
	}

	return err
}

// Value converts the instant to a time.Time. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer.
func (d Date) Value() (driver.Value, error) {
	return d.t, nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// The date is given in ISO 8601 extended format (e.g. "2006-01-02").
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The date is expected to be in ISO 8601 extended format (e.g. "2006-01-02").
func (d *Date) UnmarshalText(data []byte) (err error) {
	return d.Scan(data)
}
