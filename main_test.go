package main

import (
	"fmt"
	an "learning-golang/an9ryfr09"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	name := "an9ryfr09"
	got := an.Hello(name)
	want := an.HELLO_PREFIX + name
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	var num an.Mydata
	num.Num = 5
	num2 := 9
	got := num.Add(num2)
	want := num.Num + num2
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func ExampleAdd() {
	var num an.Mydata
	num.Num = 5
	num2 := 5
	fmt.Println(num.Add(num2))
	// output: 10
}

func ExamplePrintHello() {
	var name = "yanlei"
	fmt.Println(an.Hello(name))
	//output: hello,yanlei
}

func TestRepeat(t *testing.T) {
	ch := "a"
	count := 5
	toUpper := true
	repeated := an.Repeat(ch, count, toUpper)
	expected := strings.ToUpper(strings.Repeat(ch, 5))
	if repeated != expected {
		t.Errorf("expected: '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	ch := "b"
	time := 1000
	toUpper := true

	for i := 0; i < b.N; i++ {
		an.Repeat(ch, time, toUpper)
	}
}

func TestSum(t *testing.T) {
	var num an.Mydata
	num.Numbers = [...]int{1, 2, 3, 4, 5, 6}
	got := num.Sum()
	want := 21
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, num)
	}
}
