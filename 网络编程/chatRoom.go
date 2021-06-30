package main

import (
    "fmt"
    "net"
    "strings"
    "time"
)

// user
type Client struct {
    C chan string
    Name string
    Addr string
}

// 全局map 存储在线用户
var onlineMap map[string]Client

// 全局channel 传递用户消息
var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
    // 
    for msg := range clnt.C {
        conn.Write([]byte(msg + "\n"))
        //conn.Write([]byte(msg))
    }
}

func MakeMsg(clnt Client, msg string) (buf string) { 
    buf = "[" + clnt.Addr + "]" + clnt.Name + ": "  + msg
    return
}

func HandlerConnect(conn net.Conn) {
    defer conn.Close()
    // create user 
    // default name = ip port
    netAddr := conn.RemoteAddr().String()
    clnt := Client{make(chan string), netAddr, netAddr}

    // 将新连接的用户添加到map
    onlineMap[netAddr] = clnt

    // 创建专门给当前用户发消息的go程
    go WriteMsgToClient(clnt, conn)

    // 发送 用户上线消息到全局channel
    //message <- "[" + netAddr + "]" + clnt.Name + "login"
    message <- MakeMsg(clnt, "login")

    // 创建一个channel 用来判断用户退出状态
    isQuit := make(chan bool)
    
    // 超时 强行踢出用户
    hasMsg := make(chan bool)


    // 创建一个匿名go程 专门处理用户发送的消息
    go func() {
        buf := make([]byte, 4096)
        for {
            n, err := conn.Read(buf)
            if n == 0 {// 用户退出 断开连接
                fmt.Println("客户端：", clnt.Name, "exit")
                isQuit <- true
                return
            }
            if err != nil {
                fmt.Println("conn.Read error:", err)
                return
            }
            // 将读到的消息广播
            msg := string(buf[:n-1])// 去掉换行
            hasMsg <- true

            // 提取在线用户列表 who命令
            if msg == "who" && len(msg) == 3 {
                conn.Write([]byte("online user list:\n"))
                // 遍历map 获取在线用户
                for _, user := range onlineMap {
                    userInfo := user.Addr + ":" + user.Name + "\n"
                    conn.Write([]byte(userInfo))
                }
            } else if len(msg) >= 8 && msg[:6] == "rename" {// 改名 rename|newName
                newName := strings.Split(msg, "|")[1]
                clnt.Name = newName
                onlineMap[netAddr] = clnt// 更新map
                conn.Write([]byte("rename successful\n"))

            } else {
                message <- MakeMsg(clnt, msg)
            }

        }
    }()

    // 保证不退出
    for {
        // 监听channel上的数据流动
        select {
        case <- isQuit:
            close(clnt.C)// 用于退出 WriteMsgToClient()
            delete(onlineMap, clnt.Addr)// 将用户从map移除
            message <- MakeMsg(clnt, "logout")// 写入用户退出消息到全局channel
            return
        // 超时
        case <- hasMsg:
            // 用于刷新计时器 什么都不用做
        case <- time.After(time.Second * 30):
            close(clnt.C)// 用于退出 WriteMsgToClient()
            delete(onlineMap, clnt.Addr)// 将用户从map移除
            message <- MakeMsg(clnt, "timeout logout")// 写入用户退出消息到全局channel
            return
        }
    }
}

func Manager() {
    // 初始化map
    onlineMap = make(map[string]Client)

    for {
        // 监听全局channel中是否有数据
        msg := <- message

        // 循环发送消息给所有在线用户
        for _, clnt := range onlineMap {
            clnt.C <- msg

        }
    }
}

func main() {
    // listen
    listener, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("Listen error:", err)
        return
    }
    defer listener.Close()

    // 创建管理者go程 管理map和全局channel
    go Manager()


    // 循环监听客户端连接请求
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Accept error:", err)
            return
        }
        // 启动Go程处理客户端请求
        go HandlerConnect(conn)
    }
}
