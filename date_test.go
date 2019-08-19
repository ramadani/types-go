package types

import (
	"fmt"
	"testing"
	"time"
)

func TestDateValuer(t *testing.T) {
	date := Date{Time: time.Date(2019, time.August, 17, 10, 15, 45, 0, time.Local)}
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

func TestDateScanner(t *testing.T) {
	date := Date{}
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

func TestDateScannerWrongFormat(t *testing.T) {
	date := Date{}
	err := date.Scan("08-17-2019")
	if err == nil {
		t.Errorf("date scanner should be err")
	}

	if !date.Time.IsZero() {
		t.Errorf("time should be zero")
	}
}

func TestUnmarshalJSONDate(t *testing.T) {
	date := Date{}
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

func TestUnmarshalJSONWrongFormatDate(t *testing.T) {
	date := Date{}
	err := date.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "08-08-2019")))
	if err == nil {
		t.Errorf("should error occurs")
	}
}

func TestMarshalJSONDate(t *testing.T) {
	date := Date{Time: time.Date(2019, time.August, 17, 10, 15, 45, 0, time.Local)}
	json, err := date.MarshalJSON()
	if err != nil {
		t.Errorf("error on unmarshal json date: %v", err)
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
