package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    quit := make(chan bool)

    go func() {
        for {
            select {
            case num := <- ch:
                fmt.Println("Read:", num)
            case <- time.After(3 * time.Second):// 超过3秒还没读到数据
                quit <- true
            }
        }
    }()

    for i := 0; i < 5; i++ {
        ch <- i
        time.Sleep(time.Second)
    }

    <- quit
    fmt.Println("quit")
}
