package main

import (
    "fmt"
    "time"
)

func sing() {
    for i := 0; i < 5; i++ {
        fmt.Println("Sing something...")
        time.Sleep(100 * time.Millisecond)
    }
}


func dance() {
    for i := 0; i < 5; i++ {
        fmt.Println("Someone dancing...")
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    fmt.Println("顺序执行")
    sing()
    dance()
    // 并发执行
    fmt.Println("并发执行")
    go sing()
    go dance()

    for {
        ;
    }
}
