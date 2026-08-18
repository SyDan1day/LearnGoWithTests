package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
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
