package main

import (
    "fmt"
)

func main() {
    a := make([]int, 10)
    for i := 0; i < 10; i++ {
        a[i] = i
    }

    fmt.Println(a)

    cnt := copy(a[:5], a[5:])
    fmt.Println(a)
    fmt.Println("copy cnt = ", cnt)
}
