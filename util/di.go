package util

import (
	"fmt"
	"io"
)

//Greet ...
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
