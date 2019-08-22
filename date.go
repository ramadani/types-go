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

func (d *NullDate) String() string {
	if !d.Valid {
		return ""
	}

	return d.Time.Format(dateFormat)
}

// Value null date driver value
func (d NullDate) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return d.String(), nil
}

// Scan null date scanner
func (d *NullDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	date := value.(string)
	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

// MarshalJSON null date type
func (d NullDate) MarshalJSON() ([]byte, error) {
	date := "null"
	if d.Valid {
		date = d.String()
	}

	res := fmt.Sprintf("\"%s\"", date)
	return []byte(res), nil
}

// UnmarshalJSON date type
func (d *NullDate) UnmarshalJSON(data []byte) error {
	var date string
	if err := json.Unmarshal(data, &date); err != nil {
		return nil
	}

	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return nil
	}

	d.Valid = true
	d.Time = t
	return nil
}
