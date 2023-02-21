package utils

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	"projects/mass_user_communication_system/common/message"
	"errors"
)

// OOP思想，将每个包都用对象组织起来
// 这个包主要是工具类，用来传输/读取消息
type Transfer struct {
	// 连接
	Conn net.Conn
	// 传输时使用到的缓冲
	Buf [8096]byte
}

// 读取传输的数据：包长度+包
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// 定义字节切片来接收数据
	//buf := make([]byte, 8096)
	fmt.Println("接收来自客户端的数据...")
	// 接收包长度
	_, err = this.Conn.Read(this.Buf[ : 4])
	if err != nil {
		err = fmt.Errorf("接收包头失败！err = ", err)
		return 
	}
	// 转为长度
	pkgLen := binary.BigEndian.Uint32(this.Buf[ : 4])
	
	// 接收包
	n, err := this.Conn.Read(this.Buf[ : pkgLen])
	if n != int(pkgLen) || err != nil {
		str := fmt.Sprintf("接受包出错！err = %v\n-----接收到%d个字节，实际%d个字节-----\n", err, n, pkgLen)
		err = errors.New(str)
		return 
	}

	// 接收到的数据是字节序列，要将其反序列化为Message结构体实例
	// 注意传入mes的地址：&mes
	err = json.Unmarshal(this.Buf[ : pkgLen], &mes)
	if err != nil {
		err = errors.New("mes反序列化失败！")
		return 
	}

	return
}

// 传输返回的数据给客户端
func (this *Transfer) WritePkg(resMes []byte) (err error) {
	// 1. 发送数据长度
	var pkgLen uint32 = uint32(len(resMes))
	//var bytes [4]byte
	binary.BigEndian.PutUint32(this.Buf[ : 4], pkgLen)
	// 发送
	n, err := this.Conn.Write(this.Buf[ : 4])
	if n != 4 || err != nil {
		err = errors.New("writePkg发送包头出错！")
		return 
	}

	// 2. 发送包
	n, err = this.Conn.Write(resMes)
	if n != int(pkgLen) || err != nil {
		err = errors.New("writePkg发送包出错！")
		return 
	}
	return 

}