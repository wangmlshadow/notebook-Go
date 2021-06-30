package main

import (
    "fmt"
    "runtime"
)

// runtime.GOMAXPROCS(n) 设置当前进程使用的最大cpu内核数 返回上一次调用成功的设置值 首次调用返回默认值

func main() {
    n := runtime.GOMAXPROCS(2)// 将CPU设为双核
    fmt.Println(n)

    for {
        go fmt.Print(0)// 子go程
        fmt.Print(1)// 主go程
    }
}
