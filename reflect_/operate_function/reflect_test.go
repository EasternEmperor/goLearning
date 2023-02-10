package test

import (
	"testing"
	"reflect"
)

func TestReflect(t *testing.T) {
	// 要操作的匿名函数
	call1 := func(n1, n2 int) {
		t.Log(n1, n2)
	}
	call2 := func(n1, n2 int, s string) {
		t.Log(n1, n2, s)
	}
	
	// 反射操作
	brige := func(f interface{}, args ...interface{}) {
		// 参数个数
		n := len(args)
		// 反射调用参数切片
		params := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			params[i] = reflect.ValueOf(args[i])
		}
		// 获取封装函数的Value
		function := reflect.ValueOf(f)
		function.Call(params)
	}

	// 通过反射来调用call1/2
	brige(call1, 10, 100)
	brige(call2, 1000, 10000, "Hello")
	
}