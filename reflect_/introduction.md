### 反射的基本介绍
1. 反射可以在程序运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)
2. 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
3. 通过反射可以修改变量的值，可以调用关联的方法
4. 使用反射需要引入reflect包

### 反射重要的函数和概念
1. reflect.TypeOf(i interface{}) Type: 获取变量的类型，返回reflect.Type类型
2. reflect.ValueOf(i interface{}) Value: 获取变量的值，返回reflect.Value类型
   具体值的Value
3. 变量、interface{}和reflect.Value是可以相互转换的
    - 变量转换过程： **变量 --(函数传参)--> interface{}**
                    **i interface{} --(reflect.ValueOf(i))--> reflect.Value**
    - Value转回过程：**v reflect.Value --(v.Interface())--> interface{}**
                    **i interface{} --(类型断言)--> 变量**
4. reflect.Kind: 表示Type类型按类别的具体分类，如Int, Int8, Float64, Struct, Map, 
   Slice, Array...
    - value.Kind()和type.Kind(): 返回其持有值的类别。其中value是reflect.Value变量，
    type是reflect.Type变量

### 反射的注意事项
1. 使用reflect.Value变量的方法来获取持有的值时，如v.Int(), v.Float()，要求数据类型匹配，
   比如v原本类型是int，那么就只能用v.Int()来获取值，否则报panic
2. 通过反射来修改变量，需要调用reflect.Value的SetXxx()方法。但是注意，因为要修改
   值，一般传入的是指针，所以不能直接用v.SetXxx()，需要用: v.Elem().SetXxx()
    - v.Elem(): 返回v持有的接口保管的值的Value封装，或者v**持有的指针**指向的值的
    Value封装
    - v.Elem().SetXxx(): 由Elem()获得了Value封装，再调用SetXxx()来设置新值