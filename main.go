package main

import (
	"fmt"
	"learning-golang/util"
)

func main() {
	var a = []int{}
	var b = []int{1, 2, 3, 4, 5}
	c := util.SumAllTails(a, b)
	fmt.Println(c)
}
