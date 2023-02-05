package main

import (
	"fmt"
)

// 基类，共有字段和方法
type Student struct {
	Name string
	age int
	grade float64
}
// 打印成绩
func (stu Student) PrintGrade() {
	fmt.Println(&stu)
}
// String()
func (stu *Student) String() string {
	str := fmt.Sprintf("姓名：%s\n年龄：%d\n成绩：%.2f", stu.Name, stu.age, stu.grade)
	return str
}
// 初始化（工厂模式）
func (stu *Student) Initialization(name string, age int, grade float64) {
	stu.Name = name
	stu.age = age
	stu.grade = grade
}

// 派生类：小学生
type Pupil struct {
	Student		// 匿名结构体
}
// 考试
func (stu *Pupil) TakingExam() {
	fmt.Println("小学生", stu.Name, "在教室里考试~")
}

// 派生类：大学生
type UnderGraduate struct {
	Student		// 匿名
}
// 考试
func (stu *UnderGraduate) TakingExam() {
	fmt.Println("大学生", stu.Name, "在地狱里考试~")
}


// 嵌套有名结构体
type Graduate struct {
	s Student		// 有名结构体
}


// 匿名字段是基本数据类型
type Monster struct {
	Name string
	power float64
}
type SuperMonster struct {
	Monster			// 匿名结构体
	int				// 匿名数据类型
	n int
}


func main() {

	// 创建小学生变量
	var p = &Pupil{}
	// 初始化
	p.Student.Initialization("Tom", 8, 70.0)
	// 调用父类方法
	p.Student.PrintGrade()

	// 创建大学生变量
	var ug = &UnderGraduate{}
	// 初始化
	ug.Student.Initialization("Mary", 24, 59)
	// 调用父类方法
	ug.Student.PrintGrade()


	// 创建研究生变量：注意其中嵌套了有名结构体
	var gg Graduate
	//gg.Name = "Smith"		// 报错
	gg.s.Name = "Smith"
	gg.s.PrintGrade()


	// 在创建时给匿名结构体字段赋值
	var p2 = &Pupil{Student{"Jack", 10, 99}}
	p2.PrintGrade()

	
	// 匿名字段是基本数据类型，要直接使用该数据类型名来访问
	var sm SuperMonster
	sm.Name = "变形金刚"
	sm.power = 14.5
	sm.int = 100
	sm.n = 50
	fmt.Println("超兽变形金刚：", sm)
}