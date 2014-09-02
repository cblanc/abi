package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9999", nil)
}
