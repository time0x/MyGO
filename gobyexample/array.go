package main

import (
	"fmt"
	"reflect"
)

//在典型的 Go 程序中，相对于数组而言，slice 使用的更多。我们将在后面讨论 slice。
func main() {
	var a [5]int
	fmt.Println("arr a:", a)
	fmt.Println("arr len:", len(a))
	fmt.Println("arr type is :", reflect.TypeOf(a)) //[5]int

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get a[4]:", a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr b:", b)

	//c:= [2][3]int //is not an expression
	var c [2][3]int
	fmt.Println(len(c))
	for i := 0; i < len(c); i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2 arr :", c)
}
