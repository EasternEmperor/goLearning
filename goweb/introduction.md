### ServeMux的定义
1. 多路复用器ServeMux结构如下：
    ``` go
    type ServeMux struct {
        mu      sync.RWMutex            // 锁，处理并发
        m       map[string]muxEntry     // 路由规则，一个url对应一个mux实体
        hosts   bool                    // 是否在任意的规则中带有host信息
    }
    ```
    - muxEntry结构：
        ```go
        type muxEntry struct {
            explicit bool       // 是否精确匹配
            h        Handler    // 这个路由表达式对应的handler
            pattern  string     // 匹配字符串
        }
        ```
    - 路由器实现的ServeHTTP分发请求：
        ```go
        func (mu *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
            if r.RequestURI == "*" {
                w.Header().set("Connection", "close")
                w.WriteHeader(StatusBadRequest)
                return 
            }
            h, _ := mux.Handler(r)
            h.ServeHTTP(w, r)
        }
        ```
        - 如上所示路由器接收到请求之后，如果是 * 那么关闭链接，不然调用 mux.Handler(r) 返回对应设置路由的处理 Handler，然后执行 h.ServeHTTP(w, r)
2. Handler定义：
    ```go 
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)     // 路由实现器
    }
    ```
3. HandleFunc定义：
    ```go
    type HandleFunc func(ResponseWriter, *Request)
    // ServeHTTP calls f(w, r)
    func (f HandleFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }
    ```