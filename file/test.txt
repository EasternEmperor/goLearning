### 文件基本介绍
1. 文件在程序中是以**流**的形式来操作的
    - **流**：数据在数据源（文件）和程序（内存）之间的**路径**
    - **输入流**：文件 -> 程序
    - **输出流**：程序 -> 文件
2. Go中**os包**的**File结构体**封装所有文件相关操作
    - type File struct {...}

### 文件操作
1. func Open(name string) (file *File, err error): 打开文件以读取(O_RDONLY)
2. func Close(f *File) err: 关闭文件f
3. 概念说明：file叫法（以下表意都一样，只是叫法不同）
    - file对象
    - file指针
    - file文件句柄

### 读文件
1. 读入缓冲区：使用os.Open, file.Close, bufio.NewReader(), reader.ReadString()函数和方法
    - bufio包中的func NewReader(rd io.Reader) *Reader: 创建一个具有默认大小缓冲(4096字节)、
    从rd读取的*Reader
        - 注意：这里传入的io.Reader是一个接口，而File结构体实现了该接口，因此可以直接传入
        File实例
    - func (b *Reader) ReaderString(delim byte) (line string, err error): 读取直到第一次
    遇到**delim**字符，以string返回读取数据，若有错误则报错err
2. 一次读入：io/ioutil包中的ReadFile()方法
    - func ReadFile(name string) ([]byte, error)
    - 使用ReadFile()不需要打开（被封装在函数内部了）
    - 只适合文件较小的情况
    - 注意：返回值为[]byte！

### 写文件
1. func OpenFile(name string, flag int, perm FileMode) (file *File, err error): 更一般的打开文件方式
    - flag: 文件打开模式（可以组合使用）:
        {
            O_RDONLY
            O_WRONLY
            O_RDWR
            O_APPEND        // 追加
            O_CREATE        // 不存在则创建
            O_TRUNC         // 打开时清空文件
        }
    - perm: 权限控制(linux)
2. 带缓存的写：os.OpenFile, file.Close, bufio.NewWriter, writer.WriterString, writer.Flush函数和方法
    - 类似读取，这里是带缓存的写，通过bufio.NewWriter创建的writer，先用WriterString写入
    缓存，最后需要调用Flush方法将写入缓存的数据写入文件
    - 最后不要忘记调用Flush方法写入文件！
3. 一次写入：io/ioutil包中的WriteFile()方法
    - func WriteFile(name string, data []byte, perm os.FileMode) error
    - FileMode仅在linux下有用
    - 传入data为[]byte类型

### 判断文件是否存在
1. os.Stat(name string) (fi FileInfo, err error): err为nil说明文件存在，否则再使用os.IsNotExist(err)
   判断该err是否表示一个文件或目录不存在，true则说明文件/目录不存在
2. 总结：
    func FileExists(name string) (bool, error) {
        _, err := os.Stat(name)
        if err == nil {
            return true, nil        // 文件存在
        }
        if os.IsNotExist(err) {
            return false, nil       // 文件不存在
        }
        return false, err           // 是文件不存在以外的报错
    }

### 拷贝文件
1. io包下的Copy()函数：
    - func Copy(dst Writer, src Reader) (written int64, err error)
    - 返回的written为拷贝的字节数