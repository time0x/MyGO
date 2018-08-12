package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("发送一个值来通知我们已经完工啦。")

	done <- true
}
func main() {
	fmt.Println("使用通道来同步 Go 协程间的执行状态")

	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("chanel is waiting...")
	<-done
}
