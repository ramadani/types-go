package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

// Date type
type Date struct {
	Time time.Time
}

// NullDate type
type NullDate struct {
	Time  time.Time
	Valid bool
}

func (d *Date) String() string {
	return d.Time.Format(dateFormat)
}

// Value date driver value
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan date scanner
func (d *Date) Scan(value interface{}) error {
	date := value.(string)
	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

// MarshalJSON date type
func (d Date) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf("\"%s\"", d.String())
	return []byte(res), nil
}

// UnmarshalJSON date type
func (d *Date) UnmarshalJSON(data []byte) error {
	var date string
	if err := json.Unmarshal(data, &date); err != nil {
		return err
	}

	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}
