# 习题
声明结构体Student
type Student struct {
    Name string
    Age int
    score float64
}
// 将Student的切片，按score从小到大排序

## 解法
1. 使用sort包中的sort方法最高效：
    - func Sort(data Interface)
    - 要使[]Student能使用Sort方法，则[]Student要实现Interface接口
2. type Interface interface {
    // Len()返回元素个数
    Len() int
    // Less()用于判断下标i和下标j的元素是否要交换
    // ture则不交换(i < j)，false则交换(i > j)
    Less(i, j int) bool
    // Swap方法交换i和j下标上的元素
    Swap(i, j int)
}
3. sort.Sort()底层使用的是quickSort()，即快排，O(nlogn)