package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int, 3)// 存满3个元素之前不会阻塞
    fmt.Println("len =", len(ch), "cap =", cap(ch))

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            fmt.Println("子go程：", i)
            fmt.Println("len =", len(ch), "cap =", cap(ch))
        }
    }()
    time.Sleep(time.Second * 3)
    for i := 0; i < 5; i++ {
        num := <- ch
        fmt.Println("主go程：", num)
        fmt.Println("len =", len(ch), "cap =", cap(ch))
    }
}
