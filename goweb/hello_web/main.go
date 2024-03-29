package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path是输入地址带的信息
	fmt.Fprintln(w, "Hello World!", r.URL.Path)
}


func main() {
	http.HandleFunc("/", handler)

	// 创建路由
	http.ListenAndServe(":8080", nil)
}