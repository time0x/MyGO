package main

import (
	"fmt"
)

func main() {
	//默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
	message := make(chan string, 2)

	//因为这个通道是有缓冲区的，即使没有一个对应的并发接收方，我们仍然可以发送这些值。
	//go func() { message <- "ping" }()

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
