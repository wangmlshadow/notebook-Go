package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个子go程
    go func () {
        for i := 0; i < 5; i++ {
            fmt.Println("-----I am Goroutine-----")
            time.Sleep(time.Second)
        }
    }()

    // 主Goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("-----I am main-----")
        time.Sleep(time.Second)
        if i == 2 {
             break
        }
    }

     
    // 主Goroutine退出 子go程也会退出
    /*
    -----I am main-----
    -----I am Goroutine-----
    -----I am Goroutine-----
    -----I am main-----
    -----I am main-----
    -----I am Goroutine-----
    -----I am Goroutine-----
    */

    // runtime.Gosched() 出让当前go程所占用的cpu时间片
}
