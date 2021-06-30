package main

import (
    "fmt"
    //"time"
    "runtime"
)


// runtime.Gosched() 出让当前go程所占用的cpu时间片
// runtime.Goexit() 结束调用该函数的当前go程

func main() {
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("this is a goroutine test")
            //time.Sleep(100 * time.Microsecond)
        }
    }()

    for {
        runtime.Gosched()
        fmt.Println("this is a main test")
        //time.Sleep(100 * time.Microsecond)
    }

}
