### 管道基本介绍
1. Channel本质是一个队列数据结构
2. 数据**先进先出**
3. **线程安全**
4. channel是有类型的，即一个string管道只能存放string类型

### 管道基本使用
1. 声明：var 变量名 chan 数据类型
    - var intChan chan int
    - var mapChan chan map[string]int
    - var perChan chan Person
    - var perchan2 chan *Person
2. 管道是**引用类型**
3. 因为是引用类型，所以必须make初始化分配空间后才能存放数据
    - make(chan type, 容量)
4. 写入数据： **管道名<- 数据**
    - intChan<- num
5. 读取数据： **变量 := <-管道名**
    - num := <-intChan
    - 抛弃数据（不接收）：<-intChan
6. 管道容量由make时指定，管道内数据不能超过该容量！**（不会自动增长）**
7. 实现向管道中存放任意数据类型，则声明时数据类型设为interface{}
    - var allChan chan interface{}
    - 输出时要使用**类型断言**才能获取到结构体变量的字段
8. 写管道和读管道**频率不一致**，无所谓；如果一个管道**只有写，没有读**，则会阻塞报错

### 管道的遍历和关闭
1. 内置函数close(管道变量)
    - 关闭后不能再向管道中写数据，但是读取仍正常
2. channel支持for-range遍历：
    - 若管道**未关闭**，可能会出现**deadlock**错误（死锁，因为会一直等待写入）
    - 若管道**已关闭**，可以正常遍历
    - for v := range 管道名 {
        ...
    }
3. 读取管道同时判断是否关闭：
    - v, ok := <-intChan
    - 如果管道**已关闭**，且读到管道**末尾**，则ok为false
