package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"webproject/bookstore/dao"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//拿到请求参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//登录校验
	user, err := dao.CheckLogin(username, password)
	if err != nil {
		fmt.Println("sql异常: ", err)
	}
	if user != nil {
		//用户名,密码正确,跳转到"登录成功页"
		t := template.Must(template.ParseFiles("bookstore\\views\\pages\\user\\login_success.html"))
		t.Execute(w, "")
	} else {
		//用户名,密码不正确
		t := template.Must(template.ParseFiles("bookstore\\views\\pages\\user\\login.html"))
		t.Execute(w, "用户名或密码不正确")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	b, _ := dao.CheckRegs(username)
	if !b {
		//false 查不到,用户名不存在,校验成功. 新增用户
		err := dao.SaveUser(username, password, email)
		if err != nil {
			fmt.Println("新增用户失败: ", err)
		}
		//跳转到"注册成功页"
		t := template.Must(template.ParseFiles("bookstore\\views\\pages\\user\\regist_success.html"))
		t.Execute(w, "")
	} else {
		//用户名已存在,校验失败
		t := template.Must(template.ParseFiles("bookstore\\views\\pages\\user\\regist.html"))
		t.Execute(w, "用户名已存在")
	}
}
