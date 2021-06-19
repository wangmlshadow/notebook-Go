package main

import (
    "fmt"
)

func demo(i int) {
    var arr [10]int

    defer func() {
        //recover()// 返回值为interface
        //fmt.Println(recover())
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
    }()

    arr[i] = 99// i > 9会出现数组越界
    fmt.Println(arr)
}

func main() {
    demo(10)
    fmt.Println("test recover")
}
