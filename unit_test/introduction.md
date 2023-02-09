### 单元测试基本介绍
1. Go语言自带一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试
2. 通过单元测试解决以下问题：
    - 确保函数可运行，结果正确
    - 确保代码性能良好
    - 及时发现代码的逻辑错误，性能测试使代码在高并发时也稳定

### 编写说明
1. 导入testing包，testing提供对Go包的自动化测试，通过"go test"命令，能够自动执行如下形式
   的任何函数：func TestXxx(*testing.T)
    - Xxx可以是任何字母、数字和字符串（但第一个X不能是[a-z]），用于识别测试程序
2. 在TestXxx()中调用要测试的函数
    - t.Fatalf("...")和t.Logf("...")类似fmt.Printf()，可以格式化输出日志

### testing框架工作
1. testing框架将xxx_test.go文件引入(import)，然后自动创建一个main函数来调用其中的
   TestXxx()函数
2. xxx_test.go会引入要测试的函数xxx
3. xxx_test.go和要测试的函数xxx应在同一个包内

### 单元测试细节说明
1. 测试文件必须以**_test.go**结尾
2. 测试用例函数必须以Test开头，一般就以Test+要测试的函数名，如TestUPPer
3. 测试用例函数的形参类型**必须是***testing.T
4. 一个测试用例文件可以包含**多个**测试用例函数
5. 运行测试用例指令：
    - go test: 如果运行正确，不输出日志；错误才输出
    - go test -v: 无论正确和错误都会输出日志
6. 当出现错误时，使用t.Fatalf()输出错误信息，并退出测试函数
7. t.Logf()可以输出相应日志
8. 测试单个文件，一定要带上被测试的源文件！
    - go test -v cal_test.go cal.go
9. 测试单个方法：
    - go test -v -test.run TestAddUpper