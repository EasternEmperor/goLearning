# 使用细节
1. 循环条件返回布尔值
2. for 循环的另一种方法：
    for 循环判断条件 {
        // 循环体
    }
3. 第三种写法：配合break使用。等价for ; ; ; 
    for {
        // 循环体
    }
4. for index, val := range string/数组等 {

}
    - for range 还可以用来直接遍历汉字字符串
5. 下标访问方式要输出汉字字符要将字符串转为切片类型rune，再遍历

# 注：Go中没有while和do...while！
- 用for代替：
    for {
        if 循环结束条件 {
            break
        }
        循环体
    }