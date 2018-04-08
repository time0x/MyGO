package main

import (
	"fmt"
)

func main() {
	var arr_a [2]int // 长度为2的int 类数组
	var arr_b [3]int
	var arr_c [3]string

	arr_d := [2]int{22}
	arr_e := [20]int{19: 22}                        // 索引，前19为为默认0值
	arr_f := [...]int{1, 2, 3, 4, 5}                // ... 表示在数组内容固定的情况下，可以不指定长度
	arr_g := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5} // 索引

	// p 是指向数组的指针
	var arr_h *[2]int = &arr_d
	// 指针数组
	x, y := 1, 2
	arr_i := [...]*int{&x, &y}

	// new 一个数组
	arr_j := new([10]int)
	arr_j[3] = 66

	// 多维数组
	arr_k := [3][4]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr_a) // [0 0]
	fmt.Println(arr_b) // [0 0 0]
	fmt.Println(arr_c) // [ ]
	fmt.Println(arr_d) // [22 0]
	fmt.Println(arr_e) // 最后一位为22
	fmt.Println(arr_f) //
	fmt.Println(arr_g) // 索引及 php 数组中的 key
	fmt.Println(arr_h) // &[22 0]
	fmt.Println(arr_i) // [0xc4200160a8 0xc4200160c0]
	fmt.Println(arr_j) // &[0 0 0 66 0 0 0 0 0 0]
	fmt.Println(arr_k) // [[1 2 3 0] [4 5 6 0] [0 0 0 0]]

	// 冒泡排序
	arr_l := [...]int{266, 4, 8, 11, 77, 44, 1}
	arr_l_len := len(arr_l)
	for i := 0; i < arr_l_len; i++ {
		for j := i + 1; j < arr_l_len; j++ {
			if arr_l[i] < arr_l[j] {
				temp := arr_l[i]
				arr_l[i] = arr_l[j]
				arr_l[j] = temp
			}
		}
	}
	fmt.Println(arr_l) // [266 77 44 11 8 4 1]

}
