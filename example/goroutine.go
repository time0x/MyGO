package main

import (
	"fmt"
	"runtime"
)

func Test(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}

func main() {
	fmt.Println("当前cpu核数：", runtime.GOMAXPROCS(runtime.NumCPU()))
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go Test(c, i)
	}
	<-c
}
