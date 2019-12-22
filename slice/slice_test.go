package slice

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("bytes.Equel", func(t *testing.T) {
		a := [5]byte{'1', '2', '3', '4', '5'}
		b := a[1:3]
		got := b
		want := []byte{'2', '3'}
		if !bytes.Equal(got, want) {
			t.Error(b)
		}
	})

	t.Run("equleSliceOfInt", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		_ = copy(a, b)
		got := a
		want := []int{4, 5, 6}
		if !equelSliceOfInt(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("reflect.DeepEquel", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		_ = copy(a, b)
		got := a
		want := []int{4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("appendInt", func(t *testing.T) {
		a := []int{1, 2, 3}
		got := appendInt(a, 5)
		want := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
