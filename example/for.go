package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 21; i < 23; i++ {
		var ip = "172.16.9."
		var newip = ip + strconv.Itoa(i)
		fmt.Printf(newip + "\n")
	}
}
