package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string

	data = `STATUS=S,When=1522848281,Code=11,Msg=Summary,Description=bmminer 1.0.0|SUMMARY,E
lapsed=279527,GHS 5s=13029.30,GHS av=13158.31,Found Blocks=0,Getworks=7152,Accep
ted=33856,Rejected=3,Hardware Errors=22908,Utility=7.27,Discarded=147326,Stale=3
,Get Failures=8,Local Work=13543063,Remote Failures=2,Network Blocks=456,Total M
H=3678089971109.0000,Work Utility=185395.04,Difficulty Accepted=863617024.000000
00,Difficulty Rejected=98304.00000000,Difficulty Stale=0.00000000,Best Share=166
8460548,Device Hardware%=0.0027,Device Rejected%=0.0114,Pool Rejected%=0.0114,Po
ol Stale%=0.0000,Last getwork=1522848280|`

	temp := strings.Split(data, "|")
	fmt.Println(len(temp)) // ["a" "b" "c"]
	temp1 := strings.Split(temp[1], ",")
	num := len(temp1)
	fmt.Println(temp1)
	for i := 0; i < num; i++ {
		fmt.Println(temp1[i])
	}

}
