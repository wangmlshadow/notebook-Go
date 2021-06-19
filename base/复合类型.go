package src
//
//import "fmt"
//
//// 函数以数组作为参数
//func test(arr [10]int)  {
//	arr[2] = 10
//	fmt.Println(arr)
//}
//
//// 函数返回数组
//func test2(value int) [10]int {
//	var arr[10]int
//	for i := 0; i < 10; i++ {
//		arr[i] = value
//	}
//	return arr
//}
//
//func main() {
//	// 数组的定义
//	var arr [10]int
//
//	for i := 0; i < 10; i++ {
//		arr[i] = i
//	}
//
//	fmt.Println(arr)
//
//	// 初始化
//	var arr2 [10]int = [10]int{0, 1, 2, 3, 4, 5, 6 ,7 , 8, 9}
//	fmt.Println(len(arr2))
//
//	// 集合的遍历
//	for _, v := range arr2 {
//		fmt.Println(v)
//	}
//
//	// 数组在定义时可以只初始化一部分
//	var arr3[10]int = [10]int{0, 1, 2, 3}
//	fmt.Println(arr3)
//
//	// 使用自动化类型推导初始化数组
//	arr4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(arr4)
//
//	// 自动推导数组长度
//	arr5 := [...]int{0, 1, 2, 3 , 4}
//	fmt.Println(arr5)
//
//	// 数组是一个常量 不允许辅助 表示数组地址
//	//arr5 = 10
//
//	// 指定数组下标进行初始化数据
//	arr6 := [10]int{0, 1, 2, 9:9}
//	//fmt.Println(arr6)
//	// 数组作为函数参数时 传递数组名
//	test(arr6)
//	// 数组作为函数参数 是值传递
//	fmt.Println(arr6)
//}
