package main

import (
	"fmt"
)

// 冒泡排序
func BubbleSort(arr *[5]int) {

	flag := false
	for i := 0; i < len(*arr) - 1; i++ {

		for j := 0; j < len(*arr) - 1 -i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = tmp
				flag = true
			}
		}
		if flag == false {
			break
		} else {
			flag = false
		}

	}

}

// 快速排序
func QuickSort(arr *[5]int, beg int, end int) {
	if beg < end {
		pos := Partion(arr, beg, end)
		QuickSort(arr, beg, pos - 1)
		QuickSort(arr, pos + 1, end)
	}
}
func Partion(arr *[5]int, beg int, end int) int {
	pivot := (*arr)[beg]	// 选定基准元素
	for beg < end {
		// 比基准小的元素移到左边
		for (*arr)[end] >= pivot && beg < end {
			end--;
		}
		(*arr)[beg] = (*arr)[end]
		// 比基准大的元素移到右边
		for (*arr)[beg] <= pivot && beg < end {
			beg++;
		}
		(*arr)[end] = (*arr)[beg]
	}
	(*arr)[beg] = pivot
	return beg
}

func main() {

	var arr = [...]int{23, 80, 13, 69, 50}
	BubbleSort(&arr)
	fmt.Println("冒泡排序后，arr =", arr)

	var arr2 = [...]int{23, 80, 13, 69, 50}
	QuickSort(&arr2, 0, len(arr2) - 1)
	fmt.Println("快速排序后，arr2 =", arr2)
	
}