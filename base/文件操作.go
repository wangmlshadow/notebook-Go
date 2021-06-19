package main

import (
    "fmt"
    "os"
    "io"
    //"bufio"
)

func main123() {
    // 创建文件 如果指定路径文件存在 则覆盖 否则创建
    fp, err := os.Create("../data/a.txt")
    // 文件创建失败
    if err != nil {
        fmt.Println("file create error")
        return
    }

    fmt.Println("file create successful")

    // 延迟调用 关闭文件
    // 占用内存和缓冲区
    // 文件打开上限
    defer fp.Close()
    
    // 写入文件
    fp.WriteString("hello\n")

    // 返回写入字符数
    n, _ := fp.WriteString("world\n")
    fmt.Println(n)
    // 一个汉字代表三个字符 换行算一个
    n, _ = fp.WriteString("你好\n")
    fmt.Println(n)

    // 写入字符切片
    b := []byte{'a', 'b', 'c', 'd'}
    // 使用Write写入数据
    fp.Write(b)

    // 将字符串转为字符切片写入文件
    str := "HashSort"
    // 字符串和字符切片允许转换
    c := []byte(str)
    fp.Write(c)
}

func main124() {
    // 打开文件
    // OpenFile只能做打开操作 不能创建文件
    fp, err := os.OpenFile("../data/a.txt", os.O_RDWR, 6)

    if err != nil {
        fmt.Println("OpenFile error")
        return
    }
    defer fp.Close()

    // 获取文件字符个数
    n, _ := fp.Seek(0, io.SeekEnd)
    fmt.Println(n)

    b := []byte("HELLO\n")
    // 使用WriteAt进行指定位置插入数据时会依次覆盖
    fp.WriteAt(b, 0)// 0表示offset 表示当前位置
}

func main() {
    // 只读方式打开文件
    fp, err := os.Open("../data/a.txt")
    // 文件打开失败的原因 文件不存在 没有打开权限 打开的文件数目达到上限
    if err != nil {
        fmt.Println("Open error")
        return
    }
    defer fp.Close()

    //b := make([]byte, 1024)

    // 读取文件
    //fp.Read(b)
    //fmt.Println(string(b))

    // 创建切片缓冲区
    //r := bufio.NewReader(fp)
    // 读取一行内容
    //b, _ := r.ReadBytes('\n')
    //fmt.Println(string(b))

    b := make([]byte, 10)

    for {
        n, err := fp.Read(b)
        if err != nil {
            if err == io.EOF {
                break
            }
        }
        fmt.Print(string(b[:n]))
    }
}
