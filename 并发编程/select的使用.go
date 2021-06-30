package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)// 用于数据通信的channel
    quit := make(chan bool)// 用于判断是否退出的channel

    // 子go写数据
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            time.Sleep(time.Second)
        }
        close(ch)// ch 虽然关闭 但是还可以读到0
        quit <- true
    }()

    // 主go读数据
    for {
        // select下的case中 若果某个case可读 则执行 
        // 如果所有case都不可读 则阻塞在select
        // case中有多个满足监听条件 任选一个执行
        // 可以使用default来处理所有case都不满足监听条件的状况 通常不会这么使用  会产生忙等待
        // select自身不带有循环机制 需要借助外层for循环来监听
        // break只能跳出select
        select {
        case num := <- ch:
            fmt.Println("Read:", num)
        case <- quit:// quit 可读 退出for
            fmt.Println("quit")
            // break跳出的是select
            //break
            return
        }
        // select执行后执行
        fmt.Println("-----------------")
    }
}
