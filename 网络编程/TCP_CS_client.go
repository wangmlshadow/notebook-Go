package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("net.Dial error:", err)
        return
    }
    defer conn.Close()

    // 写数据给服务器
    conn.Write([]byte("Breaking Bad"))

    // 接受服务器返回的数据
    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("conn Read error", err)
        return
    }
    fmt.Println("read from server:", string(buf[:n]))
}
