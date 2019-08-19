package types

import (
	"fmt"
	"testing"
	"time"
)

func TestNullDateValuer(t *testing.T) {
	date := NullDate{Time: time.Date(2019, time.August, 17, 10, 15, 45, 0, time.Local), Valid: true}
	val, err := date.Value()
	if err != nil {
		t.Errorf("error on date valuer: %v", err)
	}

	got := val.(string)
	want := "2019-08-17"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestNullDateNilValuer(t *testing.T) {
	date := NullDate{}
	val, err := date.Value()
	if err != nil {
		t.Errorf("error on date valuer: %v", err)
	}

	if val != nil {
		t.Errorf("val should be nil")
	}
}

func TestNullDateScanner(t *testing.T) {
	date := NullDate{}
	err := date.Scan("2019-08-17")
	if err != nil {
		t.Errorf("error on date valuer: %v", err)
	}

	got := date.Time.Format("2006-01-02")
	want := "2019-08-17"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestNullDateNilScanner(t *testing.T) {
	date := NullDate{}
	err := date.Scan(nil)
	if err != nil {
		t.Errorf("error on date scanner: %v", err)
	}

	if !date.Time.IsZero() {
		t.Errorf("date time should be zero")
	}
}

func TestUnmarshalJSONNullDate(t *testing.T) {
	date := NullDate{}
	err := date.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "2019-08-08")))
	if err != nil {
		t.Errorf("error on unmarshal json date: %v", err)
	}

	if date.Time.IsZero() {
		t.Errorf("time should not zero")
	}

	if date.String() != "2019-08-08" {
		t.Errorf("got %s want %s", date.String(), "2019-08-08")
	}
}

func TestUnmarshalJSONWrongFormatNullDate(t *testing.T) {
	date := NullDate{}
	err := date.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "08-08-2019")))
	if err != nil {
		t.Errorf("error on unmarshal json date: %v", err)
	}

	if !date.Time.IsZero() {
		t.Errorf("time should zero")
	}
}

func TestUnmarshalJSONEmptyNullDate(t *testing.T) {
	date := NullDate{}
	err := date.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "")))
	if err != nil {
		t.Errorf("error on unmarshal json date: %v", err)
	}

	if !date.Time.IsZero() {
		t.Errorf("time should zero")
	}

	if date.String() != "" {
		t.Errorf("date should be empty")
	}
}

func TestMarshalJSONNullDate(t *testing.T) {
	date := NullDate{Time: time.Date(2019, time.August, 17, 10, 15, 45, 0, time.Local), Valid: true}
	json, err := date.MarshalJSON()
	if err != nil {
		t.Errorf("error on marshal json date: %v", err)
	}

	got := string(json)
	want := fmt.Sprintf("\"%s\"", "2019-08-17")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	if date.String() != "2019-08-17" {
		t.Errorf("got %s want %s", date.String(), "2019-08-17")
	}
}

func TestMarshalJSONValidFalseNullDate(t *testing.T) {
	date := NullDate{}
	json, err := date.MarshalJSON()
	if err != nil {
		t.Errorf("error on marshal json date: %v", err)
	}

	got := string(json)
	want := fmt.Sprintf("\"%s\"", "null")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	if date.String() != "" {
		t.Errorf("date should be empty")
	}
}
