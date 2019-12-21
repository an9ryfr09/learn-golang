package deleyexec

import "testing"

import "reflect"

func TestDefer(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	want := nums
	got := deferfunc(nums)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
