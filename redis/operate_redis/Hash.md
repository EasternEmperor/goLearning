### Hash介绍
1. redis Hash是一个**键值对集合**，类比Go中的var user1 map[string]string
2. redis Hash是一个String类型的field和value的映射表，hash特别适合用来存储
   对象（结构体实例）

### Hash基本操作
1. 增/改：hset HashName field value
    - hset user1 name "Smith"
      hset user1 age 30
      hset user1 sal 1000
2. 查：hget HashName field
    - hget user1 name

### Hash操作注意事项
1. hset/hmset
2. hget/hmget
3. hgetall: 可以获取**一个HashName的所有key-value**
    - hgetall user1
4. hdel
5. **一次性设置Hash的多个字段**：hmset
    - hmset user1 name "Smith" age 45
6. **一次性获取Hash的多个字段**：hmget
    - hmget user1 name age
7. hlen: 获取Hash的字段个数
8. hexists key field: 判断Hash中是否有field字段，返回0/1代表false/true
    - hexists user1 name