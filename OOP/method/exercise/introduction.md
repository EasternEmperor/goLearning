### 面向对象编程步骤
1. 声明结构体，确定结构体名
2. 编写结构体的字段
3. 编写结构体的方法

### 创建结构体变量时指定字段值
type Stu struct {
    name string
    age int
}
1. 方式一：
    - var s1 Stu = Stu{"Tom", 10}
    - s2 := Stu{
        name: "Tom", 
        age: 10,
        }
2. 方式二：
    - var s3 *Stu = &Stu{"Tom", 10}
    - s4 := &Stu{
        name: "Tom",
        age: 10,
        }