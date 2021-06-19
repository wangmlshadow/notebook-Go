package main

import (
    "fmt"
    "errors"
)

func test(a int, b int) (value int, err error) {
    //return a / b
    if b == 0 {
        err = errors.New("0不能作为除数")
        return
    } else {
        value = a / b
        return 
    }
}

func test2(a int, b int) (value int) {
    value = a / b
    return value
}

func test3() {
    fmt.Println("test3")
    // 调用panic程序自动终止
    panic("test3 panic")
}

func main() {
    value, err := test(10, 0)// 注意此处问题 0不能作为除数
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(value)
    }
    value1, err1 := test(10, 2)// 注意此处问题 0不能作为除数
    if err1 != nil {
        fmt.Println(err1)
    } else {
        fmt.Println(value1)
    }

    // panic异常处理
    //value2 := test2(10, 0)// panic: runtime error:
    //fmt.Println(value2)

    // 
    //test3()

    // 延迟调用
    // defer 在程序运行结束后再运行 先defer的后调用
    fmt.Println("Hello 1")
    defer fmt.Println("Hello 2")
    fmt.Println("Hello 3")
    defer fmt.Println("Hello 4")
    fmt.Println("Hello 5")
    
    a, b := 10, 20
    f := func(a int, b int) {
        fmt.Println(a)
        fmt.Println(b)
    }
    f2 := func() {
        fmt.Println(a)
        fmt.Println(b)
    }
    // 注意变量作用域
    defer f(a, b)// 输出10, 20
    defer f2()// 输出100， 200

    a, b = 100, 200

    fmt.Println("end...")
}

