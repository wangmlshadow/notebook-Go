package main

import (
    "fmt"
    "time"
    "sync"
)

// 借助channel完成数据同步
//var ch := make(chan int)

// 通过锁完成数据同步
var mutex sync.Mutex// 创建互斥锁 新建互斥锁状态为未加锁0 

func printer(str string) {
    mutex.Lock()// 访问共享数据之前加锁
    for _, ch := range str {
        fmt.Printf("%c", ch)
        time.Sleep(time.Millisecond * 300)
    }
    mutex.Unlock()// 共享数据访问结束 解锁
}

func person1() {
    printer("Hello")
    //ch <- 111
}

func person2() {
    //<- ch
    printer("World")
}

func main() {
    go person1()
    go person2()

    for {
        ;
    }

}
