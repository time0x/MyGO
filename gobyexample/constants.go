package main

import (
	"fmt"
	"math"
	"reflect"
)

const a string = "show your code"

func main() {
	fmt.Println(a)

	const b = 50000000
	fmt.Println(b)

	const c = 3e20
	fmt.Println(c)
	fmt.Println(c / b)

	fmt.Println(int64(c / b))

	fmt.Println("当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数")
	const d = 20
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(math.Sin(d))
}
