package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
GO!`

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough call to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
