package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    str := "aaabbbcccddd"
    // 判断是否存在子串
    value := strings.Contains(str, "aaa")

    if value {
        fmt.Println("str存在aaa")
    } else {
        fmt.Println("str不存在aaa")
    }
    
    fmt.Println(value)

    s := []string{"aaa", "bbb", "ccc"}
    // 字符串拼接
    str = strings.Join(s, "-")
    fmt.Println(str)

    str = "123456789"
    // 查找字串位置 返回下标值 未找到返回-1
    idx1 := strings.Index(str, "456")
    idx2 := strings.Index(str, "436")
    fmt.Println(idx1)
    fmt.Println(idx2)

    // 重复字符串
    str = "ABCD"
    res := strings.Repeat(str, 3)
    fmt.Println(res)

    // 字符串转换
    // 将其他类型转换成字符串
    str = strconv.FormatBool(false)
    fmt.Println(str)

    str = strconv.FormatInt(123, 10)// 十进制数123
    fmt.Println(str)

    str = strconv.Itoa(12445)
    fmt.Println(str)

    str = strconv.FormatFloat(1.254, 'f', 5, 64)// 浮点数 1.254 小数位5 64位
    fmt.Println(str)

    // 字符串转换为其他类型
    str = "false"
    // 忽略错误信息
    //b, _ := strconv.ParseBool(str)
    b, err := strconv.ParseBool(str)
    //fmt.Println(b)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(b)
    }

    // 将字符串转化为整形
    str = "1010101"
    //a, _ := strconv.ParseInt(str, 2, 64)// 将2进制64位整数str转为10进制数
    a, _ := strconv.ParseInt(str, 10, 64)// 将10进制64位整数str转为10进制数
    fmt.Println(a)

    str = "12.35545"
    d, _ := strconv.ParseFloat(str, 64)
    fmt.Println(d)

    // Append
    e := make([]byte, 0, 1024)
    // 将bool类型放在指定切片中
    e = strconv.AppendBool(e, false)
    e = strconv.AppendInt(e, 123, 10)
    e = strconv.AppendFloat(e, 1.234, 'f', 5, 64)
    fmt.Println(string(e))

}
