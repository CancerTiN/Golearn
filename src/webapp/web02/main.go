package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	nExpr = 3
	word1 = "使用自定义处理器处理请求！"
	word2 = "使用自定义处理器处理请求，并且详细配置服务器！"
	word3 = "使用自定义多路处理器处理请求！"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var myWord string
	switch nExpr {
	case 1:
		myWord = word1
	case 2:
		myWord = word2
	}
	n, err := fmt.Fprintln(w, myWord, r)
	if err != nil {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) error: %v.\n", w, myWord, r, err)
	} else {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) return: (%v, %v).\n", w, myWord, r, n, err)
	}
}

func main() {
	switch nExpr {
	case 1:
		defHandler()
	case 2:
		defHandlerServer()
	case 3:
		defMux()
	}
}

func defHandler() {
	pattern := "/myHandler"
	myHandler := MyHandler{}

	http.Handle(pattern, &myHandler)

	addr := ":8080"

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe(%v, %v) error: %v.\n", addr, nil, err)
	} else {
		fmt.Printf("http.ListenAndServe(%v, %v) return: %v.\n", addr, nil, err)
	}
}

func defHandlerServer() {
	addr := ":8080"
	myHandler := MyHandler{}

	server := http.Server{Addr: addr, Handler: &myHandler, ReadTimeout: 2 * time.Second}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("server.ListenAndServe() error: %v.\n", err)
	} else {
		fmt.Printf("server.ListenAndServe() return: %v.\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := word3
	n, err := fmt.Fprintln(w, s, r.URL.Path)
	if err != nil {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) error: %v.\n", w, s, r.URL.Path, err)
	} else {
		fmt.Printf("fmt.Fprintln(%v, %v, %v) return: (%v, %v).\n", w, s, r.URL.Path, n, err)
	}
}

func defMux() {
	pattern := "/"

	mux := http.NewServeMux()
	mux.HandleFunc(pattern, handler)

	addr := ":8080"

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Printf("http.ListenAndServe(%v, %v) error: %v.\n", addr, mux, err)
	} else {
		fmt.Printf("http.ListenAndServe(%v, %v) return: %v.\n", addr, mux, err)
	}
}
