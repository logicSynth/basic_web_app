package main

import (
	"fmt"
	"net/http"
)

/* creat a handler to handle function to process
request response events */
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, this is Go")
}

func main() {
	http.HandleFunc("/", index_handler)
	/* create a request listener: that listen to  a port
	and passes connection to a port */
	http.ListenAndServe(":8080", nil)
}
