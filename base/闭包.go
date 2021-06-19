package src

import "fmt"

func test(a int)  {
	fmt.Println(a)
}

// 通过匿名函数和闭包 实现函数在栈区的本地化
func test1() func() int {
	var a int = 1

	return func() int {
		a++
		return a
	}
}

func main() {
	//for i := 0; i < 10; i++ {
	//	test(i)
	//}

	f := test1()// func() int

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
