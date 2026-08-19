package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsitesChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

// 假的mockWebsitesChecker，阻塞20微秒.
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	got := CheckWebsites(mockWebsitesChecker, websites)

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if len(got) != len(want) {
		t.Fatalf("want '%v' but got '%v'", want, got)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want '%v' but got '%v'", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
