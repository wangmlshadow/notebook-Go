package main

import (
    "fmt"
)

func main() {
    ch := make(chan string)
    // len() 得到channel中剩余未读取数据个数
    // cap() 得到通道容量
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))

    go func () {
        for i := 0; i < 5; i++ {
            fmt.Println("i = ", i)
        }
        // 通知主go打印完毕
        ch <- "Completed..."
        fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    }()
    
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    str := <- ch
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    fmt.Println("主go", str)
}
