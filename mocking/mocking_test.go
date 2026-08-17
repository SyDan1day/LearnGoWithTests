package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}

	Countdown(&buffer)

	got := buffer.String()
	want := `3
			 2
			 1`

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
