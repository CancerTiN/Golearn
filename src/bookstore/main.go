package main

import (
	"bookstore/controller"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml := "views/index.html"
	fileInfo, err := os.Stat(indexHtml)
	fmt.Printf("os.Stat return: (%v, %v).\n", fileInfo, err)
	if err != nil {
		indexHtml = "D:/Workspace/Golearn/src/bookstore/views/index.html"
	}
	t := template.Must(template.ParseFiles(indexHtml))
	err = t.Execute(w, "")
	fmt.Printf("t.Execute return: (%v).\n", err)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", IndexHandler)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/regist", controller.Regist)

	http.HandleFunc("/checkUsername", controller.CheckUsername)

	fmt.Println("Server startup ...")
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("http.ListenAndServe return: (%v).\n", err)
}
