package mocking

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprintf(w, "%d", 3)
}

func main() {
	Countdown(os.Stdout)
}
