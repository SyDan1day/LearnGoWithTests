package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "GO!"
)

type Sleeper interface {
	Sleep()
}

// 用于test，使得测试不依赖于time.Sleep，而是调用次数。
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// 用于业务，在time.Sleep上包装一层Sleeper接口，使得业务和测试解耦。
type ConfigurableSleeper struct {
	time time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.time)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	s := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, s)
}
