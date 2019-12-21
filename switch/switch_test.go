package control

import "testing"

func TestSwitch(t *testing.T) {
	want := "string"
	got := mySwitch("abc")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
