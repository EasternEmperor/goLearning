### 命令行参数
1. Go中使用os包的string切片：Args来保存所有命令行参数
2. 获取命令行参数：os.Args[i]

### flag包解析命令行参数
1. 在命令行输入中有时会指定参数名：mysql -u root -p 123. 此时使用os.Args就不方便
   解析，使用flag包内函数
    - func (f *FlagSet) IntVar(p *int, name string, value int, usage string):
        IntVar用指定的名称name、默认值value、使用信息usage注册一个int类型flag，并
        将flag的值保存到p指向的变量
    - 同理还有func StringVar(), func Boolvar(), func Float64Var()...
    - 使用方法：定义要接收的参数名如：flag.IntVar(&port, "p", 3366, "端口号，默认为3366")
        - 此时运行exe时输入命令行参数-p，就会自动被port接收。如果没有-p，port默认值就是
        3366
2. 在绑定参数名后，需要调用flag.Parse()方法来将命令行参数转换到变量里保存！
    - flag.Parse()
    - 之后便可以通过变量访问这些参数