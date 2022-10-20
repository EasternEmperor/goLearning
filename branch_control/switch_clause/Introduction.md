# switch特性介绍
1. Go中switch的特性：每个case语句后不需要加break
2. 基本语法：
    switch 表达式 {
    case 表达式1, 表达式2, ...:
        语句块1
        
    case 表达式3, 表达式4, ...:
        语句块2
    ...
    default:
        语句块n
    }
3. switch后表达式值的数据类型一定要和case后的相同（否则报错），注意Go中数据类型的严格区分：int8, int16, ...
4. case后面的表达式如果是常量（字面量），则要求不能重复；但是变量可以。
5. default语句不是必须的。
6. switch后的表达式也可以省略，此时可将swich...case作为一个if-else来使用
7. switch后也可以直接声明/定义一个变量，分号结束，**不推荐**
8. switch穿透：fallthrough. 如果在一个case语句后加上fallthrough，则会继续执行下一个case的语句块（不进行该case的判断）。