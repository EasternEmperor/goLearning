package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	"projects/mass_user_communication_system/common/message"
	"errors"
)

func Login(userId int, userPwd string) (err error) {
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
	var loginMes = message.LoginMes{userId, userPwd}
	// 4. 对loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		str := fmt.Sprintf("loginMes序列化失败！err =", err)
		err = errors.New(str)
		return 
	}
	// 5. data赋给mes.Data
	mes.Data = data
	// 6. 将mes进行序列化，准备传输
	tMes, err := json.Marshal(mes)
	if err != nil {
		str := fmt.Sprintf("mes序列化失败！err =", err)
		err = errors.New(str)
		return 
	}
	// 7. 将tMes进行传输
	// 7.1 首先传输tMes的长度
	var pkgLen uint32 = uint32(len(tMes))
	var bytes [4]byte
	// 将pkgLen转为大端字节序表示的byte数组
	binary.BigEndian.PutUint32(bytes[0 : 4], pkgLen)
	// 发送
	n, err := conn.Write(bytes[ : 4])
	if n != 4 || err != nil {
		str := fmt.Sprintf("发送包长度失败！err =", err)
		err = errors.New(str)
		return 
	}
	// 7.2 发送包
	n, err = conn.Write([]byte(tMes))
	if err != nil {
		str := fmt.Sprintf("发送包失败！err =", err)
		err = errors.New(str)
		return 
	}

	// 8. 接收服务端返回数据，判断是否成功登录
	resMes, err := readPkg(conn)
	if err != nil {
		str := fmt.Sprintf("login接收服务端数据失败！err =", err)
		err = errors.New(str)
		return 
	}
	// 反序列化resMes.Data
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		str := fmt.Sprintf("login反序列化resMes.Data失败！err =", err)
		err = errors.New(str)
		return 
	}
	// 判断登陆成功与否
	if loginResMes.Code == 200 {
		fmt.Println("用户登录成功！")
	} else {
		fmt.Println(loginResMes.Error)
	}

	return 

}