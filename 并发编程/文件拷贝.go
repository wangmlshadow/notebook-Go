package main

import (
    "io"
    "os"
    "fmt"
)

func main() {
    // open
    f_r, err := os.Open("../data/a.txt")
    if err != nil {
        fmt.Println("Open error")
        return
    }
    defer f_r.Close()
    // create
    f_w, err := os.Create("../data/a_copy.txt")
    if err != nil {
        fmt.Println("Create error")
        return
    }
    defer f_w.Close()
    // 从文件读到缓冲区
    buf := make([]byte, 4096)
    // 循环读取 再写入
    for {
        n, err := f_r.Read(buf)// n 表示读取到的字节数
        if err != nil && err == io.EOF {
            fmt.Println("Read All Content")
            return
        }
        f_w.Write(buf[:n])
    }
}
