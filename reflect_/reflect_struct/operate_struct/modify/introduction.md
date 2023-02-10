1. v := reflect.ValueOf(x): 如果x是指针类型，则v也会是指针类型(*reflect.Value).
   要使其变为直接的reflect.Value实例：v = v.Elem()
2. func (v Value) FieldByName(name string) Value: 按名查找字段
    - 一般连用为: v.FieldByName(name).SetXxx(newVal), 用于给字段赋新值
3. 注意：能被反射机制获取到的字段一定是**公开的**，即首字母是大写的！