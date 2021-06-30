package main

import (
    "fmt"
)

// 单go程自己死锁
// channel应该在至少两个以上go程中进行通信 否则死锁
func main1() {
    ch := make(chan int)

    // fatal error: all goroutines are asleep - deadlock
    ch <- 748// 程序死锁 卡在这一步 等待ch被读取 而不会执行下面读取ch的那一步

    num := <- ch
    fmt.Println("Read:", num)

}

// go程间channel访问顺序导致死锁
// 使用channel时 读写两端要有同时有机会执行
func main2() {
    ch := make(chan int)
    num := <- ch// 死锁 等待读 导致子go程不会执行 即写操作不会执行
    fmt.Println("Read:", num)

    go func() {
        ch <- 789
    }()
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    // 
    go func() {
        for {
            select {
            case num := <- ch1:
                ch2 <- num
            }
        }
    }()

    for {
        select {
        case num := <- ch2:
            ch1 <- num
        }
    }
}
