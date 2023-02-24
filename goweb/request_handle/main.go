package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// 结构体
type User struct {
	id int
	Username string
	Password string
	Email string
}

// 处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送请求的请求地址后的查询字符是：", r.URL.RawQuery)
	// 请求头Header
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息有：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性是：", r.Header.Get("Accept-Encoding"))

	// 解析表单，在调用r.PostForm之前必须执行ParseForm()
	r.ParseForm()
	// 获取请求参数
	// r.Form字段：如果url中也有与form表单参数名相同的请求参数(如...?username=xxx&password=xxx)
	// 那么参数值都可以在r.Form中得到，并且form表单中的参数值会排在前面
	fmt.Fprintln(w, "请求参数有：", r.Form)
	fmt.Fprintln(w, "POST请求中form表单中的请求参数有：", r.PostForm)

	// 直接调用FormValue和PostFormValue获取请求参数
	fmt.Fprintln(w, "URL中的user请求参数：", r.FormValue("user"))
	fmt.Fprintln(w, "Post表单中username请求参数：", r.PostFormValue("username"))

	// 获取请求体的长度
	len := r.ContentLength
	// 创建byte切片以保存请求体内容
	body := make([]byte, len)
	// 读取请求体内容
	r.Body.Read(body)
	// 打印请求体
	fmt.Fprintln(w, "请求体中的内容有：", string(body))
}

// 回复客户端json数据
func jsonRes(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容的类型: application/json
	w.Header().Set("Content-Type", "application/json")
	// 创建User
	user := User{1, "king", "123", "king@qq.com"}
	// 序列化以发送
	data, _ := json.Marshal(user)
	// 返回给客户端
	w.Write(data)
}

// 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 设置重定向地址
	w.Header().Set("location", "https://www.baidu.com")
	// 设置响应码
	w.WriteHeader(302)
}

func main() {
	// 调用处理器函数
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", jsonRes)
	http.HandleFunc("/testRedirect", Redirect)
	// 监听
	http.ListenAndServe(":8080", nil)
}