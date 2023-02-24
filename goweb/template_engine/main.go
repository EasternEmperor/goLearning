package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	// 通过Must函数替我们解决异常
	// t, _ := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html", "hello.html"))
	// 执行
	// Execute只能输出解析的第一个模板index.html
	// 如果要解析第二个，应该使用ExecuteTemplate()来指定
	t.ExecuteTemplate(w, "hello.html", "进入hello.html")
	// t.Execute(w, "Hello")
}

func main() {
	// 调用处理器函数
	http.HandleFunc("/Template", testTemplate)

	// 监听
	http.ListenAndServe(":8080", nil)
}