package main

/*
条件变量
加锁 操作公共区 解锁 唤醒对端
*/

import (
    "fmt"
    "time"
    "sync"
    "math/rand"
)

var cond sync.Cond// 全局条件变量

func producer(out chan <- int, idx int) {
    for {
        // 加锁
        cond.L.Lock()
        // 判断缓冲区是否满
        for len(out) == 5 {
            cond.Wait()// 等待缓冲区有位置可写
        }
        num := rand.Intn(800)
        out <- num
        fmt.Println("Idx", idx, "Write", num)
        // 解锁
        cond.L.Unlock()
        // 唤醒对端 即消费者
        cond.Signal()
        time.Sleep(time.Millisecond * 200)
    }
}

func consumer(in <- chan int, idx int) {
    for { 
        cond.L.Lock()
        for len(in) == 0 {
            cond.Wait()
        }
        num := <- in
        fmt.Println("idx", idx, "Read", num)
        cond.L.Unlock()
        cond.Signal()
        time.Sleep(time.Millisecond * 200)
    }
}

func main() {
    ch := make(chan int, 5)
    //quit := make(chan int)
    rand.Seed(time.Now().UnixNano())

    // 指定条件变量使用的锁
    cond.L = new(sync.Mutex)

    for i := 0; i < 5; i++ {
        go producer(ch, i)
    }

    for i := 0; i < 5; i++ {
        go consumer(ch, i)
    }
    
    //<- quit
    for {
        ;
    }
}
