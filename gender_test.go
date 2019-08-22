package types

import (
	"fmt"
	"testing"
)

func TestGenderValuer(t *testing.T) {
	g := Male
	val, err := g.Value()
	if err != nil {
		t.Errorf("error on gender valuer: %v", err)
	}

	got := val.(string)
	want := "M"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestGenderScanner(t *testing.T) {
	g := new(Gender)
	err := g.Scan([]byte("M"))
	if err != nil {
		t.Errorf("error on gender scanner: %v", err)
	}

	got := *g
	want := Male

	if got != want {
		t.Errorf("got %s, want %s", got.String(), want.String())
	}
}

func TestGenderScannerWithUnknownValue(t *testing.T) {
	g := new(Gender)
	err := g.Scan([]byte("N"))
	if err != nil {
		t.Errorf("error on gender scanner: %v", err)
	}

	got := *g
	want := Unknown

	if got != want {
		t.Errorf("got %s, want %s", got.String(), want.String())
	}
}

func TestMarshalGender(t *testing.T) {
	g := Male
	b, err := g.MarshalJSON()
	if err != nil {
		t.Errorf("error on marshalling gender: %v", err)
	}

	got := string(b)
	want := fmt.Sprintf("\"%s\"", "M")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestUnmarshalGender(t *testing.T) {
	g := new(Gender)

	err := g.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "F")))
	if err != nil {
		t.Errorf("error on unmarshalling gender: %v", err)
	}

	got := *g
	want := Female

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestNullGenderValuer(t *testing.T) {
	g := NullGender{Gender: Male, Valid: true}
	val, err := g.Value()
	if err != nil {
		t.Errorf("error on null gender valuer: %v", err)
	}

	got := val.(string)
	want := "M"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestNullGenderValuerWithNil(t *testing.T) {
	g := NullGender{}
	val, err := g.Value()
	if err != nil {
		t.Errorf("error on null gender valuer: %v", err)
	}

	if val != nil {
		t.Errorf("gender should be nil")
	}
}

func TestNullGenderScanner(t *testing.T) {
	g := new(NullGender)
	err := g.Scan([]byte("M"))
	if err != nil {
		t.Errorf("error on null gender scanner: %v", err)
	}

	got := g.Gender
	want := Male

	if got != want {
		t.Errorf("got %s, want %s", got.String(), want.String())
	}
}

func TestNullGenderScannerWithNil(t *testing.T) {
	g := new(NullGender)
	err := g.Scan(nil)
	if err != nil {
		t.Errorf("error on null gender scanner: %v", err)
	}

	got := g.Gender
	want := Unknown

	if got != want {
		t.Errorf("got %s, want %s", got.String(), want.String())
	}
}

func TestMarshalNullGender(t *testing.T) {
	g := NullGender{Gender: Female, Valid: true}
	b, err := g.MarshalJSON()
	if err != nil {
		t.Errorf("error on marshalling null gender: %v", err)
	}

	got := string(b)
	want := fmt.Sprintf("\"%s\"", "F")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalNullGenderWithNil(t *testing.T) {
	g := NullGender{}
	b, err := g.MarshalJSON()
	if err != nil {
		t.Errorf("error on marshalling null gender: %v", err)
	}

	got := string(b)
	want := fmt.Sprintf("\"%s\"", "null")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestUnmarshalNullGender(t *testing.T) {
	g := new(NullGender)

	err := g.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "F")))
	if err != nil {
		t.Errorf("error on unmarshalling gender: %v", err)
	}

	got := g.Gender
	want := Female

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestUnmarshalNullGenderWithEmpty(t *testing.T) {
	g := new(NullGender)

	err := g.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", "")))
	if err != nil {
		t.Errorf("error on unmarshalling gender: %v", err)
	}

	got := g.Gender
	want := Unknown

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestUnmarshalNullGenderWithNil(t *testing.T) {
	g := new(NullGender)

	err := g.UnmarshalJSON(nil)
	if err != nil {
		t.Errorf("error on unmarshalling gender: %v", err)
	}

	got := g.Gender
	want := Unknown

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
