fmt包中有读取用户输入的函数：
1. func Scanf(format string, a ...interface{}) (n int, err error)
    - 类似C/C++的scanf()
2. func Scanln(a ...interface{}) (n int, err error)
    - 把成功读取的空白分隔的值保存进成功传递给本函数的参数。在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。