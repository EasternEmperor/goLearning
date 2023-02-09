package main

import (
	"fmt"
)

// 准备数据
func putNum(intChan chan int) {
	for i := 2; i <= 8000; i++ {
		intChan<- i
	}
	close(intChan)
}
// 判断素数
func primeNum(intChan chan int, primeChan chan int, boolChan chan bool) {
	for {
		var flag = false
		// 取数
		num, ok := <-intChan
		// 全部取完
		if !ok {
			break
		}
		for i := 2; i < num / 2; i++ {
			if num % i == 0 {
				flag = true
			}
			if flag {
				break
			}
		}
		// 素数则放入primeChan
		if !flag {
			primeChan<- num
		}
	}
	// 协程结束
	boolChan<- true
}

func main() {
	// 声明管道
	var intChan = make(chan int, 4000)
	var primeChan = make(chan int, 1000)
	var boolChan = make(chan bool, 4)
	// 启动协程
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, boolChan)
	}
	// 判断是否结束：能取到4个出来
	// 匿名函数作为协程
	go func() {
		for i := 0; i < 4; i++ {
			<-boolChan
		}
		close(primeChan)
	}()

	// 可以边计算边打印结果
	cnt := 1
	for {
		v, ok := <-primeChan
		// 如果okfalse说明素数全取完
		if !ok {
			break
		}
		// 否则打印素数
		fmt.Printf("素数%d = %d\n", cnt, v)
		cnt++
	}

	fmt.Println("主线程退出！")

}