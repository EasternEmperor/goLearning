# 服务端框架设计与改进
1. **main.go**：
    1. 监听
    2. 等待客户端的连接
    3. 初始化的工作
2. **processor.go**:
    1. 根据客户端的请求，调用相应的处理器，完成对应的任务
3. **smsProcess.go**:
    1. 处理和短消息相关的请求
    2. 群聊
    3. 点对点聊天
4. **userProcess.go**:
    1. 处理和用户相关的请求
    2. 登录
    3. 注册
    4. 注销
    5. 用户列表管理
5. **userMgr.go**:
    1. 维护用户在线列表
6. **utils.go**：
    1. 一些常用的工具函数、结构体
    2. 提供常用的方法和函数

### 接收客户端数据readPkg(conn net.Conn)
1. 接收传输的包长度，encoding/binary包中的：
    - pkgLen := binary.BigEndian.Uint32(buf[0 : 4])
    - Uint32([]byte) uint32
2. 再读取

### 处理登录请求
1. mes.Type记录了此次请求的类型，在process()里判断，进入serverProcessLogin()方法处理
2. 步骤:
    1. 从mes中取出mes.Data（string类型），并直接**反序列化为LoginMes结构体实例**
    2. 声明一个**LoginResMes**结构体实例，作为返回给客户端的数据
    3. 判断用户名密码是否正确和匹配，是则给LoginResMes的**Code**赋值200，代表登录成功；
    否则Code赋值500，登陆失败。同时根据成功与否给LoginResMes的**Error**赋值
    4. 创建一个Message结构体实例resMes作为返回给客户端的消息，resMes.Data赋值为序列化的
    loginResMes变量，resMes.Type为LoginResMesType
    5. 将resMes进行序列化进行发送，调用writePkg()函数

### 发送数据writePkg(conn net.Conn, data []byte)
1. 先发送数据长度
2. 再发送数据本身