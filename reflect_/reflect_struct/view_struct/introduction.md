### 反射结构体使用的函数
1. reflect.Type和reflect.Value结构体的方法（设实例分别为t, v）：
    - **v.Kind()**: 获得v封装变量的类别，如int, int32, float64, struct, map等

    - **func (v Value) NumField() int**: 获取v封装结构体的字段个数

    - **func (v Value) Field(i int) Value**: 以Value封装返回结构体的第i个字段，
    如果v的Kind不是结构体或i越界会panic

    - **t.Field(i).Tag.Get(key)**: 返回结构体字段tag设置的key对应的别名，如\`json:"name"\`
        - func (t Type) Field(i int) **StructField**: 返回Struct类型的第i个
    字段的类型
        - type StructField struct {...}, 其中Type表示字段的类型，Tag表示字段设置
        的标签
        - func (StructTag) Get(key string) string: 由Tag调用，返回结构体定义
        给字段设置的tag的key对应的别名

    - **func (v Value) NumMethod() int**: 返回v持有值的方法个数
    
    - v.Method(i).Call(params): 调用v持有值的第i个方法，传入对应个数的参数params为
    []Value类型；**没有参数则传入nil**
        - **func (v Value) Method(i int) Value**: 返回v持有类型的第i个方法封装到的Value类型
            - 注意：这里给方法排序不是按定义的先后，而是按方法名称的字符串(ASCII)排序
        - **func (v Value) Call(in []Value) []Value**: v为封装的函数时调用Call表示使用
        该封装的函数，通过in传递参数，返回值为[]Value类型