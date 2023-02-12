### redis连接池
1. 通过Go操作redis还可以使用连接池，以节省临时建立redis连接的时间，提高效率
2. 实现初始化一定数量的连接，放入到连接池
3. 当Go需要操作redis时，直接从连接池中取出连接即可

### 核心代码
var pool *redis.Pool
pool = &redis.Pool{
    MaxIdle: 8,         // 最大空闲连接数
    MaxActive: 0,       // 表示和数据库的最大连接数，0表示没有限制（则看MaxIdle）
    IdleTimeout: 100,   // 最大空闲时间，当一个连接在IdleTimeout内没有被使用，就会重新变为Idle
    // 初始化连接的代码，连接哪个ip的redis
    Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", "localhost:6379")
    }
}
c := pool.Get()     // 从连接池中取出一个连接
pool.Close()        // 关闭连接池，关闭后就不能再用pool来取连接了