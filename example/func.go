package main

import (
	"fmt"
)

func main() {
	a := 1
	// 参数是引用类型
	fmt.Println("不定长参数，返回的格式是 slice:")
	func_c(a, 2, 3) // [1,2,3]
	// 如果要修改参数，传递参数指针进去
	fmt.Println("如果变量是slice，则传递进去的是内存地址，可以修改，变量不可以，除非传递这个便利那个的指针进去")
	func_d(&a)

	// 函数也是一种类型
	e := func_e
	e() // func_e

	//匿名函数
	func_f := func() {
		fmt.Println("匿名函数 f")
	}
	func_f()

	// 闭包
	closure := func_g(10)
	fmt.Println(closure(2)) //12
	fmt.Println(closure(6)) //16

	//defer
	func_h()

	func_i()
	func_j()
	func_k()
}

func func_a(a int, b string) (int, string) {
	// 返回 2个参数，int 类型和 string 类型
	return a, b
}

// 3个参数都是int
func func_b() (a, b, c int) {
	// a,b,c 已经直接赋值
	a, b, c = 1, 2, 3
	//自动返回 1，2，3，但是为了代码可读性，不建议这么写
	return
}

// 不定长变参 , ...必须是参数列表最后一个参数
func func_c(a ...int) {

	fmt.Println(a)
}

func func_d(a *int) {
	*a = 99
	fmt.Println(*a)
}

func func_e() {
	fmt.Println("func_e")
}

// 闭包
func func_g(x int) func(int) int {
	fmt.Println("外层x地址：", &x)
	return func(y int) int {
		fmt.Println("内层x地址：", &x)
		return x + y
	}
}

//defer
func func_h() {
	fmt.Println("defer:a1")
	defer fmt.Println("a2")
	defer fmt.Println("a3")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	//闭包
	for i := 0; i < 3; i++ {
		defer func() {
			defer fmt.Println(i)
		}()
	}
}

func func_i() {
	fmt.Println("func_i")
}
func func_j() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Rcover in fucn_j")
		}
	}()

	panic("Panic in func_j")
}
func func_k() {
	fmt.Println("func_k")
}
