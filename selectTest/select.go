package selectTest

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) string {
	a := time.Now()
	http.Get(url1)
	aTime := time.Since(a)

	b := time.Now()
	http.Get(url2)
	bTime := time.Since(b)

	if aTime < bTime {
		return url1
	}

	return url2
}
