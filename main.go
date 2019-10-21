package main

import (
	"learning-golang/util"
	"net/http"
)

//MyGreeterHandler ...
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	util.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
