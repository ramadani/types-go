package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Gender user gender
type Gender uint

// NullGender nullable gender
type NullGender struct {
	Gender Gender
	Valid  bool
}

const (
	// Unknown gender
	Unknown Gender = iota
	// Male gender
	Male
	// Female gender
	Female
)

func (e Gender) String() string {
	values := map[Gender]string{
		Male:   "M",
		Female: "F",
	}

	if val, ok := values[e]; ok {
		return val
	}

	return values[Unknown]
}

// Value gender valuer
func (e Gender) Value() (driver.Value, error) {
	return []byte(e.String()), nil
}

// Scan gender scanner
func (e *Gender) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		e.parseFrom(src.(string))
	case []byte:
		e.parseFrom(string(src.([]byte)))
	}

	return nil
}

// MarshalJSON marshalling gender
func (e *Gender) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf("\"%s\"", e.String())
	return []byte(res), nil
}

// UnmarshalJSON unmarshalling gender
func (e *Gender) UnmarshalJSON(data []byte) error {
	src, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	e.parseFrom(src)
	return nil
}

// ParseFrom parse gender from string
func (e *Gender) parseFrom(src string) {
	values := map[string]Gender{
		"":  Unknown,
		"M": Male,
		"F": Female,
	}

	if val, ok := values[src]; ok {
		*e = val
	} else {
		*e = values[""]
	}
}

// Value null gender value
func (e NullGender) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}

	return e.Gender.Value()
}

// Scan null gender scanner
func (e *NullGender) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if err := e.Gender.Scan(value); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON marshalling null gender
func (e *NullGender) MarshalJSON() ([]byte, error) {
	var res string
	if !e.Valid || e.Gender == Unknown {
		res = fmt.Sprintf("\"%s\"", "null")
	} else {
		res = fmt.Sprintf("\"%s\"", e.Gender.String())
	}

	return []byte(res), nil
}

// UnmarshalJSON unmarshalling gender
func (e *NullGender) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	src, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	e.Gender.parseFrom(src)
	e.Valid = true
	return nil
}
