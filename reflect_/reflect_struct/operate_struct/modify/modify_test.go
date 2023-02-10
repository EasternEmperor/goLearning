package test

import (
	"testing"
	"reflect"
)

type user struct {
	Name string
	Pwd string
}

func TestReflectStruct(t *testing.T) {
	// 声明user变量
	var model *user = &user{}
	var sv reflect.Value

	// 获取结构体变量
	sv = reflect.ValueOf(model)
	// 输出sv的类别
	t.Log("reflect.ValueOf: ", sv.Kind().String())		// ptr: 即指针
	// 修改字段
	sv = sv.Elem()
	sv.FieldByName("Name").SetString("nickname")
	sv.FieldByName("Pwd").SetString("1234567")
	// 打印修改后的model值
	t.Log("model:", model)
}