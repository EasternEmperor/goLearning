1. 随机数包："math/rand"
2. rand.Intn(n int): 生成[0, n)之间的随机数
3. rand.Seed(n): 配合Seed()使用改变种子可以改变生成的随机数，否则是Seed的默认种子