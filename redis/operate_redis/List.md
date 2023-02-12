### List基本介绍
1. 列表是简单的字符串列表，按照插入顺序排序。可以添加元素到列表的头部或尾部
2. List本质上是一个**链表**，List的元素是**有序的**，**元素值可以重复**

### List基本操作
1. lpush key value1 value2 ...: 从List的左边**依次插入**value1...
    - 所以注意：最后放完，List第0个元素是valuen（最后那个元素）
2. rpush: 与lpush类似，不过是从右边插入
3. lrange key start end: 输出key列表中[start, end]间的元素（左闭右闭）
    - 这里end也可以用负数，**-1表示倒数第一个元素**，-2倒数第二个，以此类推
4. lpop key: 从左边弹出key列表中的元素
    - 弹出后原列表中就没有这个元素了！
    - 如果列表的元素**全部弹出**，那么这个列表key会**自动消失**
5. rpop key: 从右边弹出key列表中的元素
6. del: 删除列表
7. lindex key i: 获取列表索引i上的元素（从左到右）
    - lindex cities 0
8. llen key: 获取列表长度