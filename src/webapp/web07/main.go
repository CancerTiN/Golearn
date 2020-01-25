package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/web04/model"
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

	n, err = fmt.Fprintln(w, "URL中的user的值为：", r.FormValue("user"))
	fmt.Printf("fmt.Fprintln6 return: (%v, %v).\n", n, err)
	n, err = fmt.Fprintln(w, "POST表单中的username的值为：", r.PostFormValue("username"))
	fmt.Printf("fmt.Fprintln7 return: (%v, %v).\n", n, err)
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{ID: 1, Username: "admin", Password: "123456", Email: "admin@gmail.com"}
	data, err := json.Marshal(user)
	fmt.Printf("json.Marshal return: (%v).\n", err)
	n, err := w.Write(data)
	fmt.Printf("w.Write return: (%v, %v).\n", n, err)
}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedirect)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("http.ListenAndServe return: (%v).\n", err)
}
