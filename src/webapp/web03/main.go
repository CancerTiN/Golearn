package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "r: %v\n", r)
	if err != nil {
		fmt.Println("fmt.Fprintf error:", err)
	} else {
		fmt.Printf("fmt.Fprintf return: (%v, %v)\n", n, err)
	}
	n, err = fmt.Fprintln(w, "Test HTTP protocol!")
	if err != nil {
		fmt.Println("fmt.Fprintf error:", err)
	} else {
		fmt.Printf("fmt.Fprintf return: (%v, %v)\n", n, err)
	}
}

func main() {
	http.HandleFunc("/http", handler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("fmt.Fprintf return:", err)
}
