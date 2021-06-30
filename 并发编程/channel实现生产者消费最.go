package main

import (
    "fmt"
)

func producer(out chan <- int) {
    for i := 0; i < 10; i++ {
        fmt.Println("producer send", i * i)
        out <- i * i
    }
    close(out)
}

func consumer(in <- chan int) {
    for num := range in {
        fmt.Println("consumer recv", num)
    }
}

func main() {
    // 无缓冲channel实现生产者消费者
    //ch := make(chan int)
    // 有缓冲
    ch := make(chan int, 5)

    go producer(ch)// 子go程作为生产者
    consumer(ch)
}
// 无缓冲channel实现的输出
/*
producer send 0
producer send 1
consumer recv 0
consumer recv 1
producer send 4
producer send 9
consumer recv 4
consumer recv 9
producer send 16
producer send 25
consumer recv 16
consumer recv 25
producer send 36
producer send 49
consumer recv 36
consumer recv 49
producer send 64
producer send 81
consumer recv 64
consumer recv 81
*/
// 有缓冲
/*
producer send 0
producer send 1
producer send 4
producer send 9
producer send 16
producer send 25
producer send 36
consumer recv 0
consumer recv 1
consumer recv 4
consumer recv 9
consumer recv 16
consumer recv 25
consumer recv 36
producer send 49
producer send 64
producer send 81
consumer recv 49
consumer recv 64
consumer recv 81
*/
