### 获取请求URL
1. Request请求体结构：
    ```go
    type Request struct {
        // Method指定HTTP方法(Get, Post, Put...)。对客户端""表示Get
        Method string
        // URL在服务端表示被请求的URI，在客户端表示要访问的URL
        URL *url.URL
        // 表示HTTP请求的头域
        Header Header
        // Body是请求的主体
        Body io.ReadCloser
        // 请求体长度
        ContentLength int64
        // PostForm是解析好的Post或Put的表单数据
        // 本字段只有调用ParseForm后才有效
        PostForm url.Values
        ...
    }
    ```
    - Body的类型：ReadCloser实现了Reader和Closer接口，其中有两个方法用于读取和关闭连接：
        - func Read(p []byte) (n int, err error)
        - func Close() error
    - func (r *Request) FormValue(key string) string: 返回key为键查询**r.Form**字段得到
    结果[]string切片的**第一个值**
    - func (r *Request) PostFormValue(key string) string: 返回key为键查询**r.PostForm**
    字段得到结果[]string切片的**第一个值**
        - 注意：FormValue和PostFormValue都**将ParseForm()步骤封装到内部**了
2. *url.URL请求体结构：
    ```go
    type URL struct {
        Scheme string
        User *Userinfo      // 用户名和密码信息
        Host string         // host或者host:port
        Path string         // path
        RawQuery string     // 编码后的查询字符串，为'?'之后的信息
        ...
    }
    ```
    - URL类型代表一个解析后的URL. 基本格式：`scheme://[userinfo@]host/path[?query][#fragment]`,
    或者：`scheme:opaque[?query][#fragment]`
3. type Header代表HTTP头域的键值对：`type Header map[string][]string`
    - func (h Header) Get(key string) string: Get返回**键对应的第一个值**，不存在返回""。
        - 注意！这里返回的不是切片
    - func (h Header) Set(key, value string): Set添加键值对到h，如键已存在则会用**只有新值**
    一个元素的切片**取代旧值切片**
4. 服务端相应客户端使用ResponseWriter的方法写：
    ```go
    type ResponseWriter interface {
        // 返回的Header类型会被WriterHeader()发送，
        // 因此应当在调用WriterHeader()之前修改Header
        Header() Header
        // 发送HTTP回复的头域和状态码。
        // 如果没有显式调用，则第一次调用Write()时会触发隐式调用WriteHeader(http.statusOK)
        // 因此一般用于发送错误码
        WriteHeader(int)
        // 发送回复的数据
        Write([]byte) (int, error)
    }
    ```
    - Write(data): 可以传输html文本（转换为string再转为[]byte来传输）