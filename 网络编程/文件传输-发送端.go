package main

import (
    "fmt"
    "os"
    "net"
    "io"
)

func sendFile(conn net.Conn, filePath string) {
    // Read only
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Println("os.Open error:", err)
        return
    }
    defer f.Close()

    buf := make([]byte, 4096)
    for {
        // 从本地文件读 再写道网络对端
        n, err := f.Read(buf)
        if err != nil {
            if err == io.EOF {
                fmt.Println("Send Completed...")
            } else {
                fmt.Println("f.Read error:", err)
            }
            return
        }
        // write to net
        _, err = conn.Write(buf[:n])
        if err != nil {
            fmt.Println("conn.Write to net error:", err)
            return
        }
    }
}

func main() {
    // 命令行参数
    list := os.Args
    //fmt.Println(list)
    if len(list) != 2 {
        fmt.Println("格式：go run xxx.go fileName")
        return
    }
    // 获取文件属性
    path := list[1]
    fileInfo, err := os.Stat(path)
    if err != nil {
        fmt.Println("os.Stat error:", err)
        return
    }
    fmt.Println("文件名:", fileInfo.Name(), "文件大小:", fileInfo.Size())

    // connect
    conn, err := net.Dial("tcp", "127.0.0.1:8008")
    if err != nil {
        fmt.Println("net.Dial error:", err)
        return
    }
    defer conn.Close()

    // send fileName
    fileName := fileInfo.Name()
    conn.Write([]byte(fileName))

    // 读取ok
    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("conn.Read ok error:", err)
        return
    }
    if "ok" == string(buf[:n]) {
        fmt.Println("Read OK")
        // 写文件内容给服务器
        sendFile(conn, path)
    }
}
