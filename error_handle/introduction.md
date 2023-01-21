### Go的错误处理机制
1. Go中不支持try...catch...finally
2. Go中引入defer, panic, recover
3. 处理过程为：Go程序抛出一个异常，然后在defer中通过recover捕获这个异常，并正常处理

### 自定义错误
1. error.New("错误说明")，会返回一个error类型的值，表示一个错误
2. panic内置函数接收一个interface{} 类型的值作为参数，可以接收error类型的变量，**输出错误信息，并终止程序**
