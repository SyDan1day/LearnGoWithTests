package mocking

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
}

func main() {
	Countdown(os.Stdout)
}
