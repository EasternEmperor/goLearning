package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"io/ioutil"
)

// 缓冲区读取文件
func BufRead() {
	// 打开
	f, err := os.Open("introduction.md")		// 相对路径
	if err != nil {
		// 打开失败
		fmt.Println("打开文件失败！", err)
		return 
	}
	// defer关闭文件
	defer f.Close()

	// 读取内容
	reader := bufio.NewReader(f)
	for {
		// 读取一行
		str, err2 := reader.ReadString('\n')
		if err2 == io.EOF {
			fmt.Println("\n------读取结束！-------\n\n")
			break
		}
		fmt.Println(str)
	}
}

// 一次性读取：io/ioutil.ReadFile(name string)
func OneRead() {
	content, err := ioutil.ReadFile("introduction.md")
	if err != nil {
		fmt.Println("读取文件出错：err = ", err)
		return 
	}
	// 读取成功输出
	// 注意返回值为[]byte, 要把它转为string
	fmt.Println(string(content))
	fmt.Println("\n------读取结束！-------\n\n")
}


// 写
// 缓冲写
func BufWrite() {
	// 打开
	file, err := os.OpenFile("test.txt", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建文件失败！err =", err)
		return 
	}
	// 关闭
	defer file.Close()

	// 创建缓冲写writer
	writer := bufio.NewWriter(file)
	// 将缓冲区内的数据写入文件
	defer writer.Flush()

	str := "Hello World!\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

}
// 一次写
func OneWrite() {
	// 读取
	data, err := ioutil.ReadFile("introduction.md")
	if err != nil {
		fmt.Println("读取文件出错！err =", err)
		return
	}

	// 写入
	err = ioutil.WriteFile("test.txt", data, 0666)
	if err != nil {
		fmt.Println("写入文件出错！err =", err)
	}


}

// 拷贝文件io.Copy()
func CopyFile(des string, src string) (written int64, err error) {

	// 打开待读文件
	read, err := os.Open(src)
	if err != nil {
		fmt.Println("读取文件失败！err =", err)
		return 
	}
	// 关闭
	defer read.Close()
	reader := bufio.NewReader(read)

	// 打开待写文件
	write, err2 := os.OpenFile(des, os.O_CREATE | os.O_WRONLY, 0666)
	if err2 != nil {
		fmt.Println("读取文件失败！err =", err2)
		return 
	}
	// 关闭
	defer write.Close()
	writer := bufio.NewWriter(write)

	// io.Copy()拷贝
	return io.Copy(writer, reader)

}


func main() {

	// 用ReadString()读取文件
	BufRead()

	// io/ioutil.ReadFile()
	OneRead()
	
	// WriteString()写文件
	BufWrite()

	// WriteFile()写文件
	OneWrite()

	// 拷贝文件
	CopyFile("abc.txt", "test.txt")
}