package util

import (
	"fmt"
	"io"
	"time"
)

const FINAL_WORD = "Go!"
const COUNTDOWN_START = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	Duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.Duration)
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
