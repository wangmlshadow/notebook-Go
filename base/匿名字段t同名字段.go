package src

//import "fmt"
//
//type person struct {
//	name string
//	age int
//	sex string
//}
//
////  结构体嵌套
//type student struct {
//	person// 匿名字段
//	id int
//	score int
//	name string// 同名字段
//}
//
//type student2 struct {
//	*person// 指针匿名字段
//	id int
//	score int
//	name string// 同名字段
//}
//
//func main() {
//	// var stu student = student{person{...}, ...}
//	var stu student
//	stu.id = 101
//	stu.score = 100
//	stu.name = "nana"
//	stu.person.name = "lala"
//	stu.person.age = 18
//	stu.sex = "female"
//	fmt.Println(stu)
//
//	var stu2 student2
//	//stu2.person.name = "kaka"// invalid memory address or nil pointer dereference
//	stu2.person = new(person)
//	stu2.name = "baba"
//	stu2.person.name = "kaka"
//	fmt.Println(stu2)
//}
