### 模板引擎
1. 使用html/template包中的结构体和方法：
    ```go
    type Template struct {
        // 底层的模板解析树
        Tree *parse.Tree
        ...
    }
    ```
    - `func Must(t *Template, err error) *Template`: 用于**包装返回(\*Template, err)**
    的函数/方法调用，会在**err不为nil时panic**，一般**用于变量初始化**
    - `func ParseFiles(filenames ...string) (*Template, error)`: 创建一个模板，并**解析filenames
    指定**的文件里的**模板定义**，返回的模板的名字是**第一个文件的名字**（不含扩展名），内容为解析后的
    第一个文件的内容
    - `func ParseGlob(pattern string) (*Template, error)`: 创建一个模板，并**解析匹配pattern**
    的文件里的模板定义。返回的模板的名字是第一个文件的名字（不含扩展名），内容为解析后的第一个文件
    的内容。和ParseFiles可以达成相同的效果
    - `func (t *Template) Execute(wr io.Writer, data interface{}) error`: 将解析好的template应用
    到data上，并将输出**写入wr**。模板可以安全地并发执行
    - `func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error`: 
    类似Execute，但是可以**指定要输出的template的名字**（在ParseFiles/ParseGlob中输入了多个文件时，有多个
    模板，可以通过名字来使用ExecuteTemplate来输出）