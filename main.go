package main

import (
	"learning-golang/util"
	"os"
	"time"
)

func main() {
	sleeper := &util.ConfigurableSleeper{1 * time.Second}
	util.Countdown(os.Stdout, sleeper)
}
