package message

// 定义消息类型Type常量
const (
	LoginMesType 			= "LoginMes"
	LoginResMesType 		= "LoginResMes"
	RegisterMesType 		= "RegisterMes"
	RegisterResMesType 		= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType				= "SmsMes"
	LogoutMesType			= "LogoutMes"
)

// 定义用户状态Status常量
const (
	UserOnline 	= iota		// 自增
	UserOffline
)

// C/S通信传输消息的格式
type Message struct {
	Type string	`json:"type"`	// 消息类型
	Data string	`json:"data"`	// 消息
}

// 先定义两个登录时用的消息
// 客户端
type LoginMes struct {
	UserId int `json:"userId"`			// id
	UserPwd string `json:"userPwd"`		// 密码
	UserName string `json:"userName"`	// 用户名
}
// 服务端
type LoginResMes struct {
	Code int `json:"code"`		// 返回状态码: 500表示用户未注册，200表示登录成功，404表示用户不存在
	Error string `json:"error"`	// 返回错误信息
	// 新字段：返回一个在线用户id的切片
	UserIds []int				// 在线用户id切片
}

// 定义两个注册时用的消息
// 客户端
type RegisterMes struct {
	// 用户信息
	User User	`json:"user"`
}
// 服务端
type RegisterResMes struct {
	Code int		`json:"code"`	// 400表示用户已被注册，200表示注册成功
	Error string	`json:"error"`	// 错误信息
}

// 定义用户状态变化时，告知其他客户端以更新它们那里保存的状态
type NotifyUserStatusMes struct {
	UserId int	`json:"userId"`	// 用户id
	Status int	`json:"status"`	// 用户状态
}

// 增加短消息类型
// SmsMes：发送
type SmsMes struct {
	Content string	`json:"content"`
	User		// 匿名结构体，即继承
}

// 增加登出消息类型
type LogoutMes struct {
	UserId int	`json:"userId"`
	Status int	`json:"status"`
}