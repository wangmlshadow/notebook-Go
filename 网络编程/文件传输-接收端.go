package main

import (
    "fmt"
    "net"
    "os"
)

func recvFile(conn net.Conn, fileName string) {
    // 创建文件
    f, err := os.Create(fileName)
    if err != nil {
        fmt.Println("os.Create error:", err)
        return
    }
    defer f.Close()

    // read
    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        if n == 0 {
            fmt.Println("Read Completed...")
            return
        }
        if err != nil {
            fmt.Println("conn.Read from net error:", err)
            return
        }
        // 将读取的数据 写入本地文件
        f.Write(buf[:n])
    }
}

func main() {
    // listen
    listener, err := net.Listen("tcp", "127.0.0.1:8008")
    if err != nil {
        fmt.Println("net.Listen error:", err)
        return
    }
    defer listener.Close()

    // 阻塞监听 
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Accept error:", err)
        return
    }
    defer conn.Close()

    // 获取文件名 保存
    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("conn.Read error:", err)
        return
    }
    fileName := string(buf[:n])

    // 回写 ok
    conn.Write([]byte("ok"))

    // 获取文件内容
    recvFile(conn, fileName)
}
