package racer

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func racer(server1 string, server2 string, sec time.Duration) (winner string, err error) {
	select {
	case <-ping(server1):
		return server1, nil
	case <-ping(server2):
		return server2, nil
	case <-time.After(sec):
		return "", fmt.Errorf("timed out waiting for %s and %s", server1, server2)
	}
}
