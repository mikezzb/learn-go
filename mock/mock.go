package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 0; i <= 3; i++ {
		sleeper.Sleep()
		if i != 3 {
			fmt.Fprintf(w, "%d\n", 3-i)
		} else {
			fmt.Fprintf(w, "Go!")
		}
	}

}

type CountdownSleeper struct {
	duration time.Duration
}

func (c *CountdownSleeper) Sleep() {
	time.Sleep(c.duration)
}

func main() {
	sleeper := &CountdownSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
