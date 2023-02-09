package cal

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {

	// 调用
	res := addUpper(10)
	if res != 55 {		// 函数出错
		// t.Fatalf()输出错误信息，并退出
		t.Fatalf("AddUpper(10) 执行错误，期望值=%d, 实际值=%d\n", 55, res)
	}

	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...\n")
}

func TestHello(t *testing.T) {
	// 不调用t也不会报错
	fmt.Println("测试Hello()...")
}