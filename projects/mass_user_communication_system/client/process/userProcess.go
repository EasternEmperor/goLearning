package process

import (
	"fmt"
	"net"
	"encoding/json"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/client/utils"
)

type UserProcess struct {
	// 暂时不用字段
}

// 处理登录请求
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 1. 创建连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("创建连接失败！err =", err)
		return 
	}
	// 及时关闭
	defer conn.Close()

	// 2. 创建Message实例
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3. 创建LoginMes实例
	var loginMes = message.LoginMes{
		UserId : userId, 
		UserPwd : userPwd,
	}
	// 4. 对loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		err = fmt.Errorf("loginMes序列化失败！err = %v", err)
		return 
	}
	// 5. data赋给mes.Data
	mes.Data = string(data)
	// 6. 将mes进行序列化，准备传输
	tMes, err := json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("mes序列化失败！err = %v", err)
		return 
	}
	// 7. 将tMes进行传输
	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg([]byte(tMes))
	if err != nil {
		return 
	}

	// 8. 接收服务端返回数据，判断是否成功登录
	// 创建Transfer实例用于收取数据
	// tf := &utils.Transfer{conn}
	resMes, err := tf.ReadPkg()
	if err != nil {
		err = fmt.Errorf("login接收服务端数据失败！err = %v", err)
		return 
	}
	// 反序列化resMes.Data
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		err = fmt.Errorf("login反序列化resMes.Data失败！err = %v", err)
		return 
	}
	// 判断登陆成功与否
	if loginResMes.Code == 200 {
		fmt.Println("用户登录成功！")

		// 初始化CurUser变量，保存当前用户的信息
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// 登录成功打印在线用户列表
		fmt.Println("当前在线的用户：")
		for _, id := range loginResMes.UserIds {
			fmt.Printf("%d\t", id)

			// 完成客户端的onlineUsers初始化
			user := &message.User{
				UserId : id, 
				UserStatus : message.UserOnline,
			}
			onlineUsers[id] = user
		}
		fmt.Println("\n快去找TA聊天吧~")

		// 启动一个协程保持和服务端的通信
		go processMes(conn)

		// 显示登陆菜单
		var flag = true
		for flag {
			flag = ShowLoginMenu()
		}
	} else {
		err = fmt.Errorf(loginResMes.Error)
		return
	}

	return 

}

// 处理登出/掉线，告知服务器
func (this *UserProcess) Logout(userId int) (err error) {
	// 1. 创建Message实例
	var mes message.Message
	mes.Type = message.LogoutMesType

	// 2. 创建LogoutMes实例
	var logoutMes = message.LogoutMes{userId, message.UserOffline}

	// 3. 序列化logoutMes作为mes的Data
	data, err := json.Marshal(logoutMes)
	if err != nil {
		err = fmt.Errorf("Logout序列化logoutMes失败！err = %v", err)
		return
	}
	mes.Data = string(data)

	// 4. 序列化mes准备传输
	pkg, err := json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("Logout序列化mes失败！err = %v", err)
		return
	}

	// 5. 创建Transfer实例传输data
	var tf = &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(pkg)
	if err != nil {
		return
	}

	return
}

// 处理注册请求
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1. 创建连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("创建连接失败！err =", err)
		return 
	}
	// 及时关闭
	defer conn.Close()

	// 2. 创建Message实例
	var mes message.Message
	mes.Type = message.RegisterMesType
	// 3. 创建LoginMes实例
	var user = message.User{
		UserId : userId, 
		UserPwd : userPwd, 
		UserName : userName,
	}
	var RegisterMes = message.RegisterMes{user}
	// 4. 对loginMes序列化
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		err = fmt.Errorf("RegisterMes序列化失败！err = %v", err)
		return 
	}
	// 5. data赋给mes.Data
	mes.Data = string(data)
	// 6. 将mes进行序列化，准备传输
	tMes, err := json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("mes序列化失败！err = %v", err)
		return 
	}
	// 7. 将tMes进行传输
	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg([]byte(tMes))
	if err != nil {
		return 
	}

	// 8. 接收服务端返回数据，判断是否成功登录
	// 创建Transfer实例用于收取数据
	// tf := &utils.Transfer{conn}
	resMes, err := tf.ReadPkg()
	if err != nil {
		err = fmt.Errorf("Register接收服务端数据失败！err = %v", err)
		return 
	}
	// 反序列化resMes.Data
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(resMes.Data), &registerResMes)
	if err != nil {
		err = fmt.Errorf("Register反序列化resMes.Data失败！err = %v", err)
		return 
	}
	// 判断登陆成功与否
	if registerResMes.Code == 200 {
		fmt.Println("==>用户注册成功！请尝试登录您的新账号~~")
	} else {
		fmt.Println(registerResMes.Error)
	}

	return 
}