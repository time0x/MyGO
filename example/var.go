package main

import (
    "fmt"
    "strconv"
)

type (
    文本 string
)

var (
    aa int
    bb bool
    str string
)
func main() {
    var a 文本
    a = "中文类型名，UTF-8 优势体现"
    fmt.Println(a)

    var b int
    b = 1
    fmt.Println(b)

    var c = 1 // 省略变量类型，系统根据赋值来推断类型
    fmt.Println(c)

    d := false // 极简写法
    fmt.Println(d)

    var e,f,g int = 2,3,4
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)

    // _ 是一个空白符号
    var h,i,_,j = 5,6,7,8// h,i,_,j := 5,6,7,8
    fmt.Println(h)
    fmt.Println(i)
    fmt.Println(j)

    var k float32 = 9.1
    l := int(k) // 如果之前 l 被声明过则不需要：
    fmt.Println(l)

    var m = 65
    n := string(m)
    fmt.Println(n) // 输出A，因为根据 ACSCII 65表示A

    // 整型转换为字符串
    o := strconv.Itoa(m)
    fmt.Println(o)

    p,_ := strconv.Atoi(o)
    fmt.Println(p)
}
