package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("递归递减相乘中，n=：", n)
	return n * fact(n-1)
}
func main() {
	fmt.Println("递归 demo：")
	fmt.Println(fact(6))
}
