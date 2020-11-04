package main

import (
	"html/template"
	"net/http"
	"webproject/bookstore/controller"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("bookstore\\views\\index.html"))
	t.Execute(w, "")
}
func main() {
	//静态资源过滤
	//访问以"/static/"为前缀的资源(css样式等), 会定位匹配到的其相对路径
	//即将"/static/"替换为"bookstore/views/static"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore\\views\\static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore/views/pages"))))
	//首页
	http.HandleFunc("/index", IndexHandler)
	//登录
	http.HandleFunc("/login", controller.Login)
	//注册
	http.HandleFunc("/regist", controller.Regist)

	http.ListenAndServe(":8080", nil)
}
