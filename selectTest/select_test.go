package selectTest

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "www.google.com"
	fastURL := "www.baidu.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("Want '%v' but got '%v'", want, got)
	}
}
