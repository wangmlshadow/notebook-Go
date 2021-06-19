package src

//import "fmt"
//
//// 指针作为函数参数
//func test(p *int) {
//    *p = 99
//    fmt.Println(*p)
//}
//
//// 切片指针作为函数参数
//func test2(p *[]int) {
//    *p = append(*p, 999, 999)
//}
//
//type student struct {
//    id int
//    name string
//}
//
//func main() {
//    var a int = 10
//    // 定义整形指针变量 指向a的地址
//    // 与C++类似
//    var pa *int = &a
//    fmt.Println(pa)
//    *pa = 20
//    fmt.Println(a)
//
//    // 为指针变量创建一块内存空间
//    var p1 *int
//    // 堆
//    p1 = new(int)
//    fmt.Println(*p1)
//    fmt.Println(p1)
//
//    test(&a)
//    test(p1)
//    fmt.Println(a)
//    fmt.Println(*p1)
//
//    // 数组指针
//    var arr [5]int = [5]int{1, 2, 3, 4, 5}
//    fmt.Println(arr)
//    fmt.Printf("%p\n", &arr)
//
//    // 数组指针定义时给出的数组大小需要和赋值给他的数组大小一样
//    // var parr *[5]int
//    // parr = &arr
//    parr := &arr
//    fmt.Println(*parr)
//    fmt.Println(parr)
//    fmt.Printf("%T\n", parr)
//
//    // 切片指针
//    var slice []int = []int{1, 2, 3, 4, 5}
//    ppslice := &slice
//    fmt.Println(slice)
//    fmt.Println(*ppslice)
//    fmt.Println(ppslice)// 二级指针
//    fmt.Printf("%T\n", ppslice)
//    fmt.Println(*ppslice)
//    // 切片名本身就是一个地址
//    (*ppslice)[0] = 9
//    fmt.Println(slice)
//    // ppslice[0] = 9 error 和数组指针不同
//    test2(ppslice)
//    fmt.Println(slice)
//
//    // 用new创建切片指针空间
//    var p3 *[]int
//    fmt.Printf("%p\n", p3)
//    p3 = new([]int)
//    fmt.Printf("%p\n", p3)
//    *p3 = append(*p3, 1, 2, 3, 4, 5)
//    for i := 0; i < len(*p3); i++ {
//        fmt.Println((*p3)[i])
//    }
//
//    // 指针数组 指针切片
//    a, b, c := 1, 2, 3
//    var pointarr [3]*int = [3]*int{&a, &b, &c}
//    *pointarr[0] = 99
//    fmt.Println(a)
//
//    // 指针切片
//    var pointslice []*int
//    pointslice = append(pointslice, &a, &b, &c)
//    *pointslice[2] = 99
//    fmt.Println(c)
//
//    // 结构体指针
//    var st student = student{101, "bob"}
//    fmt.Println(st)
//    var pstudent *student = &st
//    pstudent.id = 102
//    pstudent.name = "marrys"
//    fmt.Println(st)
//
//    // 多级指针
//    x := 10
//    px := &x
//    ppx := &px
//
//}
