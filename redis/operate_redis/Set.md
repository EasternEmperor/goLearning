### Set基本介绍
1. redis的Set是String类型的无序集合，底层是HashTable数据结构
2. Set可以存放很多字符串元素，是**无序**存放的，而且元素的值**不能重复**

### Set基本操作
1. sadd key value1 value2 ...
2. smembers key: 显示key集合中所有的元素
3. sismember key value: 判断value是否在key集合中
4. srem key value: 删除指定值