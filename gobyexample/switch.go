package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("now is the weekend")
	default:
		fmt.Println("now is weekday")
	}

	h := time.Now().Hour()
	fmt.Println(h)
	switch {
	case h < 12:
		fmt.Println("now is before noon")
	default:
		fmt.Println("now is after noon")
	}

}
