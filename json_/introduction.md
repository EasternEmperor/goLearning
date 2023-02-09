### json基本介绍
1. 概述：JSON(JavaScript Object Notation)是一种**轻量级数据交换格式**
2. json易于机器解析和生成，并有效地提升**网络传输效率**，通常程序在网络传输时会先将数据
   （结构体、map等）**序列化**为json字符串，到接收方得到字符串后，再**反序列化**恢复成
   原来的数据类型
3. 应用场景：
    - C/S模式：Client/Server, tcp编程
    - B/S模式：Browser/Server, Web编程

### json数据格式说明
1. json使用**键值对**来保存数据：
    - [
        {"key1" : value1, "key2" : value2, "key3" : [value3_1, value3_2]},
        {"key1" : value1, "key2" : value2, "key3" : [value3_1, value3_2]}
    ]
2. www.json.cn

### json的序列化
1. json序列化：将**有key-value结构**的数据类型（如结构体，map，切片）序列化为
   json字符串的操作
2. func Marshal(v interface{}) ([]byte, error): 返回v的json编码