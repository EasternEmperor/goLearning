### 处理静态文件
1. 对于HTML页面中的css及js等静态文件，需要使用**net/http**包下的方法来处理：
    - func StripPrefix(prefix string, h Handler) Handler: **返回一个处理器**，该处理器会将
    请求的URL.Path字段给定**前缀prefix去除再交由h处理**。如果URL.Path中没有给定前缀会**返回404**
    - func FileServer(root FileSystem) Handler: 返回一个使用FileSystem接口root提供文件访问
    服务的HTTP处理器
        - 如：http.Handle("/", http.FileServer(http.Dir("/tmp")))
            - 使用操作系统的FileSystem接口实现，可使用**http.Dir**
2. FileSystem**接口**实现了**对一系列命名文件的访问**。文件路径的**分隔符为'/'**，**不管主机操作系统**
   的惯例如何
    ```go
    type FileSystem interface {
        Open(name string) (File, error)
    }
    ```
3. Dir使用限制到指定目录树的本地文件系统**实现了http.FileSystem接口**: `type Dir string`
    - func (d Dir) Open(name string) (File, error)