package main

import (
	"fmt"
	_ "time"
)

// 一个写协程向管道中写入数据，一个读协程向管道中读取数据
// 写
func WriteChan(intChan chan int) {
	// 写入50个数据
	for i := 0; i < 100; i++ {
		intChan<- i
		fmt.Println("写入数据: ", i)
	}
	// 关闭管道
	close(intChan)
}
// 读
func ReadChan(intChan chan int, boolChan chan bool) {
	// 读取数据
	for {
		v, ok := <-intChan
		// 如果管道关闭，且已读完，则ok为false
		if !ok {
			break
		}
		// 打印
		fmt.Println("读取数据：", v)
	}
	// 读完后要让主线程停止忙等，向boolChan写入true表明可以退出
	boolChan<- true
	close(boolChan)
}

func main() {

	// 声明channel变量
	var intChan chan int
	// make分配空间
	intChan = make(chan int, 3)
	fmt.Println("intChan =", intChan, "\n&intChan =", &intChan)
	// 存放数据
	num := 100
	intChan<- 10
	intChan<- num
	// 获取长度和容量
	fmt.Println("intChan长度=", len(intChan), "\tintChan容量=", cap(intChan))
	// 读取数据
	num = <-intChan
	<-intChan		// 抛弃该数据
	fmt.Println("intChan长度=", len(intChan), "\tintChan容量=", cap(intChan))
	fmt.Println("读取到num =", num)

	// 创建读写协程，判断何时可以结束
	var intChan2 = make(chan int, 50)
	var boolChan = make(chan bool, 1)
	go WriteChan(intChan2)				// 写
	go ReadChan(intChan2, boolChan)	// 读
	// 判断读写是否结束
	for {
		_, ok := <-boolChan
		if !ok {
			fmt.Println("读取完毕！")
			break
		}
	}

}