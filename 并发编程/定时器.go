package main

import (
    "fmt"
    "time"
)
/*
func main() {
    fmt.Println("Now time:", time.Now())
    // 创建定时器
    myTimer := time.NewTimer(time.Second * 2)
    nowTime := <- myTimer.C// 当系统向定时器中的channel写完后 再从中读
    fmt.Println("Now time:", nowTime)

}
*/
/*
time.Timer
定时器，由channel实现，当设定的时间到达时，系统会向定时器中的channel写
type Timer struct {
    C <- chan Time
    r runtimeTimer
}
*/

// 3种定时方法
/*
func main() {
    // 1 sleep
    time.Sleep(time.Second)
    // 2 Timer.C
    // 如上述代码
    // 3 time.After
    nowTime := <- time.After(time.Second * 2)
}
*/
/*
// 定时器的停止和重置
func main() {
    myTimer := time.NewTimer(time.Second * 3)
    myTimer.Reset(1 * time.Second)// 重置定时器
    go func() {
        <- myTimer.C
        fmt.Println("子go程读取定时完毕")
    }()

    //myTimer.Stop()// 设置定时器停止 子go程无法从定时器读到任何数据
    for {
        ;
    }
}
*/

// 周期定时
func main() {
    // 创建一个是否终止的channel
    quit := make(chan bool)
    fmt.Println("now:", time.Now())
    // 周期定时 每隔一秒 系统会向Ticker.C写一次
    myTicker := time.NewTicker(time.Second)
    i := 0
    go func() {
        for {
            nowTime := <- myTicker.C
            i++
            fmt.Println("nowTime:", nowTime)

            // 定时器循环了五次后 向quit写数据
            // 主go程从quit读到数据后 程序退出
            if i == 5 {
                quit <- true
            }
        }
    }()

    <- quit
}

