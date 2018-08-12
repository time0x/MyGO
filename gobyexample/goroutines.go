package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 6; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	fmt.Println("Go 协程 在执行上来说是轻量级的线程")
	fmt.Println("我们使用一般的方式调并同时运行。")
	f("direct")

	fmt.Println("使用协程并行执行")
	go f("goroutine")

	go func(msg string) {
		fmt.Println("msg")
	}("going")

	fmt.Println("输入任意键结束 main ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
