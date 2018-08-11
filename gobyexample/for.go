package main

import (
	"fmt"
)

func main() {
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("first for,i=", i)

	for {
		i++
		fmt.Println(i)
		if i == 5 {
			break
		}

	}

	fmt.Println("second for,i=", i)

	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	fmt.Println("third for,i=", i)
}
