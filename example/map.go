package main

import (
	"fmt"
	"sort"
)

func main() {
	// 关键字 [key类型]value类型
	var map_a map[int]string
	// 初始化1
	map_a = map[int]string{}
	// 初始化2
	map_a = make(map[int]string)
	fmt.Println(map_a)
	// 初始化3
	var map_b map[int]string = make(map[int]string)
	fmt.Println(map_b)
	// 初始化4
	map_c := make(map[int]string)
	fmt.Println(map_c)

	map_c[1] = "time"
	fmt.Println(map_c)
	fmt.Println(map_c[1]) // time
	fmt.Println(map_c[2]) // 空

	delete(map_c, 1) // 删除

	// 复杂map
	map_d := make(map[int]map[int]string)
	map_d[1] = make(map[int]string) // 初始化key为1的map
	map_d[1][1] = "time"
	fmt.Println(map_d[1][1])

	// 迭代操作,类比 foreach

	slice_a := make([]map[int]string, 5)
	for i, v := range slice_a {
		v = make(map[int]string, 1)
		v[i] = "foreach"
		fmt.Println(v)
	}
	fmt.Println(slice_a)

	map_e := map[int]string{1: "e", 2: "f", 3: "t", 4: "y"}
	slice_map_e := make([]int, len(map_e))

	i := 0
	for k, _ := range map_e {
		slice_map_e[i] = k
		i++
	}
	// 排序
	sort.Ints(slice_map_e)

	// 颠倒 map_e 和 map_f
	map_f := map[string]int{}

	for k, v := range map_e {
		map_f[v] = k
	}

	fmt.Println(map_f)

}
