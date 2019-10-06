package main

import (
	"fmt"
	u "learning-golang/util"
)

func main() {
	var n u.Mydata
	var num int
	n.Numbers = [...]int{1, 2, 3, 4, 5, 6}
	for _, number := range n.Numbers {
		num += number
	}
	fmt.Printf("%v", num)
}
