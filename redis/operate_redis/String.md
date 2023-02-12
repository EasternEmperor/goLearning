### String的CRUD
1. set key value（如果key存在就是修改，不存在就是添加）
2. get key
3. del key

### String操作的使用细节
1. setex: 可以设置类型及其失效时间(set expire), setex key time value
    - setex key1 5 beijing: 则beijing值在5s后会自动被删除
2. mset: 同时设置多个key-value对
    - mset key1 value1 key2 value2 ...
3. mget: 同时获取多个key对应的value
    - mget key1 key2 ...