package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type Sleeper interface {
	Sleep()
}

type ConfigurablesSleeper struct {
	duration time.Duration
}

func (o *ConfigurablesSleeper) Sleep() {
	time.Sleep(o.duration)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(w, "Go!")

}

func main() {
	sleeper := &ConfigurablesSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
