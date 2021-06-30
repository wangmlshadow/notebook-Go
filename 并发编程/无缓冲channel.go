package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("子go程写， i =", i)
            ch <- i
        }
    }()

    //time.Sleep(time.Second * 2)

    for i := 0; i < 5; i++ {
        num := <- ch
        fmt.Println("主go程读，i = ", num)
    }
}

// 第一次输出
// Print打印会阻塞
/*
子go程写， i = 0
子go程写， i = 1
主go程读，i =  0
主go程读，i =  1
子go程写， i = 2
子go程写， i = 3
主go程读，i =  2
主go程读，i =  3
子go程写， i = 4
主go程读，i =  4
*/
