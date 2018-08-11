package main

import (
	"fmt"
)

func main() {
	fmt.Println("map 在别的语言中称为哈希或字典")

	m := make(map[string]int)

	m["k1"] = 66
	m["k2"] = 168
	fmt.Println("map m:", m)
	fmt.Println("map len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1 :", v1)

	delete(m, "k2")
	fmt.Println("delete k2", m)

	_, prs := m["k2"]
	fmt.Println("当从一个 map 中取值时，可选的第二返回值指示这个键是否在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 ‘’ 而产生的歧义")
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
