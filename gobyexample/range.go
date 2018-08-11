package main

import (
	"fmt"
)

func main() {
	fmt.Println("类似于 php 中的 foreach")

	nums := []int{2, 3, 4}
	sum := 0
	//使用 空值定义符_ 来忽略索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("slice 的元素个数 :", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index is ", i)
		}
	}

	kvs := map[string]string{"name": "time", "sex": "boy"}
	for k, v := range kvs {
		fmt.Println(k + "->" + v)
	}

	fmt.Println("range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
