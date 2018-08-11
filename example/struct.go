package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	fmt.Println(person{}) //{ 0}

	//初始化写法，最好使用& 是使用指针
	p := &person{
		Name: "satoshi",
	}
	p.Age = 66

	fmt.Println(p)

	func_a(p)
	fmt.Println(p)

	//匿名写法
	p1 := struct {
		Name string
		Age  int
	}{
		Name: "time",
		Age:  20,
	}

	fmt.Println(p1)

	// 嵌入结构
	teach := teacher{human: human{Sex: 0}, Name: "tom", Age: 30}
	stude := student{human: human{Sex: 1}, Name: "lily", Age: 10}

	teach.Name = "tom2"
	teach.Sex = 2
	teach.human.Sex = 2 // 这样的写法是避免字段的冲突
	fmt.Println(teach, stude)

}

func func_a(p *person) {
	p.Name = "time"
	fmt.Println("func_a:", p)
}
