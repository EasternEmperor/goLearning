### 切片的基本介绍
1. 切片是数组的一个引用，因此切片是引用类型，在进行传递时遵循引用传递的机制
2. 切片的使用和数组类似，遍历、访问元素和求切片长度len(slice)的操作都和数组一样
3. 切片长度是可动态变化的，因此切片是一个**可以动态变化的数组**
4. 基本语法：
    var 切片名 []类型
    如：var s []int
5. cap(slice): 求切片的容量，一般是len(slice) * 2（预先多分配空间）
6. slice内存上类似结构体(struct)，首先保存第一个元素的地址，然后是len(slice)，最后是cap(slice)。
   slice通过地址来获取元素
   类似是：
        type slice struct {
            ptr *[2]int
            len int
            cap int
        }

### 切片的使用
1. slice := array[beg : end]
2. 使用**make**：
    var slice []int = make([]type, len, cap)    // cap为可选参数
    make出来的切片中元素为对应元素类型的默认元素
        - 通过make创建的切片对应的数组是由make底层维护，对外透明，只能通过slice访问
3. 指定具体数组，原理类似make：
    var slice = []int{1, 2, 3}

### 切片的遍历
1. 常规for遍历：
    for i := 0; i < len(slice); i++ {...}
2. for ... range 遍历：
    for index, value := range slice {...}

### 内置函数append
1. 用数组初始化切片时，可以和python一样简写
2. slice也可以用slice来初始化（切片可以继续切片）：
    var slice2 = slice1[beg : end]
3. 内置函数**append**:
    - 用append给切片追加元素：
        slice = append(slice, n1, n2, ...)  // 注意append会将追加完毕的slice作为返回值，要手动接收！
    - 追加切片：    
        slice = append(slice, slice2...)    // 注意追加的切片后有三个点"..."！
    - 扩容/append时，实际上是新开辟足够大的空间将原slice复制过去再append。旧空间会被回收，且如果赋值给
    同名变量地址也不会变化（见main.go）

### 内置函数copy
1. copy(para1, para2)   参数的数据类型是**切片**
2. 这里拷贝是**深拷贝**，拷贝完两个切片互不影响