package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	name := "an9ryfr09"
	got := hello(name)
	want := "hello," + name
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
