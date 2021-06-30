package main

import (
    "fmt"
    "time"
    "net"
)

func main() {
    // 创建一个UDP地质结构 指定服务器的IP:port
    srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
    if err != nil {
        fmt.Println("net.ResolveUDPAddr error:", err)
        return
    }
    fmt.Println("Create udpAddr ...")
    // 创建用于通信的socket
    udpConn, err := net.ListenUDP("udp", srvAddr)
    if err != nil {
        fmt.Println("net.ListenUDP error:", err)
        return
    }
    defer udpConn.Close()
    fmt.Println("wait for read ...")
    // Read
    buf := make([]byte, 4096)
    n, cltAddr, err := udpConn.ReadFromUDP(buf)
    if err != nil {
        fmt.Println("ReadFromUDP error:", err)
        return
    }
    // 处理
    fmt.Println("Server Read:", string(buf[:n]), "total length:", n)
    // echo
    daytime := time.Now().String()

    _, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)

    if err != nil {
        fmt.Println("WriteUDP error:", err)
        return
    }

}

// 客户端：nc -u 127.0.0.1 8003
