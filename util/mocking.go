package util

import (
	"fmt"
	"io"
)

const FINAL_WORD = "Go!"
const COUNTDOWN_START = 3

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

const write = "write"
const sleep = "sleep"

//Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := COUNTDOWN_START; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := COUNTDOWN_START; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, FINAL_WORD)
}
