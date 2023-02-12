### Go连接到redis
1. 通过**github.com/garyburd/redigo/redis**包中的Dial()方法：
    - func Dial(network, address string, options ...DialOption) (Conn, error):
    用来连接Go程序和redis数据库: conn, err := redis.Dial("tcp", "127.0.0.1:6379")
2. func Do(cmd string, args ...interface{}) (interface{}, error): 操作redis数据库执行
   cmd里的指令
    - conn.Do("set", "name", "Tom"): 向redis数据库中添加一对name:Tom的String型值
        - cmd是redis中操作数据的指令，如set, get, mset, mget...
    - 返回的interface{}是获取到的redis中的数据，要转为能操作的Go中类型用该包的内置函数：
        - redis.String()/redis.Strings()（多个String类型的, 返回切片）
        - redis.Float64()
        - redis.Int64()
        - ...
3. 给数据设置有效时间: c.Do("expire", key, i)    // i是int型，单位是s