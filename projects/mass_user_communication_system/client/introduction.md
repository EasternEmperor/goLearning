# 客户端框架设计与改进
1. **main.go**:
    1. 显示第一级菜单
    2. 根据用户输入调用相应的处理器
2. **smsProcess.go**:
    1. 处理短消息相关逻辑
    2. 私聊
    3. 群发
3. **userProcess.go**:
    1. 处理和用户相关的业务
    2. 登录
    3. 注册等
4. **server.go**:
    1. 显示登录成功界面
    2. 保持和服务器通讯【启动协程】
    3. 当读取到服务器发送的消息后，就显示在界面上
5. **utils.go**:
    1. 提供常用的工具函数和结构体


### login.go传输用户输入
1. 连接到服务端
2. 构造Message结构体实例，首先要给定Type类型:
    type Message struct {
        Type string
        Data string
    }
3. 创建LoginMes结构体实例：
    type LoginMes struct {
        UserId int
        UserPwd string
        UserName string     // 可有可无
    }
4. 将LoginMes序列化为string类型(data)
5. 将data赋值给mes.Data准备传输
6. 将mes进行序列化：传输时都是用string类型
7. 此时，tMes是我们要传输的消息，按照设计需要先把tMes的长度传给服务端
    - 传输len(data)：获取到tMes长度，将该长度转为byte数组（用4个字节表示）：
        var pgkLen uint32
        pkgLen = uint32(len(tMes))
        var bytes [4]byte
        binary.BigEndian.Putuint32(bytes[0 : 4], pkgLen)
        // 发送
        n, err := conn.Write(bytes)
        - encoding/binary包中有对网络传输字节序列进行处理的方法，其中BigEndian代表大端字节序，
        LittelEndian代表小端字节序。PutUint32([]byte, uint32)，同样的也有PutUint16/PutUint64
    
### 客户端发送消息
1. 新增一个消息结构体SmMes
2. 新增一个model CurUser
3. 在smsProcess.go 增加相应的方法SendGroupMes，发送群聊消息