package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Monster struct {
	Name string		`json:"name"`
	age int
	Skill string	`json:"skill"`
	Level int		`json:"level"`
}
// 定义其方法
func (m Monster) String() string {
	str := fmt.Sprintf("---------妖精----------\n%s\n%d\n%s\n%d", 
						m.Name, m.age, m.Skill, m.Level)
	return str
}
func (m Monster) Cal(n1, n2 int) int {
	return n1 + n2 + m.age
}
func (m Monster) Print() {
	fmt.Println(&m)
	fmt.Println("-----------------------")
}

// 反射结构体
func ReflectStruct(i interface{}) {

	// Value
	rVal := reflect.ValueOf(i)
	// Type
	rTyp := reflect.TypeOf(i)

	// 判断类别是否为结构体，只有结构体才能继续，否则退出函数
	kind := rVal.Kind()
	if kind != reflect.Struct {
		fmt.Println("\n传入变量非结构体，退出...")
		return
	}

	// 获取字段个数
	fieldNum := rVal.NumField()
	fmt.Printf("结构体有%d个字段\n", fieldNum)
	// 输出字段，若有tag则一并输出
	for i := 0; i < fieldNum; i++ {
		// 获取第i个字段的封装Value变量
		fieldVal := rVal.Field(i)
		// 获取该字段的名字
		fieldName := rTyp.Field(i).Name
		// 输出字段+值
		fmt.Printf("第%d个字段为：%s，值为：%v\n", i, fieldName, fieldVal)
		// 获取该字段key为json的tag
		fieldTag := rTyp.Field(i).Tag.Get("json")
		// 有返回则输出
		if fieldTag != "" {
			fmt.Println("----->该字段json tag =", fieldTag)
		}
	}

	fmt.Println()

	// 获取方法个数
	methodNum := rVal.NumMethod()
	fmt.Printf("结构体有%d个方法\n", methodNum)
	// 调用第2个方法
	rVal.Method(1).Call(nil)
	// 调用第1个方法
	// 设置两个参数
	var params []reflect.Value
	num1 := reflect.ValueOf(10)
	num2 := reflect.ValueOf(101)
	params = append(params, num1)
	params = append(params, num2)
	// 接收返回值
	M0Res := rVal.Method(0).Call(params)
	fmt.Println("Method 0 返回值：", M0Res[0].Int())

}

// 反射修改结构体变量字段值
func ReflectSetField(i interface{}) {

	// Value
	rVal := reflect.ValueOf(i)

	// 获取字段个数
	fieldNum := rVal.NumField()
	// 循环修改字段值
	for i := 0; i < fieldNum; i++ {
		// 获取封装字段的Value变量
		fieldVal := rVal.Field(i)
		switch fieldVal.Kind() {
		case reflect.Int, reflect.Int64:
			fieldVal.
		}
	}

}

func main() {
	// 声明结构体
	m := Monster{"九尾狐", 500, "欺诈宝珠", 99}

	// 反射结构体
	ReflectStruct(m)
	
	// 试图反射int
	num := 10
	ReflectStruct(num)
}