package bytes

import "testing"
import "bytes"

func TestBytes(t *testing.T) {
	var b bytes.Buffer
	var delim = byte('e')
	b.Write([]byte("abcde"))
	got, err := b.ReadBytes(delim)
	if err != nil {
		t.Error(err)
	}
	want := []byte("abcde")
	if !bytes.Equal(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
