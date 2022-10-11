### string转基本数据类型
使用strconv包里的函数：
1. func ParseInt(s string, base int, bitSize int) (i int64, err error)
    - s：待转的字符串
    - base：要转成的进制
    - bitSize：对返回值进行大小限制，0、8、16、32、64 分别代表在 int、int8、int16、int32、int64 的范围内来返回转换值（溢出处理）
    - 注意该函数有两个返回值，接收时需进行处理
    - 注意该函数只能返回int64类型的整型
2. func ParseBool(str string) (value bool, err error)
    - 转换bool
3. func ParseFloat(s string, bitSize int) (f float64, err error)
    - bitSize：同样对返回值进行大小限制，32、64 分别代表在 float32、float64 的范围内来返回转换值
    - 注意该函数只能返回float64类型的浮点型
4. func Atoi(s string) (i int, err error)
    - ParseInt(s, 10, 0)的简写版，参数少，更方便

#### 注意事项
若使用strconv的函数，传入不对应的string，会将其转换为该转换类型的零值。如ParseInt输入"hello"，则会直接将其转为0