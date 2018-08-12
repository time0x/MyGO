package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2 * (r.height + r.width)
}
func main() {
	fmt.Println("Go 支持在结构体类型中定义方法 。")

	fmt.Println("这里的 area 方法有一个接收器类型 rect。")
	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}
