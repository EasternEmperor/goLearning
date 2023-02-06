package model

import "fmt"


// 客户结构体
type Customer struct {
	Id int
	name string
	gender bool			// true: female, false: male
	age int
	PhoneNumber string
	email string
}

// 工厂模式创建Customer结构体
func NewCustomer(id int, name string, gender bool, age int, phoneNumber string, email string) Customer {
	return Customer{
		Id : id,
		name : name,
		gender : gender,
		age : age,
		PhoneNumber : phoneNumber,
		email : email,
	}
}
// 工厂模式返回字段
func (c *Customer) GetName() string {
	return c.name
}
func (c *Customer) GetGender() string {
	if c.gender {
		return "女"
	} else {
		return "男"
	}
}
func (c *Customer) GetAge() int {
	return c.age
}
func (c *Customer) GetPhoneNumber() string {
	return c.PhoneNumber
}
func (c *Customer) GetEmail() string {
	return c.email
}

// 改
func (c *Customer) ModifyId(id int) {
	c.Id = id
}
func (c *Customer) ModifyName(name string) {
	c.name = name
}
func (c *Customer) ModifyGender(gender bool) {
	c.gender = gender
}
func (c *Customer) ModifyAge(age int) {
	c.age = age
}
func (c *Customer) ModifyPhoneNumber(pn string) {
	c.PhoneNumber = pn
}
func (c *Customer) ModifyEmail(email string) {
	c.email = email
}

// String()输出方法
func (c *Customer) String() string {
	gender := ""
	// 判断性别
	if c.gender {	// true为女female
		gender = "女"
	} else {
		gender = "男"
	}
	str := fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s", c.Id, c.name,
			gender, c.age, c.PhoneNumber, c.email)
	return str
}