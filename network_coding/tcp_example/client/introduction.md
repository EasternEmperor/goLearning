### 客户端
1. 调用net包中的Dial()函数：
    - func Dial(network, address string) (Conn, error): 在网络network上连接
    地址address，返回Conn接口。可用的network类型有："tcp", "tcp4", "tcp6", 
    "udp", "udp4", "udp6", "ip"...对于address，如果是tcp/udp协议，则格式为
    "ip:端口号"；如果是ip协议，则只需要"ip"即可
2. 向服务器发送消息：
    - conn, err := net.Dial(net, addr)  // 请求建立连接
    - reader := bufio.NewReader(os.stdin)   // 从键盘获取输入
    - data, err := reader.ReadString('\n')  // 读取键盘输入
    - n, err2 := conn.Write([]byte(data))   // 传输