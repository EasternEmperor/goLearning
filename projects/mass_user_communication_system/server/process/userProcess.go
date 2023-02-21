package process

import (
	"fmt"
	"net"
	"encoding/json"
	"projects/mass_user_communication_system/common/message"
	"projects/mass_user_communication_system/server/utils"
	"projects/mass_user_communication_system/server/model"
	"errors"
)

type UserProcess struct {
	// 字段
	Conn net.Conn
	// 添加字段，以表示是哪个用户的UserProcess
	UserId int
}

// 当用户状态改变后，要将状态发送给所有在线用户的客户端，使它们显示的状态改变
func (this *UserProcess) NotifyOtherOnlineUsers(status int) (err error) {

	// 遍历onlineUsers，对每一个在线用户发送当前用户的状态
	for id, up := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == this.UserId {
			continue
		}
		// 服务端给对应的客户端发送状态更新
		err = up.NotifyClientStatus(this.UserId, status)
		if err != nil {
			err = fmt.Errorf("对%d用户状态更新失败！err = %v\n", id, err)
		}
	}
	return
}
// 服务端向客户端发送要更新userId的状态
func (this *UserProcess) NotifyClientStatus(userId int, status int) (err error) {
	// 声明Message实例以传输
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	
	// 声明NotifyUserStatusMes实例
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = status

	// 序列化notifyClientStatus赋给mes.Data
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		err = fmt.Errorf("NotifyClientStatus对notifyClientStatus序列化失败！err = %v", err)
		return
	}
	mes.Data = string(data)

	// 序列化mes以发送
	data, err = json.Marshal(mes)
	if err != nil {
		err = fmt.Errorf("NotifyClientStatus对mes序列化失败！err")
		return 
	}

	// 创建Transfer实例发送data
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		err = fmt.Errorf("NotifyClientStatus发送mes失败！err =", err)
	}

	return
}

// 处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1. 获取登录信息
	var loginMes message.LoginMes
	// 将mes.Data反序列化为结构体实例
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		err = errors.New("serverProcessLogin反序列化mes.Data失败！")
		return 
	}
	
	// 2. 声明LoginResMes变量
	var loginResMes message.LoginResMes

	// 3. 判断用户名密码
	// 使用MyUserDAO去数据库中验证
	user, err := model.MyUserDAO.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_PWD {
			// 密码错误
			loginResMes.Code = 500
			// Error()类似String()，可以将错误信息转为字符串
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_NOTEXISTS {
			// 用户不存在
			loginResMes.Code = 404
			loginResMes.Error = err.Error()
		}

	} else {
		// 登录成功
		loginResMes.Code = 200
		// 添加UserProcess新字段UserId
		this.UserId = loginMes.UserId
		fmt.Println("login UserId =", loginMes.UserId)
		// 将该用户添加到在线用户列表中
		userMgr.AddOnlineUser(this)
		// 通知其他用户：我上线了
		err = this.NotifyOtherOnlineUsers(message.UserOnline)
		// 同时将所有在线用户的id放入loginResMes的用户id切片中，以返回给客户端
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登录成功！")
	}
	// 输出出现的错误
	fmt.Println(err)

	// 4. 创建新的Message实例作为返回消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	data, err := json.Marshal(loginResMes)
	if err != nil {
		err = errors.New("serverProcessLogin序列化loginResMes失败！")
		return 
	}
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，调用writePkg()发送
	pkg, err := json.Marshal(resMes)
	if err != nil {
		err = errors.New("serverProcessLogin序列化resMes失败！")
		return 
	}
	// 创建Transfer实例，调用其中的write方法
	var tf = &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(pkg)
	return 
}

// 处理离线请求
func (this *UserProcess) ServerProcessLogout(mes *message.Message) (err error) {
	// 1. 获取登出信息
	var logoutMes message.LogoutMes
	// 将mes.Data反序列化为结构体实例
	err = json.Unmarshal([]byte(mes.Data), &logoutMes)
	if err != nil {
		err = errors.New("serverProcessLogout反序列化mes.Data失败！")
		return 
	}

	// 2. 补充该up的UserId字段
	this.UserId = logoutMes.UserId
	
	// 3. 通知其他用户：我离线了
	err = this.NotifyOtherOnlineUsers(logoutMes.Status)

	// 4. 删除userMgr中的该用户
	userMgr.DelOnlineUser(this)

	return
}

// 处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 1. 获取注册信息
	var registerMes message.RegisterMes
	// 将mes.Data反序列化为结构体实例
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		err = errors.New("serverProcessRegister反序列化mes.Data失败！")
		return 
	}
	
	// 2. 声明RegisterResMes变量
	var registerResMes message.RegisterResMes

	// 3. 判断用户Id
	// 使用MyUserDAO去数据库中验证
	err = model.MyUserDAO.Register(registerMes.User.UserId, registerMes.User.UserPwd, 
								   registerMes.User.UserName)
	if err != nil {

		if err == model.ERROR_USER_EXISTS {
			// 用户Id已被注册
			registerResMes.Code = 500
			// Error()类似String()，可以将错误信息转为字符串
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 404
			registerResMes.Error = err.Error()
		}

	} else {
		// 注册成功
		registerResMes.Code = 200
		fmt.Println("用户", registerMes.User.UserId, "注册成功！")
	}

	// 4. 创建新的Message实例作为返回消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	data, err := json.Marshal(registerResMes)
	if err != nil {
		err = errors.New("serverProcessRegister序列化registerResMes失败！")
		return 
	}
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，调用writePkg()发送
	pkg, err := json.Marshal(resMes)
	if err != nil {
		err = errors.New("serverProcessRegister序列化resMes失败！")
		return 
	}
	// 创建Transfer实例，调用其中的write方法
	var tf = &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(pkg)
	return 
}