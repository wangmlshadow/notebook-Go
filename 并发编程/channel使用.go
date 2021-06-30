package main

import (
    "fmt"
    "time"
)

// 全局定义channel 完成数据同步
var channel = make(chan int)


func printer(s string) {
    for _, ch := range s {
        fmt.Printf("%c", ch)
        time.Sleep(300 * time.Millisecond)
    }
}

// 先执行
func person1() {
    printer("Hello")
    channel <- 1// 向channel写数据 如果写的数据没有被读走 channel阻塞
}

// 后执行
func person2() {
    <- channel// 从channel读 
    printer("World")
}

func main() {
    go person1()
    go person2()
    // 输出WHeorllldo person1 person2 交替使用标准输出 导致输出结果乱序

    for {
        ;
    }
}
