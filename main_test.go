package lib_test

import (
	"testing"

	"github.com/drewwells/simple-golang"
)

func TestString(t *testing.T) {
	s := lib.String()
	if e := "Hello World!"; e != s {
		t.Errorf("Mismatch! got: %s wanted: %s", s, e)
	}
}

func TestNumber(t *testing.T) {
	n := lib.Number()
	if e := 1; e != n {
		t.Errorf("Mismatch! got: %d wanted: %d", n, e)
	}
}

func TestConverted(t *testing.T) {
	n := lib.Converted()
	if e := 2; e != n {
		t.Errorf("Mismatch! got: %d wanted: %d", n, e)
	}
}
