package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // 主动发起请求
    conn, err := net.Dial("tcp", "127.0.0.1:8001")
    if err != nil {
        fmt.Println("net.Dial error:", err)
        return
    }
    defer conn.Close()

    go func() {
        str := make([]byte, 4096)
        for {
            n, err := os.Stdin.Read(str)
            if err != nil {
                fmt.Println("stdin read error:", err)
                continue
            }
            conn.Write(str[:n])
        }
    }()
    
    // echo
    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("conn.Read error:", err)
            continue
        }
        fmt.Println("client Read:", string(buf[:n]))
    }

}
