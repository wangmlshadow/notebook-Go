package src

//import "fmt"
//
//// 切片作为函数参数
//func test(s []int) {
//	s[0] = 88// 切片传递地址 源切片数据也会被改变
//	fmt.Println(s)
//}
//
//func main() {
//	// 数组定义 var 数组名 [元素个数]数据类型
//	// 切片定义 var 切片名 []数据类型
//
//	//var s []int
//	//fmt.Println(s)
//
//	// 自动推导类型创建切片
//	s := make([]int, 5)// 设置长度为5
//	s[0] = 1
//	s[1] = 2
//
//	//s[6] = 7// error 越界
//	// 使用append添加元素
//	fmt.Println(len(s))
//	s = append(s, 6, 7 , 8, 9)
//	fmt.Println(len(s))
//	fmt.Println(s)
//	// 查看容量
//	fmt.Println(cap(s))
//
//	// go的切片不是python的切片 可以看作C++的vector
//
//	// 推荐使用切片而不是数组
//
//	// 切片的截取 类似于python的切片
//	fmt.Println(s[2:])
//	fmt.Println(s[:4])
//	fmt.Println(s[2:4])
//	fmt.Println(s[0:2:4])// low = 0 height = 2 max = 4 cap = max - low
//
//	slice := s[2:4]// 切片数据仍然指向原始的s 修改slice的话s也会被修改
//	slice[0] = 9
//	slice[1] = 9
//	fmt.Println(slice)
//	fmt.Println(s)
//
//	s2 := make([]int, len(s))
//	copy(s2, s)// 拷贝操作 深拷贝 s2需要有足够的空间存放拷贝过来的数据
//	s2[0] = 100
//	fmt.Println(s)
//	fmt.Println(s2)
//	// 也可因copy切片的截取部分
//
//	test(s)
//	fmt.Println(s)
//}
