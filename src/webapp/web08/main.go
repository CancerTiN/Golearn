package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//filenames := "D:/Workspace/Golearn/src/webapp/web08/index.html"
	//t, err := template.ParseFiles(filenames)
	//fmt.Printf("template.ParseFiles return: (%v, %v).\n", t, err)
	//err = t.Execute(w, "Hello Template")
	//fmt.Printf("t.Execute return: (%v).\n", err)

	filename0 := "D:/Workspace/Golearn/src/webapp/web08/index.html"
	filename1 := "D:/Workspace/Golearn/src/webapp/web08/indexalpha.html"
	t := template.Must(template.ParseFiles(filename0, filename1))
	err := t.ExecuteTemplate(w, filepath.Base(filename1), "我要去indexAlpha中")
	fmt.Printf("t.ExecuteTemplate return: (%v).\n", err)
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("http.ListenAndServe return: (%v).\n", err)
}
