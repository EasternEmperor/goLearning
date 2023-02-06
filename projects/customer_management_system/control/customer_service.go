/**
结构体保存客户切片，对客户进行增删改查
*/
package control

import (
	"projects/customer_management_system/model"
)

type CustomerService struct {
	// 保存所有客户
	customers []model.Customer
	// 客户数
	cnt int
}

// 工厂模式创建CustomerService结构体变量
func NewCustomerService() *CustomerService {
	// 测试用，先加入一个客户
	// customer := model.NewCustomer(1, "Tom", true, 30, "88888888", "tom@qq.com")
	// var customers []model.Customer
	// customers = append(customers, customer)
	return &CustomerService{
		// customers : customers,
		cnt : 0,
	}
}
// 增删改查
// 查：返回客户切片以供View打印客户列表
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
// 增: 接收View提供的客户信息，在切片中增加客户
func (this *CustomerService) Add(name string, gender bool, age int, pn string, email string) {
	c := model.NewCustomer(this.cnt + 1, name, gender, age, pn, email)
	this.customers = append(this.customers, c)
	
	// 客户数+1
	this.cnt++
}
// 删
func (this *CustomerService) Del(id int) bool {
	index := this.FindById(id)
	// 找到该客户，删除用append([:index], [index+1:]...)
	if index != -1 {
		// 删除idx
		this.customers = append(this.customers[ : index], this.customers[index + 1 : ]...)
		return true
	} else {
		return false
	}
}
// 按id查找客户: 返回下标, 返回-1则说明没有该id的客户
func (this *CustomerService) FindById(id int) int {
	idx := -1		// 下标
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			idx = i
		}
	}
	return idx
}
// 改
func (this *CustomerService) ModifyCustomer(idx int, name string, gender string, 
											age int, pn string, email string) {
	// 判断是否要修改
	if name != "" {
		this.customers[idx].ModifyName(name)
	}
	if gender != "" {
		var flag bool
		if gender == "男" {
			flag = false
		} else {
			flag = true
		}
		this.customers[idx].ModifyGender(flag)
	}
	if age != -1 {
		this.customers[idx].ModifyAge(age)
	}
	if pn != "" {
		this.customers[idx].ModifyPhoneNumber(pn)
	}
	if email != "" {
		this.customers[idx].ModifyEmail(email)
	}
}