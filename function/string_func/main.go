package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	// 统计字符串长度：len(str)
	// Go中字符按utf-8编码，汉字为3个字节
	str1 := "hello 中国"
	fmt.Println("len(str1) = ", len(str1))

	// 字符串遍历，处理有中文的字符串：r := []rune(str)
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("r[i] = %c\n", r[i])
	}

	// 字符串转整数：n, err := strconv.Atoi(str)
	// 转换成功err为nil
	n, err := strconv.Atoi("hello")
	if err == nil {
		fmt.Println(n)
	} else {
		fmt.Println("转换错误！")
	}

	// 整数转字符串：str := strconv.Itoa(n)
	str2 := strconv.Itoa(123)
	fmt.Printf("str2 = %s, type = %T\n", str2, str2)

	// 字符串转byte切片：bytes := []byte(str)
	bytes := []byte(str1)
	fmt.Printf("bytes = %v\n", bytes)

	// byte切片转字符串：str := string(bytes)
	fmt.Printf("str1 = %s\n", string(bytes))

	// 十进制转为指定进制，返回字符串：str := strconv.FormatInt(123, 2)
	fmt.Println("123的2进制 = ", strconv.FormatInt(123, 2))
	fmt.Println("123的8进制 = ", strconv.FormatInt(123, 8))

	// 判断包含关系：b := strings.Contains(source, substr)
	fmt.Println("INT_ABC contains ABC: ", strings.Contains("INT_ABC", "ABC"))
	fmt.Println("INT_ABC contains abc: ", strings.Contains("INT_ABC", "abc"))
	
	// 计数出现次数：n := strings.Count(source, substr)
	fmt.Printf("cheese contains %d e\n", strings.Count("cheese", "e"))

	// 不区分大小写的判断相等：b := strings.EqualFold(s1, s2)
	fmt.Println("Abc equals ABC: ", strings.EqualFold("Abc", "ABC"))

	// 返回substr第一次出现在母串的下标：idx := strings.Index("INT_ABC", "ABC")
	fmt.Printf("ABC first appears at %d pos of INT_ABC\n", strings.Index("INT_ABC", "ABC"))

	// 返回substr最后一次出现在母串的下标：idx := strings.LastIndex("cheese", "e")
	fmt.Printf("e last appears at %d pos of cheese\n", strings.LastIndex("cheese", "e"))

	// 替换对应串n次，n=-1则全部替换：str := strings.Replace("cheese", "e", "E", -1)
	fmt.Println(strings.Replace("cheese", "e", "E", -1))

	// 按照给定分隔符分割字符串为数组：strs := strings.Split("Hello,world,Go", ",")
	fmt.Printf("strs = %v\n", strings.Split("Hello,world,Go", ","))

	// 转大小写：str := strings.ToLower(s) / strings.ToUpper(s)
	fmt.Println(strings.ToLower("Go"), strings.ToUpper("Go"))

	// 删除两边的空格：str := strings.TrimSpace(s)
	fmt.Printf("str = %q\n", strings.TrimSpace("  H ello World !  "))

	// 删除左右两侧包含在sub中的字符：str := strings.Trim(s, sub)
	fmt.Println(strings.Trim(" ! Hello World !  ", "! "))

	// 删除左侧/右侧包含在sub中的字符：str := strings.TrimLeft(s, sub) / strings.TrimRight(s, sub)
	fmt.Println(strings.TrimLeft("!!!!Hello world!!!!", "!"))
	fmt.Println(strings.TrimRight("!!!!Hello world!!!!", "!"))

	// 判断是否以prefix开头：b := strings.HasPrefix(s, prefix)
	fmt.Println(strings.HasPrefix(" Hello 中国", " "))

	// 判断是否以suffix结尾：b := strings.HasSuffix(s, suffix)
	fmt.Println(strings.HasSuffix(" Hello 中国", "国"))
}