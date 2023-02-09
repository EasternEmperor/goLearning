package main

import (
	"fmt"
)

func main() {
	// 只读管道
	// var intChan <-chan int
	// 只写管道
	// var intChan chan<- int

	// select
	var chan1 = make(chan int, 5)
	var chan2 = make(chan string, 10)
	for i := 0; i < 5; i++ {
		chan1<- i
	}
	for i := 0; i < 10; i++ {
		chan2<- "Hello World" + fmt.Sprintf("%d", i)
	}
	// select 取数据
//label:
	for {
		select {
		case v := <-chan1:
			fmt.Println("取chan1:", v)
		case v := <-chan2:
			fmt.Println("取chan2:", v)
		default:
			fmt.Println("取不到了，在此处加入后续逻辑")
			return 
			//break label
		}
	}
}