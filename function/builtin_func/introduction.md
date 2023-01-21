### 内置函数
1. len(v): 求长度，如string, map, channel, array, slice
2. new(type): 用来分配内存，主要用来分配值类型，比如int, int32, float, struct...返回指针
3. make(type): 用来分配内存，主要用来分配引用类型，比如channel, map, slice