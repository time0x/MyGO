package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name string
}

type TEST int

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	fmt.Println(a.age) //可以访问此处的私有字段

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var c TEST
	c.Print()
	(TEST).Print(c)
}

// 引用才可以修改，即传递指针
func (a *A) Print() {
	a.Name = "a name"
	a.age = 10
	fmt.Println("A")
}

// 只是副本，修改无效
func (b B) Print() {
	b.Name = "b name"
	fmt.Println("B")
}

//
func (c TEST) Print() {
	fmt.Println("TEST")
}
