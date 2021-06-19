package main

import "fmt"

//func add(a, b int) int {
//	return a + b
//}

// 起别名
type Int int

// 方法
// func (方法接收者)方法名(参数列表)返回值类型
func (a Int) add (b Int) Int {
	return a + b
}

type student struct {
	name string
	age int
	sex string
}

// 为结构体定义别名
func (stu student) PrintInfo () {
	fmt.Println(stu.sex)
	fmt.Println(stu.name)
	fmt.Println(stu.age)
}

func (stu *student) Rename () {
	stu.name = "asdasdasd"
}

// 方法的继承
type test struct {
	student
	index int
}

// 方法的重写
func (t test) PrintInfo () {
	t.student.PrintInfo()
	fmt.Println(t.index)
}

// 方法类型和方法的值

func main() {
	// 根据数据类型绑定方法
	var a Int = 1
 	var b Int = 3
	fmt.Println(a.add(b))

 	stu := student{"sdsad", 10, "male"}
 	stu.PrintInfo()
 	stu.Rename()
 	fmt.Println(stu)

 	var t test
 	t.index = 101
 	t.age = 19
 	t.sex = "male"
 	t.name = "Bbb"
 	t.student.PrintInfo()
 	t.PrintInfo()
 	fmt.Println(t)
 	t.Rename()
 	fmt.Println(t)

 	f1 := t.PrintInfo
 	f2 := t.student.PrintInfo
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)

 	f1()
 	f2()

}
