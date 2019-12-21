package translocation

import "testing"

func TestTranslocation(t *testing.T) {
	t.Run("left translocation:", func(t *testing.T) {
		num := 10
		bit := 4
		got := leftTranslocation(num, bit)
		want := 160
		if got != want {
			t.Errorf("got:%d, want:%d", got, want)
		}
	})

	t.Run("right translocation:", func(t *testing.T) {
		num := 4
		bit := 2
		got := rightTranslocation(num, bit)
		want := 1
		if got != want {
			t.Errorf("got:%d, want:%d", got, want)
		}
	})
}
