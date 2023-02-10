1. func New(typ Type) Value: 返回持有指向类型为Typ的新申请的零值的指针，**返回值也是指针**
    - 一般用于反射中创建结构体实例(var st reflect.Type)：
        var model *User
        st := reflect.TypeOf(model)    // 注意此时的st也为指针
        st = st.Elem()      // 先将st从指针变为Value实例
        elem := reflect.New(st)         // 创建新对象，注意！这里返回的仍为指针，要修改字段得调用Elem()方法
2. 创建实例后给字段赋值：
    - elem.FieldByName(name).SetXxx(newVal)