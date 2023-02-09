1. 用4个协程求[1, 8000]的素数个数
    - putNum(): 往管道intChan中按顺序放[1, 8000]的数
    - primeNum(): 从intChan中取数，判断它是否为素数，是则放入primeChan中。同时
    在发现intChan关闭且取完后向boolChan中放入true表示本协程结束
    - main(): 创建协程，通过boolChan中是否有4个true来判断是否结束