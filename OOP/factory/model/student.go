package model

// 结构体首字母小写，则在其他包不能直接使用结构体名来创建变量，
// 此时可通过工厂模式！（即创建公有函数）
type student struct {
	Name string
	age int
}

func NewStudent(name string, age int) (*student) {
	return &student{
		Name: name,
		age: age,
	}
}

// 同样，由于student结构体中age首字母小写，其他包也不能直接访问，
// 通过工厂模式来访问！
func (s *student) GetAge() int {
	return s.age
}