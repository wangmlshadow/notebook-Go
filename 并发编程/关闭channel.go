package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 8; i++ {
            ch <- i
        }
        close(ch)// 写端写完数据 主动关闭channel
    }()

    for {
        // 检测对端是否关闭
        if num, ok := <- ch; ok == true {// ok == true, 读到数据
            fmt.Println("Read num =", num)
        } else {// channel已经关闭
            break
        }
    }

    // 或者换种写法
    // for num := range ch {}
}

/*
数据未发送完不应该关闭channel

无缓冲channel 无法向已经关闭的channel中写数据 但是还可以读
*/
