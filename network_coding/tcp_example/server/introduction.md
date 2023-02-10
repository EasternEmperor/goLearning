### 服务器
1. tcp连接方法在**net**包中
2. 监听：func Listen(net, laddr string) (Listener, error)
    - net: 必须是面向流的网络协议, "tcp", "tcp4", "tcp6"
    - 返回的Listener结构体：type Listener struct {
                            // Addr返回该接口的网络地址
                            Addr() Addr
                            // Accept等待并返回下一个连接到该接口的连接
                            Accept() (c Conn, err error)
                            // Close关闭该接口，并使任何阻塞的Accept操作都不再阻塞并返回错误
                            Close() error
                        }
        - **Conn**接口代表通用的面向流的网络连接
            - type Conn struct {
                // Read从连接中读取数据
                **Read(b []byte) (n int, err error)**
                // Write从连接中写入数据
                **Write(b []byte) (n int, err error)**
                // Close关闭连接，并使所有阻塞中的Read和Write方法不再阻塞返回错误
                Close() error
                // 返回本地网络地址
                LocalAddr() Addr
                // 返回远端网络地址
                RemoteAddr() Addr
                // 设定读写时间限制，超过该时间未读/写成功则报错
                // 等价于同时调用SetReadDeadline和SetWriteDeadline
                SetDeadline(t time.Time) error
                // 设定读时间限制
                SetReadDeadline(t time.Time) error
                // 设定写时间限制
                SetWriteDeadline(t time.Time) error
                
            }
            - Addr类型：
                - type Addr interface {
                    Network() string    // 网络号
                    String() string     // 格式化输出
                }