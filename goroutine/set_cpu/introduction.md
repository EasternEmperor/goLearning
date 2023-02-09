### go设置运行CPU数量
1. **runtime**包中的函数：
    - func NumCPU() int: 返回本地机器的逻辑CPU个数
    - func GOMAXPROCS(n int) int: 设置可同时执行的最大CPU数，并返回先前的设置。
    若传入的n<1，它就不会更改当前设置
2. 在go 1.8以后，编译器会自动让程序运行在多核上，不需要设置了