package controller

import (
	"bookstore/dao"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Printf("r.PostFormValue('username') return: (%v).\n", username)
	fmt.Printf("r.PostFormValue('password') return: (%v).\n", password)
	user, err := dao.CheckUsernameAndPassword(username, password)
	fmt.Printf("dao.CheckUsernameAndPassword return: (%v, %v).\n", user, err)
	if err == nil && user.ID != 0 {
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		_ = t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		_ = t.Execute(w, "用户名或密码不正确！")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	fmt.Printf("r.PostFormValue('username') return: (%v).\n", username)
	fmt.Printf("r.PostFormValue('password') return: (%v).\n", password)
	fmt.Printf("r.PostFormValue('email') return: (%v).\n", email)
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		_ = t.Execute(w, "")
		return
	}
	user, err := dao.CheckUsername(username)
	fmt.Printf("dao.CheckUsernameAndPassword return: (%v, %v).\n", user, err)
	if err == nil && user.ID != 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		_ = t.Execute(w, "用户名已存在！")
	} else {
		err = dao.SaveUser(username, password, email)
		fmt.Printf("dao.SaveUser return: (%v).\n", err)
		var t *template.Template
		if err != nil {
			t = template.Must(template.ParseFiles("views/pages/user/regist.html"))
		} else {
			t = template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		}
		_ = t.Execute(w, "")
	}
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	fmt.Printf("r.PostFormValue('username') return: (%v).\n", username)
	user, err := dao.CheckUsername(username)
	fmt.Printf("dao.CheckUsernameAndPassword return: (%v, %v).\n", user, err)
	if err == nil && user.ID != 0 {
		_, _ = w.Write([]byte("用户名已存在！"))
	} else {
		_, _ = w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
