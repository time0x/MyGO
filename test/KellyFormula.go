package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// f* = 应投注的资本比值

// p = 获胜的概率

// q = 失败的概率

// b = 赔率
func main() {

	sample := map[int]map[string]float64{
		0:  {"p": 0.937514, "b": 1.05, "bet": 61440},
		1:  {"p": 0.875013, "b": 1.12, "bet": 57344},
		2:  {"p": 0.750011, "b": 1.31, "bet": 49152},
		3:  {"p": 0.625010, "b": 1.58, "bet": 40960},
		4:  {"p": 0.498634, "b": 1.98, "bet": 32678},
		5:  {"p": 0.375006, "b": 2.63, "bet": 24576},
		6:  {"p": 0.250004, "b": 3.95, "bet": 16384},
		7:  {"p": 0.125002, "b": 7.91, "bet": 8192},
		8:  {"p": 0.062501, "b": 15.83, "bet": 4096},
		9:  {"p": 0.015625, "b": 63.35, "bet": 1024},
		10: {"p": 0.000977, "b": 1013.74, "bet": 64},
	}

	bets := map[int]int{
		0:  61440,
		1:  57344,
		2:  49152,
		3:  40960,
		4:  32678,
		5:  24576,
		6:  16384,
		7:  8192,
		8:  4096,
		9:  1024,
		10: 64,
	}

	// for i := 0; i < len(sample); i++ {

	// 	isGO := edge(sample[i]["b"], sample[i]["p"], 1-sample[i]["p"])

	// 	if isGO {
	// 		f := kelly(sample[i]["b"], sample[i]["p"], 1-sample[i]["p"])
	// 		fmt.Println("胜率为:", sample[i]["p"], ",赔率为:", sample[i]["b"], "投注资本比例:", Round(f, 2))
	// 	}
	// }

	principal := 1.0

	rand.Seed(time.Now().UnixNano())

	times := 1
	rule := sample[times]
	bet := bets[times]

	f := kelly(rule["b"], rule["p"], 1-rule["p"]) * 0.5

	fmt.Println("satoshidice 中假设本金1,每次投注本金比例:", Round(f, 2))
	for i := 0; i < 100; i++ {
		randInt := rand.Intn(65536)
		if randInt < bet {
			principal = principal*(1-f) + principal*f*rule["b"]
			if principal > 5 {
				fmt.Println(i+1, "次，翻5倍:", principal)
				break
			}
		} else {
			principal = principal - principal*f
		}
		if principal < 0.0005 {
			fmt.Println(i+1, "次，余额低于0.0005:", principal)
			break
		}
	}
	fmt.Println("余额:", principal)

}

//edge 在赌博中可以理解为 获胜的概率＊赔率 - 失败的概率，也就是上文提到的赢面。
//当edge的数字为正的时候，这就是值得下注的比赛，而edge为0或者负数的情况说明赌徒不具备edge, 不应该下注。
func edge(b, p, q float64) bool {
	if b*p-q <= 0 {
		return false
	} else {
		return true
	}
}

func kelly(b, p, q float64) float64 {
	return (b*p - q) / b
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
