package main

import (
	"fmt"
)

func main() {
	var num1 = 10

	fmt.Println(num1 == 10)          // true
	fmt.Println(1 == 1 && 2 == 2)    // true
	fmt.Println(1 ^ 2)               // 3 ^是二元运算符
	fmt.Println(^2)                  // -3 ^是一元运算符
	fmt.Println(1 << 10)             // 1024  << 1左移10位，2进制
	fmt.Println(1 << 10 << 10)       // 1048576
	fmt.Println(1 << 10 << 10 >> 10) // 1048576

	/*位运算
		6  : 0110
	    11 : 1011
	    ---------
	     &   0010
	     |   1111
	     ^   1101
	     &^  0100
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

}
