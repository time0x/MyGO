package main

import (
	"fmt"
	"strconv"
)

func main() {
	//for
	for i := 21; i < 23; i++ {
		var ip = "172.16.9."
		var newip = ip + strconv.Itoa(i)
		fmt.Println(newip)
	}

	var var_a, var_b = 0, 0
	//  相当于 while 循环
	for {
		var_a++
		if var_a > 3 {
			break
		}
		fmt.Println(var_a)
	}

	// 相当于 do while
	for var_b < 5 {
		var_b++
		fmt.Println(var_b)
	}
	fmt.Println("this for is over")

	//if
	if 1 < 2 {
		fmt.Println("yes ")
	}
	if var_a := 1; var_a < 2 {
		fmt.Println("'i want you'")
	}

	// switch

	switch var_a { // switch var_a:=0 { 也可以初始化，是局部变量
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println(var_a)
	}

	switch {
	case var_a > 0:
		fmt.Println("a >0")
		fallthrough // 如果不加 fallthrough 则此条件满足后，自动跳出(break)，加上则会继续往下执行
	case var_a > 1:
		fmt.Println("a >1")
	default:
		fmt.Println("a is none")
	}

	// 跳转语句 break continue goto
LABEL1: // 标签
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break // 只可以跳出一层循环
			}
		}
	}

}
