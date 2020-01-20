package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := "Hello, World!"
	n, err := fmt.Fprintln(w, s, r.URL.Path)
	if err != nil {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) error: %v.\n", w, s, r.URL.Path, err)
	} else {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) return: (%v, %v).\n", w, s, r.URL.Path, n, err)
	}
}

func main() {
	pattern := "/"

	http.HandleFunc(pattern, handler)

	addr := ":8080"

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe(%v, %v) error: %v.\n", addr, nil, err)
	} else {
		fmt.Printf("http.ListenAndServe(%v, %v) return: %v.\n", addr, nil, err)
	}
}
