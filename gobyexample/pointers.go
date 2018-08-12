package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("在函数体内部，*iptr将指针/内存地址 解引用(dereferendes)为地址指向的值。所以赋值语句为*i=0这中形式。")
	fmt.Println("简单理解就是传入的是指针，*把指针对应的值取出来")
}

func main() {
	fmt.Println("go 允许在程序中通过引用传递值或者数据结构")

	i := 1
	fmt.Println("i 初始化 =", i)

	zeroval(i)
	fmt.Println("zeroval,传递的拷贝，并没有改变i的值：", i)

	zeroptr(&i)
	fmt.Println("zeroptr，传递的指针，可以改变i的值：", i)
	fmt.Println("指针，pointer", &i)
}
