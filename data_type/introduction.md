### 基本数据类型：
1. 整数类型(int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, byte)
2. 浮点类型(float32单精度, float64双精度)
3. 字符型（没有专门的字符型，使用byte来保存单个字母字符）
4. 布尔型(bool)
5. 字符串(string)
6. rune（与int32一样，表示Unicode码）

#### 基本数据类型默认值
- 整型：0
- 浮点型：0
- string：""    (空串)
- bool：false

### 派生/复杂数据类型：
1. 指针(Pointer)
2. 数组
3. 结构体(struct)       ->      替代class
4. 管道(Channel)
5. 函数(也是一种类型)
6. 切片(slice)
7. 接口(interface)
8. map

### 格式化输出
- %c
- %d
- %f
- %v：原值输出，可以用来输出结构体
- %T：数据类型
- %t：bool值"false"/"true"输出