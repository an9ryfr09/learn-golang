package main

import "fmt"

type Str string

func (s *Str) print() {
	fmt.Println(*s)
}

func main() {
	var message Str = "golang is the best language!!!!!!!"
	message.print()
}
