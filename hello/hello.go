// 输出hello world

package main		// hello.go所在的包
/*
package表示该go文件在哪个包下，这里是main包。同时在一个包下只能有一个main函数！
*/

import "fmt"

func main() {
	fmt.Println("hello, world!")
}

// go build xx.go	编译
// go build -o xx.exe xx.go		可以指定.exe文件的名字
// go build -race xxx.go		编译时查看该文件运行是否会出现资源竞争情况
// xx.exe			运行
// go run xx.go		编译运行，较慢

// 规范格式：
// gofmt xx.go		将规范格式的源代码输出
// gofmt -w xx.go	会对原文件用规范格式覆盖（写入）

// golang.org		官方网站，可以查包