package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看CPU数
	cpuNum := runtime.NumCPU()
	fmt.Println("本机CPU数：", cpuNum)

	// 设置运行的CPU数
	preRun := runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("启动时运行的CPU个数：", preRun)
}