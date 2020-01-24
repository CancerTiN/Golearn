package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		n   int
		err error
	)

	// GET
	n, err = fmt.Fprintln(w, "发送请求的请求地址：", r.URL.Path)
	fmt.Printf("fmt.Fprintln1 return: (%v, %v).\n", n, err)
	n, err = fmt.Fprintln(w, "发送请求的查询字串：", r.URL.RawQuery)
	fmt.Printf("fmt.Fprintln2 return: (%v, %v).\n", n, err)
	n, err = fmt.Fprintln(w, "请求头中的所有信息：", r.Header)
	fmt.Printf("fmt.Fprintln3 return: (%v, %v).\n", n, err)
	n, err = fmt.Fprintln(w, "请求头中的Accept-Encoding信息：", r.Header["Accept-Encoding"])
	fmt.Printf("fmt.Fprintln4 return: (%v, %v).\n", n, err)
	n, err = fmt.Fprintln(w, "请求头中的Accept-Encoding属性：", r.Header.Get("Accept-Encoding"))
	fmt.Printf("fmt.Fprintln5 return: (%v, %v).\n", n, err)

	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(n)
	fmt.Printf("rand.Intn(%d) return: (%v).\n", n, rn)
	if rn > n/2 {
		// POST Form Data
		length := r.ContentLength
		body := make([]byte, length)
		n, err = r.Body.Read(body)
		fmt.Printf("r.Body.Read return: (%v, %v).\n", n, err)
		n, err = fmt.Fprintln(w, "请求体中的内容有：", string(body))
		fmt.Printf("fmt.Fprintln6 return: (%v, %v).\n", n, err)
	} else {
		// POST Form Data Parse
		err = r.ParseForm()
		fmt.Printf("r.ParseForm return: (%v).\n", err)
		n, err = fmt.Fprintln(w, "请求参数被解析为：", r.Form)
		fmt.Printf("fmt.Fprintln7 return: (%v, %v).\n", n, err)
		n, err = fmt.Fprintln(w, "POST请求的Form表单被解析为：", r.PostForm)
		fmt.Printf("fmt.Fprintln8 return: (%v, %v).\n", n, err)
	}
}

func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("http.ListenAndServe return: (%v).\n", err)
}
