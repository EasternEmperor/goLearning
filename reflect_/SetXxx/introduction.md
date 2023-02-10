1. 通过reflect.Value变量的SetXxx()来修改变量值时，值类型要传入指针才能修改
    - 如：
    var str = "Mary"
    fs := reflect.ValueOf(&str)     // 传入指针
    fs.Elem().SetString("Jack")     // 同时要使用Elem()方法获取封装的Value