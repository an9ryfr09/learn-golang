package main

import (
	"fmt"
	"learning-golang/util"
	"reflect"
	"strings"
	"testing"
)

func checkSumsSlice(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkFloat(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestHello(t *testing.T) {
	name := "an9ryfr09"
	got := util.Hello(name)
	want := util.HELLO_PREFIX + name
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	var num util.Mydata
	num.Num = 5
	num2 := 9
	got := num.Add(num2)
	want := num.Num + num2
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func Example_add() {
	var num util.Mydata
	num.Num = 5
	num2 := 5
	fmt.Println(num.Add(num2))
	// output: 10
}

func Example_printHello() {
	var name = "yanlei"
	fmt.Println(util.Hello(name))
	//output: hello,yanlei
}

func TestRepeat(t *testing.T) {
	ch := "a"
	count := 5
	toUpper := true
	repeated := util.Repeat(ch, count, toUpper)
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
		util.Repeat(ch, time, toUpper)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := util.Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestAllSum(t *testing.T) {
	got := util.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	checkSumsSlice(t, got, want)
}

func TestAllSumTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := util.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSumsSlice(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := util.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSumsSlice(t, got, want)
	})
}

func TestPerimeter(t *testing.T) {
	var rectangle = &util.Rectangle{20.0, 10.0}
	got := rectangle.Perimeter()
	want := 60.0
	checkFloat(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   util.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &util.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: &util.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &util.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}

}
