package main

import (
	"fmt"
	"sort"
	"math/rand"
)

// Student结构体
type Student struct {
	Name string
	Age int
	score float64
}
// String()方法用于输出
func (s Student) String() string {
	str := fmt.Sprintf("姓名：%s 年龄：%d 成绩：%.2f", s.Name, s.Age, s.score)
	return str
}

// Student结构体切片类型：[]Student
type StudentSlice []Student
// 使StudentSlice实现Interface接口
// Len()
func (ss StudentSlice) Len() int {
	return len(ss)
}
// Less(i, j int) int
func (ss StudentSlice) Less(i, j int) bool {
	// i更小则返回true，不交换；否则false，交换
	return ss[i].score < ss[j].score
}
// Swap(i, j int)
func (ss StudentSlice) Swap(i, j int) {
	// tmp := ss[i]
	// ss[i] = ss[j]
	// ss[j] = tmp
	// 可以简写如下
	ss[i], ss[j] = ss[j], ss[i]
}

// 验证
func main() {

	// 创建切片
	var students = make(StudentSlice, 10)
	// 添加元素
	for i := 0; i < 10; i++ {
		student := Student{
			Name : fmt.Sprintf("%d号学生", rand.Intn(50)),
			Age : rand.Intn(100),
			score : rand.Float64() * 100,
		}
		students[i] = student
	}
	// 排序前打印
	fmt.Println("---------------排序前-----------------")
	for _, v := range students {
		fmt.Println(&v)
	}

	// 排序
	sort.Sort(students)

	// 排序后打印
	fmt.Println("--------------排序后-----------------")
	for _, v := range students {
		fmt.Println(&v)
	}

}