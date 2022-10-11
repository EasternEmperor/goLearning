### 基本数据类型转string
1. 通过fmt.Sprintf()。
- Sprintf()会将其中格式化的数据转为string类型返回
    - eg: str := fmt.Sprintf("%d", num)
2. 通过strconv包里的函数
- func FormatInt(i int64, base int) string
    - 传入int64类型的整型，base为要转成的进制
- func Itoa(i int) string
    - 整型转string，只有一个参数，更简单
- func FormatFloat(f float64, fmt byte, prec, bitSize int) string
    - f：传入float64类型的双精度浮点；
    - fmt：格式。'f'：为普通10进制小数格式；'b'：二进制指数；'e'：十进制指数；'E'：同'e'，但指数标记用E
    - prec：小数位数；若传入-1，则使用最少的必要小数位数来转换f
    - bitSize：表示f来源为float32还是float64，会根据此项进行取舍
- func FormatBool(b bool) string
    - b：待转换的bool值