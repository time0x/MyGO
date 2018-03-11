// 当前程序的包名
package main

// 导入的其他包名,导入的包必须使用，否则会报错
import (
    // 如果别名使用 . 的话，调用包的方法时候可以直接调用
    otherFmtName "fmt"
)

// 常量的赋值
const PAI = 3.1415926

// 全局变量的声明与赋值
var name = "time"

// 一般类型赋值
type age int

// 结构的声明
type goxx struct{}

// 接口的声明
type goxxx interface{}

// 由main 函数作为程序入口启动
func main() {
    // fmt.Println("hello world!")
    // Println("hello world!")  别名为 . 可以直接调用
    otherFmtName.Println("hello world!")
}