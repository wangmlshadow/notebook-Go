package main

import (
    "fmt"
    "math/rand"
    "time"
//    "sync"
)

// 读写锁
//var rwMutex sync.RWMutex

// 在go中尽量不要将互斥锁 读写锁 与 channel混用 可能造成隐形死锁
// 下面程序会死锁
// 不使用channel 而是用全局变量
/*
func readGo(in <- chan int, idx int) {
    for {
        rwMutex.RLock()// 读 加锁
        num := <- in
        fmt.Println("Id", idx, "Read", num)
        rwMutex.RUnlock()// 读 解锁
    }
}

func writeGo(out chan <- int, idx int) {
    for {
        // 生成随机数
        num := rand.Intn(1000)
        rwMutex.Lock()// 写 加锁
        out <- num
        fmt.Println("Id", idx, "Write", num)
        //time.Sleep(time.Millisecond * 300)
        rwMutex.Unlock()
        time.Sleep(time.Millisecond * 300)
    }
}

func main() {
    // 随机数种子
    rand.Seed(time.Now().UnixNano())

    ch := make(chan int)
    //quit := make(chan bool)

    // 5个读go程 5个写go程
    for i := 0; i < 5; i++ {
        go readGo(ch, i)
    }

    for i := 0; i < 5; i++ {
        go writeGo(ch, i)
    }

    //<- quit
    for {
        ;
    }
}
*/

/*
// 使用全局变量
var value int// 定义全局变量 模拟共享数据

func readGo(idx int) {
    for {
        rwMutex.RLock()// 读 加锁
        num := value
        fmt.Println("Id", idx, "Read", num)
        time.Sleep(time.Millisecond * 300)
        rwMutex.RUnlock()// 读 解锁
    }
}

func writeGo(idx int) {
    for {
        // 生成随机数
        num := rand.Intn(1000)
        rwMutex.Lock()// 写 加锁
        value = num
        fmt.Println("Id", idx, "Write", num)
        time.Sleep(time.Millisecond * 300)
        rwMutex.Unlock()
    }
}

func main() {
    // 随机数种子
    rand.Seed(time.Now().UnixNano())

    //ch := make(chan int)
    //quit := make(chan bool)

    // 5个读go程 5个写go程
    for i := 0; i < 5; i++ {
        go readGo(i)
    }

    for i := 0; i < 5; i++ {
        go writeGo(i)
    }

    //<- quit
    for {
        ;
    }
}
*/

// 使用channel模拟读写锁
var value int

func readGo(in <- chan int, idx int) {
    for {
        num := <- in
        fmt.Println("Id", idx, "Read", num)
        time.Sleep(time.Millisecond * 300)
    }
}

func writeGo(out chan <- int, idx int) {
    for {
        // 生成随机数
        num := rand.Intn(1000)
        out <- num
        fmt.Println("Id", idx, "Write", num)
        time.Sleep(time.Millisecond * 300)
    }
}

func main() {
    // 随机数种子
    rand.Seed(time.Now().UnixNano())

    ch := make(chan int)
    //quit := make(chan bool)

    // 5个读go程 5个写go程
    for i := 0; i < 5; i++ {
        go readGo(ch, i)
    }

    for i := 0; i < 5; i++ {
        go writeGo(ch, i)
    }

    //<- quit
    for {
        ;
    }
}
