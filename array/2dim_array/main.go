package main

import (
	"fmt"
)

func main() {

	// 保存3个班、每班5名同学的成绩
	// 求出每个班的平均分，以及所有班级的平均分
	var classes [3][5]float64
	// 读取
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("请输入%d班第%d名同学的成绩：", i + 1, j + 1)
			fmt.Scanln(&classes[i][j])
		}
	}

	// 计算
	sum1, sum2 := 0.0, 0.0		// 默认是float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			sum1 += classes[i][j]
		}
		fmt.Printf("%d班平均分：%.02f\n", i + 1, sum1 / 5)
		sum2 += sum1
		sum1 = 0
	}
	fmt.Printf("3个班平均分：%.02f\n", sum2 / 15)

}
/*
80 88 89 90 75
70 90 99 86 91
76 79 77 81 60
*/