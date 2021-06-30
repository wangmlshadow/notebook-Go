package main

import (
    "fmt"
    "net"
    "strings"
)

func HandlerConnect(conn net.Conn) {
    defer conn.Close()
    // 获取连接的客户端
    addr := conn.RemoteAddr()
    fmt.Println(addr, "contected...")

    // 循环读
    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        // 对端关闭
        if n == 0 {
            fmt.Println("client closed...")
            return
        }
        if err != nil {
            fmt.Println(addr, "Read error:", err)
            return
        }
        //fmt.Println("Read from:", addr, "Data:", string(buf[:n]), "len:", n)
        fmt.Println("Read from:", addr, "Data:", string(buf[:n]))
        // 数据处理
        conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
    }
}

func main() {
    // listen
    listener, err := net.Listen("tcp", "127.0.0.1:8001")
    if err != nil {
        fmt.Println("net.Listen error:", err)
        return
    }
    defer listener.Close()
    fmt.Println("wait for new connection...")

    // 监听
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Accept error:", err)
            return
        }
        //defer conn.Close()

        // 处理
        HandlerConnect(conn)
    }

}
