package main

import "fmt"

type person struct {
    name string
    sex string
    age int
}

type student struct {
    person
    score int
}

type teacher struct {
    person
    subject string
}


func (s *student)SayHello(){
    fmt.Printf("name: %s, score: %d\n", s.name, s.score)
}

func (t *teacher)SayHello(){
    fmt.Printf("name: %s, subject: %s\n", t.name, t.subject)
}

// 接口的定义
// 接口定义了规则 方法实现了规则
type TestInterface interface {
    // 方法的声明 没有具体实现 
    // 接口中定义的方法必须全部有具体的实现
    SayHello()
}

// 多态
type Person interface {
    SayHello()
}

// 多态实现
// 多态是将接口类型作为函数参数
func SayHello(p Person){
    p.SayHello()
}

// 接口的继承
type Speaker interface {
    Person
    Sing(string)
}

func (s *student)Sing(name string){
    fmt.Printf("Sing %s Name %s\n", name, s.name)
}

func main () {
    var stu student = student{ person{"Bob", "male", 18}, 99 }
    var tea teacher = teacher{ person{"marry", "male", 28}, "Math" }

    stu.SayHello()
    tea.SayHello()

    // 定义接口类型
    var h TestInterface
    h = &stu
    h.SayHello()

    h = &tea
    h.SayHello()

    var p Person

    // 接口
    p = &student{ person{"Bob", "male", 18}, 99 }
    // 多态
    SayHello(p)

    p = &teacher{ person{"marry", "male", 28}, "Math" }
    SayHello(p)

    var s Speaker
    s = &stu
    s.SayHello()
    s.Sing("lalala")

    // 接口的转换
    // 将超集转换为子集
    p = s
    p.SayHello()

    // 不允许子集转换为超集
    // s = p // error

    // 空接口
    var i interface{}
    i = 10
    fmt.Println(i)
    i = "Hello"
    fmt.Println(i)

    // 类型断言
    if data, ok := i.(string); ok {
        fmt.Println("string ", data)
    }
}
