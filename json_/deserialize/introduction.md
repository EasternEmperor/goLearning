### json反序列化
1. json反序列化：将json字符串反序列化成对应的数据类型（如结构体，map，切片）的操作
2. encoding/json包中的Unmarshal():
    - func Unmarshal(data []byte, v interface{}) error
    - data: 传入要转换的json串
    - v: 要转成的数据类型变量
3. 使用Unmarshal()反序列化map或切片等时，不需要对传入Unmarshal()的变量进行make，
   因为Unmarshal()内部封装了对这些变量的make操作
4. 序列化为结构体变量时：
    - 如果序列串中有结构体**没有的字段**，不会报错，会**忽略**该字段
    - 反序列化时，会忽略字段名的大小写。

### 反序列化注意事项
1. 在反序列化json字符串时，要确保**反序列化的串的原本数据类型**和**序列化前的数据类型**一致