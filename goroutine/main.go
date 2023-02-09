package main

import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World" + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}
}

func main() {

	// 开启协程
	go test()

	// 线程任务
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}