package main

import (
	"net/http"
	"text/template"
)

// 处理器函数去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行模板，不传入
	t.Execute(w, "")
}

func main() {
	// 设置处理静态资源，如css和js文件
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages", http.StripPrefix("/pages/", http.FileServer(http.Dir("/views/pages/"))))
	// 展示html页面
	http.HandleFunc("/main", IndexHandler)
	// 监听
	http.ListenAndServe(":8080", nil)
}