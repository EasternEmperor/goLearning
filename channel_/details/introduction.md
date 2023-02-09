### 管道使用细节
1. 默认情况下管道是双向的，可读可写，也可声明为只读/只写管道：
    - 只写：**var 管道名 chan<- 数据类型**
        var intChan chan<- int
    - 只读：**var 管道名 chan-> 数据类型**
        var intChan2 <-chan int
    - 只读/只写管道一般声明在**函数形参**上，以免函数内部对该管道进行**不想要**的读写
2. 在无法判定管道是否关闭的情况下对管道读写，可以使用select语句：
    - select {
        case v := <-chan1:
            ...
        case v := <-chan2:
            ...
        ...
        default:
            ...     // 添加读不到时的逻辑
    }
3. goroutine中使用recover，解决协程中可能出现的panic，导致整个程序崩溃的问题
    - defer + recover