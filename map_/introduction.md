### map的声明
1. 基本语法：
    - var 变量名 map[keytype]valuetype
    - keytype一般是int和string类型，也可以是其他数据类型：bool, 指针, channel, 接口, 结构体, 数组。
    slice, map, function不能作key，因为这仨不能通过==来判断
    - valuetype和keytype类型一样
    - 例: var a map[int]string
          var a map[int]map[int]string      // key是int，value是map[int]string。多重map
2. 注意：map通过var声明后是不分配空间的，所以不能像python, java一样声明后就m[0] = "...". 
   需要通过make给其分配空间才能赋值
3. make给map分配空间：
    - m = make(map[keytype]valuetype, 大小)
    - 如：
    var a map[int]string
    a = make(map[int]string, 10)
4. 注意map中的key不能重复，如果给相同的key赋值value，后赋值的会覆盖原先的
5. map中的键值对是没有顺序的，不会按照key或者value来自动排序

### map的使用
1. 先声明再make分配空间：
    var a map[int]string
    a = make(map[int]string, 10)
2. 声明时make分配空间：
    var a = make(map[int]string)    // 空间自动增长！
3. 声明时赋值：
    var a = map[string]string{
        "1" : "无用",
        "2" : "卢俊义",             // 注意最后也要有逗号！    
    }
4. map在make分配空间后，都会自动增长空间。比如1.中如果a中元素超过10，a会自动扩容！

### map的增删改查
1. 增/改：
    m[key] = value      // 如果key已存在就是改
2. 删：内置函数delete(map, key)
    delete(m, key)
    - delete删除时，如果key存在于m，则删除；如果不存在则不操作，不会报错
    - 如果想删除所有key：
        - 遍历map逐一删除
        - make一个新空间赋给map：m = make(map[keytype]valuetype, 大小)
3. 查：
    value, ok := m[key]
    if ok {     // true存在
        ...
    } else {    // false不存在
        ...
    }

### map的遍历
1. for ... range: 
    - for key, value := range m {...}
2. map的长度：len(map)

### map的切片
1. 声明map切片：slice := []map[string]string    // 即切片中每一个元素都是一个map
    - 可以用在切片元素有多个属性需要保存、并且这些属性需要快速查找的场合

### map的排序
1. Go中没有专门方法用于map键值排序
2. Go中map默认无序，每次输出的顺序都会不一样
3. Go中map的排序是先将key进行排序，然后根据key值遍历输出即可
4. **sort**包中有给切片排序的函数：
    - sort.Ints(): 将int切片按**递增**排序
    - sort.Float64s(): 将float64切片按递增排序
    - sort.Strings(): 将string切片按递增排序

### map使用细节
1. map是引用类型，作为实参传入函数，函数中修改形参会影响实参
2. map容量满后，会自动扩容，即能动态增长
3. map的value也可以是结构体，可以一个键对应更多数据