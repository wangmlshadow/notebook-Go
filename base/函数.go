package src

//import (
//	"fmt"
//)
//
//// 函数定义的格式
//func add(a int, b int)  {
//	sum := a + b
//	fmt.Println(sum)
//
//}
//
//// 不定参函数
//// 注意类型统一
//func addAll(args ...int)  {
//	fmt.Println(args)
//	var sum int = 0
//
//	//for i := 0; i < len(args); i++ {
//	//	sum += args[i]
//	//}
//
//	// range
//	for _, data := range args {// 使用 _ 匿名变量 只接收不处理 使用i的话 在循环中必须要使用i
//		sum += data
//	}
//
//	fmt.Println(sum)
//}
//
//// 函数的返回值
////func sub(a int, b int) int {
////    return a - b
////}
//
//func sub(a int, b int) (sum int) {// 提前定义好返回值变量
//	sum = a - b
//	return
//}
//
//// 多个返回值、
//func mlt(a int, b int) (c int, d int) {
//	c = a
//	d = b
//	return
//}

//func main() {
//	a, b := 1, 2
//	add(a, b)// 形参 实参
//
//	//
//	c := 3
//	addAll(a, b, c)
//}
