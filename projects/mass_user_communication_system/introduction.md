# 客户端
## 登录
1. 接收输入的id和pwd
2. 发送id和pwd给服务器
3. 接收到服务端返回的结果
4. 判断登录成功还是失败，并显示相应页面

### 关键的问题是如何组织发送的数据？
1. 设计消息协议
    - **消息类型**：
        type Message struct {
            Type string
            Data string
        }
        LoginMes str {
            userId int
            userPwd string
        } **序列化**作为传输Message的Data
2. **发送的流程**：
    1. 先创建Message结构体变量
    2. mes.Type = 登录消息类型
    3. mes.Data = 登陆消息的内容（序列化后）
    4. **对mes进行序列化**
    5. 在网络传输中，最怕丢包，因此发送数据可以分两步：
        - 先发送给服务器mes的长度（有多少个字节n）
        - 再发送mes本身

# 服务端
## 登录
1. 比对id，pwd【goroutine】
2. 比较
3. 返回结果

### 接收数据流程
1. 接收到客户端的数据长度n
2. 再接收到消息mes本身
3. 根据n判断mes长度是否符合（是否丢包）
4. 如果不相等，就有**纠错协议**
5. 相等，则对mes --**（反序列化）**--> Message实例
6. 取出message.Type判断是登录还是发送消息
7. 取出message.Data --**（反序列化）**--> LoginMes实例
8. 取出loginMes.userId和loginMes.userPwd和数据库中的比较
9. 根据比较结果组建Message实例返回给客户端
    - 返回时Data可以定义服务端自己的结构体序列化的string类型
    - type LoginResMes str {
        code int        // 成功/失败
        error string    // 出错
    } **序列化**

# 编译
1. 服务端指令：go build -o server.exe projects/mass_user_communication_system/server/main
2. 客户端指令：go build -o client.exe projects/mass_user_communication_system/client/main