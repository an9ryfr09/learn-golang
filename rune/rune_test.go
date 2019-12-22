package rune

import (
	"reflect"
	"testing"
)

func TestRune(t *testing.T) {
	var runes = []rune("hello, ")
	for _, r := range "世界!" {
		runes = append(runes, r)
	}
	got := runes
	want := []rune("hello, 世界!")
	if reflect.DeepEqual(got, want) {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
