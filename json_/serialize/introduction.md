### json的序列化
1. encoding/json包中的Marshal()函数：
    - func Marshal(v interface{}) ([]byte, error)
2. 对基本数据类型的序列化，其实就是将其转换为了字符串

### struct序列化须知
1. **小写**字母为首的字段**不会被序列化**！因为不是公开的！
2. 使用**tag**指定字段进行序列化后的名字
    - type Monster struct {
        Name string    \`json:"name"\`
    }
    - 注意！！！json:后**不能有空格**，应**直接**跟"..."
