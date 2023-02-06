## 客户信息管理系统

### 项目需求说明
1. 菜单
2. 增删改查客户信息

### 程序框架图
1. customer_view.go: 界面(view)
    - 可以显示界面
    - 可以接受用户输入
    - 根据用户输入，调用customer_service的方法完成客户的增删改查
2. customer_service.go: 控制器(control)
    - 完成对用户的各种操作
    - 对客户的增删改查操作
    - 创建一个customer切片以保存多个客户信息
3. customer.go: 模型(model)
    - 表示一个客户
    - 客户的各种字段