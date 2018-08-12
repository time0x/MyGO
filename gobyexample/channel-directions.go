package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	fmt.Println("当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。这个特性提升了程序的类型安全性")

	fmt.Println("ping 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	fmt.Println("pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据")
	pong(pings, pongs)
	fmt.Println("pongs 接收到:", <-pongs)

}
