# init函数
每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用。

# 使用细节
1. Go执行流程：全局变量定义 -> init -> main
2. init函数最主要作用就是完成一些初始化工作
3. 如果main.go引入包utils，且包中go文件也有init函数，则最先执行的是引入包的init函数