### 搭建简单服务器
1. 调用处理器函数：func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    - 注册一个处理器handler和对应的模式pattern
    - pattern表示**地址后的什么类型串需要用这个handler来处理**
    - HandleFunc如果传入的是函数，则会自动将其转换为一个处理器
2. 处理器接口：
    ```go
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
    ```
    - 如果实现了该接口就可以直接使用`http.Handle(pattern string, handler Handler)`函数来使用
        - http.Handler('/', &MyHandler)
3. 监听TCP地址和端口：func ListenAndServe(addr string, handler Handler) error
    - handler为设置**多路复用器**，如果为nil则使用默认的多路复用器**DefaultServeMux**
4. type HandlerFunc: 该类型是一个适配器，相当于处理器函数。通过**类型转换**我们可以将普通函数作为http
   处理器函数使用HandlerFunc(f)
    - type HandlerFunc func(ResponseWriter, *Request)
5. 监听也可以自己创建Server实例调用方法ListenAndServe()：
    ```go
    type Server struct {
        Addr            string          // 监听的TCP地址
        Handler         Handler         // 调用的处理器，如为nil会调用http.DefaultServeMux
        ReadTimeout     time.Duration   // 读最大持续时间
        WriteTimeout    time.Duration   // 写最大持续时间
        MaxHeaderBytes  int             // 请求的最大头域长，如为0则DefaultMaxHeaderBytes
        TLSConfig       *tls.Config     // 可选的TLS配置，用于ListenAndServeTLS方法
        ...
    }
    ```
    - func (srv *Server) ListenAndServe() error: 监听srv中的Addr，调用Handler来处理
6. 多路复用器：
    ```go
    type ServeMux struct {}
    ```
    - 创建新的多路复用器：func NewServeMux() *ServeMux
    - 使用ServeMux实例也可以调用处理器解决请求：Handle(), HandleFunc(), ServeHTTP()