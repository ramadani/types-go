package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

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
