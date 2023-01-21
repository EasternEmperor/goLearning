### 时间和日期相关函数
1. now := time.Now(): 当前时间
2. t.Year() / t.Month() / t.Day() / t.Hour() / t.Minute() / t.Second(): 获取time类型的年月日时分秒
3. t.Format("2006-01-02 15:04:05"): 格式化time类型为指定格式字符串
                                    注意：格式可以变，但是这个时间不能变！是固定的！
4. 时间的常量：在time中定义了时间间隔类型duration
    const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
    )

    在程序中用于获得指定时间单位的时间，比如100毫秒：100 * time.millisecond
5. 结合time.Sleep(t)可以对程序进行定时休眠工作
6. t.Unix(): 1970.1.1 00:00:00 到当前时间的秒数
7. t.UnixNano(): 1970.1.1 00:00:00 到当前时间的纳秒数