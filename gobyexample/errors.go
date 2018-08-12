package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 55 {
		return -1, errors.New("can't work with 55")
	}

	return arg + 5, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 55 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 5, nil
}

func main() {
	fmt.Println("错误通常是最后一个返回值并且是 error 类型，一个内建的接口,返回错误值为 nil 代表没有错误。")
	for _, i := range []int{3, 66, 55} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{3, 66, 55} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	fmt.Println("通过实现 Error 方法来自定义 error 类型是可以的。这里使用自定义错误类型来表示上面的参数错误")
	_, e := f2(55)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
