### Go中字符串相关的函数
1. len(str): 统计字符串的长度
2. r := []rune(str): 字符串遍历，同时处理有中文的问题
3. n, err := strconv.Atoi(str): 字符串转整数
4. n := strconv.Itoa(n): 整数转字符串
5. bytes := []byte(str): 字符串转byte切片（utf-8/ascii码值）
6. str := string(bytes): byte切片转字符串
7. str := strconv.FormatInt(n, 进制): 将十进制的n转为指定进制
8. b := strings.Contains(s, substr): 判断substr是否在s中
9. n := strings.Count(s, sep): 统计一个字符串中有几个指定的子串
10. b := strings.EqualFold("abc", "ABc"): 不区分大小写的字符串比较
11. idx := strings.Index(s1, s2): 返回s2在s1中第一次出现的下标，如果没有返回-1
12. idx := strings.LastIndex(s1, substr): 返回substr在s1中最后一次出现的下标，没有则返回-1
13. str := strings.Replace(s, source, re, n): 用re来替换s中的source子串，n为按顺序替换个数，n=-1时表示全部替换
                                              注意：s本身不会变化，替换后的作为返回值
14. strs := strings.Split(source, sep): 将source以sep为分割符分割成数组
15. str := strings.ToLower(s) / str := strings.ToUpper(s): 将字符串转为全小写/大写
                                                           注意：s本身不变化
16. str := strings.TrimSpace(s): 将s左右两边的空格全部删除
17. str := strings.Trim(s, sub): 将s左右两边的包含在sub中的字符全部删除
18. str := strings.TrimLeft(s, sub) / strings.TrimRight(s, sub): 删除左边/右边的包含在sub中的字符
19. b := strings.HasPrefix(s, prefix): 判断s是否以prefix开头
20. b := strings.HasSuffix(s, suffix): 判断s是否以suffix结尾