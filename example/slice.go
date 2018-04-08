package main

import (
	"fmt"
)

func main() {
	slice_a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice_a) // [0 0 ...]

	// slice_a := make([],len,cap)
	// fmt.Println(slice_a)

	slice_a_1_5 := slice_a[5:10] // [6 7 8 9 0] 索引 5，6，7，8，9 的数
	fmt.Println(slice_a_1_5)

	slice_a_9 := slice_a[9] // 0 不存在索引10 ，0-9
	fmt.Println(slice_a_9)

	slice_a_4_end := slice_a[4:] // [5 6 7 8 9 0],或者：slice_a[4:len(slice_a)]
	fmt.Println(slice_a_4_end)

	slice_a_start_4 := slice_a[:4] // [1 2 3 4]
	fmt.Println(slice_a_start_4)

	// 正式写法
	slice_b := make([]int, 5, 10)           // 20是容量，切片小于10内存不会重新分配，大于的话会重新分配一块连续的内存空间
	fmt.Println(len(slice_b), cap(slice_b)) // 10,10

	slice_c := make([]int, 5)               // 20是容量，切片小于10内存不会重新分配，大于的话会重新分配一块连续的内存空间
	fmt.Println(len(slice_c), cap(slice_c)) // 5,10

	arr_a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	slice_d := arr_a[2:5]
	fmt.Println(slice_d)                    //[99 100 101] 对应的ASIC码
	fmt.Println(string(slice_d))            // cde
	fmt.Println(len(slice_d), cap(slice_d)) // 3,6 这个切片虽然切的是 cde 但是容量为6，从c-h

	// Reslice 对切片切片
	reslice_a := slice_d[1:2]
	reslice_b := slice_d[0:6]
	fmt.Println(string(reslice_a)) // d
	fmt.Println(string(reslice_b)) // cdefgh

}
