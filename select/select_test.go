package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var timeoutSec time.Duration

func mockHttpServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	timeoutSec = 5
	fastServer := mockHttpServer(4 * time.Second)
	slowServer := mockHttpServer(5 * time.Second)

	winner, err := racer(fastServer.URL, slowServer.URL, timeoutSec*time.Second)
	want := fastServer.URL

	if err != nil {
		t.Errorf("%v", err)
	}

	if winner != want {
		t.Errorf("winner not is %s", fastServer.URL)
	}
}
