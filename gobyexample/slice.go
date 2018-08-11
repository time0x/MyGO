package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("slice 是一个比数组更加强大的序列接口")

	s := make([]string, 3)
	fmt.Println("slice s:", s)
	fmt.Println("s type is :", reflect.TypeOf(s)) // []string,数组的类型类似  [5]int

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set s:", s)
	fmt.Println("get s[2]:", s[2])
	fmt.Println("s len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append some :", s)

	c := make([]string, len(s))
	copy(c, s) //
	fmt.Println("c is copy s:", c)

	s[3] = "change"
	fmt.Println("change s:", s)
	fmt.Println("now c :", c)

	l := s[2:5]
	fmt.Println("s[2:5]:", l)

	l = s[:5]
	fmt.Println("s[:5]:", l)

	l = s[3:]
	fmt.Println("s[3:]:", l)

	t := []string{"time", "food", "air"}
	fmt.Println("slice t:", t)

	twoS := make([][]int, 3)
	for i := 0; i < len(twoS); i++ {
		innerLen := i + 1
		twoS[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoS[i][j] = i + j
		}
	}
	fmt.Println("Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同")
	fmt.Println("twoS :", twoS)
}
