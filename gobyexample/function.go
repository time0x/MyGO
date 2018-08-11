package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 11, 68
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("求和：", total)
}
func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	fmt.Println("go 支持多返回值")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("仅想返回值的一部分的话，你可以使用空白定义符 _")
	_, c := vals()
	fmt.Println(c)

	fmt.Println("可变参数函数。可以用任意数量的参数调用")
	fmt.Println("这个函数使用任意数目的 int 作为参数")
	sum(1, 2)
	sum(1, 2, 3, 4)
	fmt.Println("如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用:func(slice...)")
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)
}
