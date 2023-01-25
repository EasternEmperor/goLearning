1. defer: 在函数中，程序员经常需要创建资源（数据库连接、锁等），为了**在函数执行完后，
及时释放资源**，Go中提供defer（延时机制）
2. 当执行到defer时，defer之后的语句暂时不执行，会将该语句压入一个栈中，当函数执行完
毕后，再从defer栈，按照先入后出的方式弹出栈，并执行
3. 在defer将语句入栈时，也会将涉及到的值也同时入栈。因此如果后面对这些值进行了修改，
并不影响defer语句的值
4. defer可以让open和close同时成对出现，不容易因为疏忽而漏掉close
    func test() {
        // 打开关闭文件
        file = openfile(文件名)
        defer file.close()
        // 其他代码
    }