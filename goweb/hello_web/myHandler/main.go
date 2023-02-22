package main

import (
	"fmt"
	"net/http"
	"time"
)

// 创建自己的处理器
type MyHandler struct {}
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "我自己的处理器！", r.URL.Path)
}

func main() {
	myHandler := &MyHandler{}

	// 调用我的处理器
	http.Handle("/MyHandler", myHandler)

	// 监听端口
	// http.ListenAndServe(":8080", nil)

	// 创建server实例处理连接请求
	server := &http.Server{
		Addr : ":8080",
		Handler : myHandler,
		ReadTimeout : 2 * time.Second,
	}
	server.ListenAndServe()

}