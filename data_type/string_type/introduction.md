1. string底层是一个byte数组，其类似结构体：
    type string struct {
        ptr *[]byte
        len int
    }
2. 底层的byte数组对外透明
3. string本身是不可变的！即不能通过`str[0] = 'o'`来改变字符串内容
4. 如果需要修改字符串，需要先将string转为[]byte / []rune （两类切片） -> 修改 -> 重写为string