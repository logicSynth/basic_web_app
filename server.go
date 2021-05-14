package main

import (
	"fmt"
	"net/http"
)

//handle function to process request response events
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, this is Go")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
