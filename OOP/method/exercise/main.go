package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}
func (student Student) Say() string {
	str := fmt.Sprintf("name: %s\ngender: %s\nage: %d\nid: %d\nscore: %.2f",
						student.name, student.gender, student.age, student.id, student.score)
	return str
}

func main() {
	var s = Student{"Tom", "female", 18, 110, 59.9}
	var info = s.Say()
	fmt.Println(info)
}