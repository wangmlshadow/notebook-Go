# Go并发编程

[TOC]

本文是作者学习Go并发编程的笔记总结，主要内容是Go并发编程的示例代码，下面是与本文相关的链接：

[Go基础部分](https://www.cnblogs.com/lnlin/p/14903466.html)

本文中包含的代码段的完整代码可以去作者的[Github](https://github.com/wangmlshadow/notebook-Go/blob/main/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/%E5%B9%B6%E5%8F%91%E4%BD%BF%E7%94%A8.go)下载

[进程、线程、协程、Go程的解释](https://www.cnblogs.com/yinzhengjie2020/p/12556606.html)

## 说明：

在Go并发编程的学习过程中，除了具体如何使用Go实现并发编程外，还包括进程、线程、协程、生产者消费者模型、互斥量、锁、条件变量等，下文并不会详细说明这些概念，如果有想要详细了解这些内容，可以去看Unix系统编程和Unix网络编程这两本书。

## Go程

Go在语言级别支持协程，叫goroutine。Go语言标准库提供的所有系统调用操作(包括所有同步IO操作)，都会出让CPU给其他goroutine。这让轻量级线程的切换管理不依赖于系统的线程和进程，也不需要依赖于CPU的核心数量。

### Go程的创建与使用

创建时只需要使用关键字 **go**

```go
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
```

### Go程使用的相关函数说明：

Gosched()、GOMAXPROCS()

```go
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
```

```go
func main() {
    // 创建一个子go程
    go func () {
        for i := 0; i < 5; i++ {
            fmt.Println("-----I am Goroutine-----")
            time.Sleep(time.Second)
        }
    }()

    // 主Goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("-----I am main-----")
        time.Sleep(time.Second)
        if i == 2 {
             break
        }
    }

     
    // 主Goroutine退出 子go程也会退出
    /*
    -----I am main-----
    -----I am Goroutine-----
    -----I am Goroutine-----
    -----I am main-----
    -----I am main-----
    -----I am Goroutine-----
    -----I am Goroutine-----
    */

    // runtime.Gosched() 出让当前go程所占用的cpu时间片
}
```

```go
// runtime.GOMAXPROCS(n) 设置当前进程使用的最大cpu内核数 返回上一次调用成功的设置值 首次调用返回默认值

func main() {
    n := runtime.GOMAXPROCS(2)// 将CPU设为双核
    fmt.Println(n)

    for {
        go fmt.Print(0)// 子go程
        fmt.Print(1)// 主go程
    }
}
```

## Channel

channel是一种数据类型(管道)，主要用于解决go程同步问题以及协程之间数据共享的问题。

特点：一端写一端读

### Channel的定义与使用

```go
/*
make(chan 在channel中传递的数据类型, 容量)
容量为0表示无缓冲
容量大于0表示有缓冲
*/
// 全局定义channel 完成数据同步
var channel = make(chan int)

func printer(s string) {
    for _, ch := range s {
        fmt.Printf("%c", ch)
        time.Sleep(300 * time.Millisecond)
    }
}

// 先执行
func person1() {
    printer("Hello")
    channel <- 1// 向channel写数据 如果写的数据没有被读走 channel阻塞
}

// 后执行
func person2() {
    <- channel// 从channel读 
    printer("World")
}

func main() {
    go person1()
    go person2()
    // 输出WHeorllldo person1 person2 交替使用标准输出 导致输出结果乱序

    for {
        ;
    }
}
```

### Channel同步传递数据

```go
func main() {
    ch := make(chan string)
    // len() 得到channel中剩余未读取数据个数
    // cap() 得到通道容量
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))

    go func () {
        for i := 0; i < 5; i++ {
            fmt.Println("i = ", i)
        }
        // 通知主go打印完毕
        ch <- "Completed..."
        fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    }()
    
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    str := <- ch
    //fmt.Println("len(ch) =", len(ch), "cap(ch) =", cap(ch))
    fmt.Println("主go", str)
}
```

### 无缓冲Channel和有缓冲Channel

```go
func main() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("子go程写， i =", i)
            ch <- i
        }
    }()

    //time.Sleep(time.Second * 2)

    for i := 0; i < 5; i++ {
        num := <- ch
        fmt.Println("主go程读，i = ", num)
    }
}
```

```go
func main() {
    ch := make(chan int, 3)// 存满3个元素之前不会阻塞
    fmt.Println("len =", len(ch), "cap =", cap(ch))

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            fmt.Println("子go程：", i)
            fmt.Println("len =", len(ch), "cap =", cap(ch))
        }
    }()
    time.Sleep(time.Second * 3)
    for i := 0; i < 5; i++ {
        num := <- ch
        fmt.Println("主go程：", num)
        fmt.Println("len =", len(ch), "cap =", cap(ch))
    }
}
```

### Channel的关闭

```go
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
```

### 单向Channel的使用

有时在函数中，我们只需要从channel中读取数据或者写入数据，这时我们可以使用单向channel

```go
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
```

### 使用Channel实现生产者消费者模型

生产者：发送端
消费者：接收端

缓冲区作用：
        解耦（降低生产者与消费者之间的耦合度）
        并发（生产者与消费者数量不对等时 能保持正常通信）
        缓存（生产者与消费者数据处理速度不一致时 暂存数据）

```go
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
```

## 定时器

### 创建定时器

```go
func main() {
    fmt.Println("Now time:", time.Now())
    // 创建定时器
    myTimer := time.NewTimer(time.Second * 2)
    nowTime := <- myTimer.C// 当系统向定时器中的channel写完后 再从中读
    fmt.Println("Now time:", nowTime)
}
/*
time.Timer
定时器，由channel实现，当设定的时间到达时，系统会向定时器中的channel写
type Timer struct {
    C <- chan Time
    r runtimeTimer
}
*/
```

### 定时的3种方式

```go
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
```

### 周期定时

```go
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
```

## select

使用select监听每个channel

### select的使用

```go
func main() {
    ch := make(chan int)// 用于数据通信的channel
    quit := make(chan bool)// 用于判断是否退出的channel

    // 子go写数据
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            time.Sleep(time.Second)
        }
        close(ch)// ch 虽然关闭 但是还可以读到0
        quit <- true
    }()

    // 主go读数据
    for {
        // select下的case中 若果某个case可读 则执行 
        // 如果所有case都不可读 则阻塞在select
        // case中有多个满足监听条件 任选一个执行
        // 可以使用default来处理所有case都不满足监听条件的状况 通常不会这么使用  会产生忙等待
        // select自身不带有循环机制 需要借助外层for循环来监听
        // break只能跳出select
        select {
        case num := <- ch:
            fmt.Println("Read:", num)
        case <- quit:// quit 可读 退出for
            fmt.Println("quit")
            // break跳出的是select
            //break
            return
        }
        // select执行后执行
        fmt.Println("-----------------")
    }
}
```

### select超时处理

```go
func main() {
    ch := make(chan int)
    quit := make(chan bool)

    go func() {
        for {
            select {
            case num := <- ch:
                fmt.Println("Read:", num)
            case <- time.After(3 * time.Second):// 超过3秒还没读到数据
                quit <- true
            }
        }
    }()

    for i := 0; i < 5; i++ {
        ch <- i
        time.Sleep(time.Second)
    }

    <- quit
    fmt.Println("quit")
}
```

## 同步相关

### 使用channel产生死锁

```go
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
```

### 互斥锁

```go
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
```

### 读写锁

#### 读写锁的使用

```go
// 读写锁
//var rwMutex sync.RWMutex

// 在go中尽量不要将互斥锁 读写锁 与 channel混用 可能造成隐形死锁
// 下面程序会死锁
// 不使用channel 而是用全局变量

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
```

```go
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
```

#### 使用channel模拟读写锁

```go
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
```

### 条件变量

```go
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
```



