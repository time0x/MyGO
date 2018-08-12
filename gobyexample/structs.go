package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Go 的结构体 是各个字段字段的类型的集合")

	fmt.Println(person{"time", 20})
	fmt.Println(person{name: "bitcoin", age: 10})
	fmt.Println("省略的字段将被初始化为零值。")
	fmt.Println(person{name: "xmr"})
	fmt.Println("& 前缀生成一个结构体指针。")
	fmt.Println(&person{name: "cet", age: 1})

	p := person{name: "bitcoin", age: 10}
	fmt.Println("get p.name:", p.name)

	cp := &p
	fmt.Println("cp是 p 的指针拷贝,返回的是p的地址:", cp.age)

	cp.age = 2
	fmt.Println("set cp.age = 2,cp.age=", cp.age)
	fmt.Println("set cp.age = 2,p.age=", p.age)

	*cp = person{"blockchain", 10}
	fmt.Println("cp 是p的指针拷贝，如果要将cp的内存存放新的对象下")
	fmt.Println("cp ：", cp)
	fmt.Println("*cp ：", *cp)
}
