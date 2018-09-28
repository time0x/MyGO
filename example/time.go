package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前纳秒
	fmt.Println(time.Now().UnixNano())
	//秒数
	fmt.Println(time.Now().Unix())
}
