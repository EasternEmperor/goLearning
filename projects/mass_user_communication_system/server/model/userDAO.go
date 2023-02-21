package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"errors"
)

// 定义一个全局UserDAO实例以供服务端使用
var MyUserDAO *UserDAO

// 定义一个UserDAO结构体，用于完成对用户信息的增删改查
type UserDAO struct {
	// 用redis线程池作为字段，可以方便地连接数据库
	pool *redis.Pool
}

// 因为pool是私有字段，要通过工厂模式来创建
func NewUserDAO(pool *redis.Pool) (userDAO *UserDAO) {
	userDAO = &UserDAO{pool}
	return
}

// 1. 查找用户，根据用户Id返回一个User实例+err
func (this *UserDAO) getUserId(conn redis.Conn, id int) (user *User, err error) {
	// 通过特定id在redis中查找
	res, err := redis.String(conn.Do("Hget", "users", id))
	if err != nil {
		// redis.ErrNil表示未找到
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	// 找到则将其反序列化然后返回
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		err = errors.New("userDAO.getUserId反序列化redis串失败！")
		return
	}
	return
}

// 处理登录请求
func (this *UserDAO) Login(userId int, userPwd string) (user *User, err error) {
	// 通过this连接redis数据库
	conn := this.pool.Get()
	// 记得关闭
	defer conn.Close()
	// 获取数据库中的user信息
	user, err = this.getUserId(conn, userId)
	if err != nil {
		return
	}
	// 获取到用户信息，比对密码
	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return 
	}
	return 
}

// 处理注册请求
func (this *UserDAO) Register(userId int, userPwd string, userName string) (err error) {
	// 通过this连接redis数据库
	conn := this.pool.Get()
	// 记得关闭
	defer conn.Close()
	// 获取数据库中的user信息
	user, err := this.getUserId(conn, userId)
	// 没有报错说明该用户ID已被注册，无法注册
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// 可以注册，将用户信息写入
	// 创建User实例序列化后存入数据库
	user = &User{userId, userPwd, userName}
	data, err := json.Marshal(*user)
	if err != nil {
		str := fmt.Sprintf("userDAO.Register序列化user失败！err = %v", err)
		err = errors.New(str) 
		return
	}
	// 存入数据库
	_, err = conn.Do("Hset", "users", userId, string(data))
	if err != nil {
		str := fmt.Sprintf("userDAO.Register将user写入数据库失败！err = %v", err)
		err = errors.New(str)
		return
	}
	return 
}