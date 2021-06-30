package main

import (
    "fmt"
    "os"
)

func main() {
    // 输入打开的path
    fmt.Println("输出待查询的目录：")
    var path string
    fmt.Scan(&path)

    // open dir
    // os.ModeDir 指定打开的是一个目录
    f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
    if err != nil {
        fmt.Println("Open Dir error")
        return
    }
    defer f.Close()

    // 读取目录项
    // -1 表示获取所有目录信息
    info, err := f.Readdir(-1)
    // 遍历返回的切片信息
    for _, fileInfo := range info {
        if fileInfo.IsDir() {
            fmt.Println(fileInfo.Name(), "is a Dir")
        } else {
            fmt.Println(fileInfo.Name(), "is a File")
        }
    }
}
