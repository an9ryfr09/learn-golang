package main

import (
	"fmt"
	"learning-golang/util"
)

func checkArea(shape *util.Rectangle, want float64) {
	got := shape.Area()
	fmt.Printf("got %.2f want %.2f", got, want)
}

func main() {
	var rectangle = util.Rectangle{Width: 12, Height: 6}
	want := 72.0
	checkArea(&rectangle, want)
}
