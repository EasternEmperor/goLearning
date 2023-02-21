package process

import (
	"fmt"
)

// 声明全局变量，在服务端维护所有在线用户
var userMgr *UserMgr 

// 定义一个结构体保存在线用户id和对应的UserProcess实例
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}	
}

// 添加用户
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}
// 删除用户
func (this *UserMgr) DelOnlineUser(up *UserProcess) {
	delete(this.onlineUsers, up.UserId)
}
// 返回当前在线的所有用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}
// 根据id返回对应值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 从onlineUsers中取出
	up, ok := this.onlineUsers[userId]
	// 不存在
	if !ok {
		err = fmt.Errorf("用户%d不存在或者不在线~", userId)
		return 
	}
	return 
}