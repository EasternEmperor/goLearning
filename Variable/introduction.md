### 命名规则
1. 英文字母，0-9，_
2. 数字不能作为标识符开头
3. 区分大小写
4. 标识符本身不能包含空格
5. 下划线'_'在go中是一个特殊标识符，称为“空标识符”。仅作为占位符使用，它对应的值会被忽略（如忽略某个返回值）
6. 不能以系统保留关键字作为标识符

### 注意事项
1. 包名：保持package的名字和目录名一致，且不与标准库冲突
2. 变量名、函数名和常量名，采用驼峰法
3. 如果变量名、函数名和常量名首字母大写，则可以被其他包访问（导入后），相当于是public的；如果首字母小写，则不能被其他包访问，相当于是private的。但是注意：go中是**没有**public和private等关键字

### 常量介绍
1. 常量使用**const**修饰
2. 常量在声明时，必须**初始化**
3. 常量不能修改
4. 常量只能为**bool, 数值类型(int, float等), string**类型
5. 语法：const 变量名 [type] = value   
    - type可有可无
6. 声明方式：
    - const a = 10
    - const (
        a = 10
        b = 11
    )
    - const (
        a = iota
        b
        c
    )
        - iota表示给a赋值为0，b, c值为a依次加1
7. 仍然通过**首字母大小写**来控制访问范围