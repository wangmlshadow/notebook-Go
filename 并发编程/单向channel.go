package main

import (
    "fmt"
)
/*
func main() {
    // 双向channel 默认
    ch := make(chan int)

    var sendCh chan <- int// 单向写channel
    // 可以将双向channel转换为单向channel 但是反之不行
    sendCh = ch
    sendCh <- 754

    // 出错 单向写channel不能读
    //num := <- sendCh

    var recvCh <- chan int = ch// 单向读channel
    num := <- recvCh
    fmt.Println(num)

    // 反向赋值 出错
    //var ch2 chan int = sendCh
}
*/
func send(out chan <- int) {
    out <- 88
    close(out)
}

func recv(in <- chan int) {
    n := <- in
    fmt.Println("Recv num =", n)
}

func main() {
    ch := make(chan int)// 双向channel
    
    go func(){
        send(ch)// 双向channel转为写channel
    }()

    recv(ch)
}
