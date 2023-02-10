package test

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name string
}

func TestReflect(t *testing.T) {

	// 声明要用到的变量
	var (
		model *User
		st reflect.Type
		elem reflect.Value
	)

	// 获取类型*User（注意是指针）
	st = reflect.TypeOf(model)
	// 输出类型: ptr
	t.Log("reflect.TypeOf: ", st.Kind().String())
	// 更改为Type实例
	st = st.Elem()
	// 输出此时类型
	t.Log("reflect.TypeOf.Elem: ", st.Kind().String())
	// 分配实例空间，获取指向该空间的指针*Value
	elem = reflect.New(st)
	// 输出类型
	t.Log("reflect.New: ", elem.Kind().String())
	t.Log("reflect.New.Elem: ", elem.Elem().Kind().String())

	// model指向新分配的空间
	// 先将elem转为interface{}, 再用类型断言转为*User类型
	model = elem.Interface().(*User)

	// 更改为Value实例
	elem = elem.Elem()
	// 给实例字段赋值
	elem.FieldByName("UserId").SetString("1234567")
	elem.FieldByName("Name").SetString("钟起龙")

	// 打印
	t.Log("model: ", model, "model.Name", model.Name)

}