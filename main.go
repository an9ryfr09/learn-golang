package main

import (
	"fmt"
	u "learning-golang/util"
)

func main() {
	var a = []int{1, 2, 3}
	var b = []int{1, 2, 3, 4, 5}
	c := u.SumAll(a, b)
	fmt.Println(c)
}
