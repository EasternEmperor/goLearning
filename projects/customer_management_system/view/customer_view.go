package main

import (
	"fmt"
	 "projects/customer_management_system/control"
	// "projects/customer_management_system/model"
)

// 结构体
type customerView struct {

	// 定义必要字段
	choice string		// 接收选项
	loop bool			// 判断是否需要重复打印主菜单

	// 所有客户信息和方法
	customerService *control.CustomerService

}
// 工厂模式创建customerView实例
func NewCustomerView() customerView {
	return customerView{"", true, control.NewCustomerService()}
}
// 显示主菜单的方法
func (this *customerView) MainMenu() {

	// 循环打印菜单
	for this.loop {
		fmt.Println("----------------客户信息管理软件-----------------")
		fmt.Println("                1 添加客户")
		fmt.Println("                2 修改客户")
		fmt.Println("                3 删除客户")
		fmt.Println("                4 客户列表")
		fmt.Println("                5 退出")
		fmt.Print("请选择(1-5):")

		// 获取选项
		fmt.Scanln(&this.choice)

		// 根据选项输出
		switch (this.choice) {
		case "1":
			// 添加客户
			this.add()
		case "2":
			// 修改客户
			this.modify()
		case "3":
			// 删除客户
			this.del()
		case "4":
			// 打印客户列表
			this.list()
		case "5":
			// 退出菜单
			this.exit()
		default:
			fmt.Println("选项错误，请重新输入...\n\n")
		}

	}

}
// 增删改查
// 查：打印客户列表
func (this *customerView) list() {
	fmt.Println("-------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	// 从Control里的List()方法获取到客户切片，来输出信息
	for _, v := range this.customerService.List() {
		fmt.Println(&v)
	}
	fmt.Println("\n-----------------客户列表完成------------------\n\n")
}
// 增
func (this *customerView) add() {
	fmt.Println("-------------------添加客户-------------------")

	// 待使用字段
	var name = ""
	var gender = ""
	var age = 0
	var pn = ""
	var email = ""

	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&pn)
	fmt.Print("邮箱：")
	fmt.Scanln(&email)

	var flag bool
	if gender == "男" {
		flag = false
	} else {
		flag = true
	}
	this.customerService.Add(name, flag, age, pn, email)

	fmt.Println("\n-----------------添加客户完成------------------\n\n")
}
// 删
func (this *customerView) del() {
	fmt.Println("-------------------删除客户-------------------")

	var id int
	var yn string
	fmt.Print("请输入要删除的客户id（-1退出）：")
	fmt.Scanln(&id)

	// 退出删除
	if id == -1 {
		fmt.Println("\n")
		return 
	}

	fmt.Print("确认是否删除(Y/N): ")
	fmt.Scanln(&yn)

	// 不删除
	if yn != "Y" && yn != "y" {
		fmt.Println("\n")
		return 
	}

	if this.customerService.Del(id) {
		fmt.Println("-------------------删除完成-------------------\n\n")
	} else {
		fmt.Println("-----------删除失败！此id客户不存在！------------\n\n")
	}
}
// 改
func (this *customerView) modify() {
	fmt.Println("-------------------修改客户-------------------")
	fmt.Print("请输入要修改的客户id（-1退出）：")

	var id int
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("\n")
		return
	}
	idx := this.customerService.FindById(id)
	if idx == -1 {
		fmt.Println("-----------修改失败！此id客户不存在！------------\n\n")
		return
	}

	var name = ""
	var gender = ""
	var age = -1
	var pn = ""
	var email = ""
	customers := this.customerService.List()
	fmt.Printf("姓名（%s）：<直接回车表示不修改>", customers[idx].GetName())
	fmt.Scanln(&name)
	fmt.Printf("性别（%s）：<直接回车表示不修改>", customers[idx].GetGender())
	fmt.Scanln(&gender)
	fmt.Printf("年龄（%d）：<直接回车表示不修改>", customers[idx].GetAge())
	fmt.Scanln(&age)
	fmt.Printf("电话（%s）：<直接回车表示不修改>", customers[idx].GetPhoneNumber())
	fmt.Scanln(&pn)
	fmt.Printf("邮箱（%s）：<直接回车表示不修改>", customers[idx].GetEmail())
	fmt.Scanln(&email)

	this.customerService.ModifyCustomer(id, name, gender, age, pn, email)

	fmt.Println("-------------------修改完成-------------------\n\n")
}
// 退出时确认
func (this *customerView) exit() {
	fmt.Print("确认是否退出(Y/N)：")
	for {
		var flag string
		fmt.Scanln(&flag)
		if flag == "Y" || flag == "y" || flag == "N" || flag == "n" {
			// 退出
			if flag == "Y" || flag == "y" {
				this.loop = false
				fmt.Println("你已退出客户信息管理软件...")
				return 
			} else {		// 不退出
				this.loop = true
				fmt.Println("\n")
				return
			}
		} else {
			fmt.Print("非法输入，请确认是否退出(Y/N)：")
		}
	}
}


func main() {

	// 创建customerView实例
	cv := NewCustomerView()
	cv.MainMenu()

}