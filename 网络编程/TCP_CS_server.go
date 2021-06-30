package main

import (
    "fmt"
    "net"
)

func main() {
    // 指定服务器 通信地址 IP地址 端口号
    listener, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("listen error")
        return
    }
    defer listener.Close()
    fmt.Println("wait for new connection...")
    // 阻塞监听客户端连接请求
    // 成功建立连接 返回用于通信的socket
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Accept error")
        return
    }
    defer conn.Close()
    fmt.Println("established new connection")
    // 读取客户端发送的数据
    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("conn read error")
        return
    }
    // 处理数据
    // echo
    conn.Write(buf[:n])
    // Print
    fmt.Println("Read data:", string(buf[:n]))

}
