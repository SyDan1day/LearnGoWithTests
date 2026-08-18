package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfiguableSleeper struct {
	time time.Duration
}

func (o *ConfiguableSleeper) Sleep() {
	time.Sleep(o.time)
}

type CountdownOperationSpy struct {
	Calls []string
}

func (o *CountdownOperationSpy) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

func (o *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	s.Sleep()
	fmt.Fprint(w, "GO!")
}

func main() {
	s := &ConfiguableSleeper{1 * time.Second}
	Countdown(os.Stdout, s)
}
