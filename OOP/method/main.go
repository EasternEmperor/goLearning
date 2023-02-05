package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}
func (person Person) test() {
	person.Name = "Jack"		// 不改变实参
	fmt.Println("test() name =", person.Name)
}
func (person *Person) test2() {
	person.Name = "Jack"		// 编译器底层会自动加上取值符*	
	fmt.Println("test() name =", person.Name)
}
func (person Person) String() string {
	str := fmt.Sprintf("Name = [%v] Age = [%v]", person.Name, person.Age)
	return str
}

// 自定义类型都有方法
type Integer int
func (integer Integer) print() {
	fmt.Println("Integer =", integer)
}
func (integer *Integer) add1() {
	*integer++		// 注意取值符*
}

func main() {

	// test()
	var p Person
	p.Name = "Tom"
	fmt.Println(p)
	p.test()
	fmt.Println(p)

	// test2()
	fmt.Println(p)
	// 标准写法：&p.test2()，但是同样，Go编译器会自动加&，所以也可以不加&
	p.test2()
	fmt.Println(p)

	// Integer
	var i = Integer(10)		// 注意需强转！
	i.print()
	// 实际上应为(&i).add1()
	i.add1()				// 这里同样可以不用地址符&
	i.print()

	// 自定义类型的String()方法
	fmt.Println(&p)

	// 值拷贝和地址拷贝与否，只看绑定的类型，与调用形式无关
	var p2 Person
	p2.Name = "Mary"
	fmt.Println(&p2)
	(&p2).test()		// 不改变实参
	fmt.Println(&p2)		


}