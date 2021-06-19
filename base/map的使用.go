package src

//import "fmt"
//
//// map作为函数参数 是地址传递
//func  test(m map[int]string)  {
//	m[1] = "AAAAAA"
//}
//
//func main() {
//	// map key 必须是基本数据类型
//    //var m map[int]string
//    m := make(map[int]string, 1)// map自动扩容
//    fmt.Println(m)
//    m[1] = "asdasd"
//    m[2] = "safasd"
//    m[9] = "asfkasd"
//    fmt.Println(m)
//    fmt.Println(m[2])
//
//    for k, v := range m {
//    	fmt.Println(k, v)
//	}
//
//	// 判断是否存在key
//	v1, ok1 := m[1]
//	v5, ok5 := m[5]
//
//	if ok1 {
//		fmt.Println(v1)
//	} else {
//		fmt.Println("key 1 not existed")
//	}
//
//	if ok5 {
//		fmt.Println(v5)
//	} else {
//		fmt.Println("key 5 not existed")
//	}
//
//	// 删除map中的元素
//
//	fmt.Println(m)
//	delete(m, 1)
//	// delete删除map元素时 表示的时若可key存在就删除 不存在也不会报错
//	delete(m, 5)
//	fmt.Println(m)
//
//	test(m)
//
//	fmt.Println(m)
//}
