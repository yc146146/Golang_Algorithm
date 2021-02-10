package main

import "fmt"

//优先使用面值最大
const N = 5

var Money = []int{5, 2, 2, 3, 5}
var Value = []int{1, 5, 10, 50, 100}

func Min(a, b int) int {
	if a < b {
		return a

	} else {
		return b
	}
}

func Pay(money int) int {
	num := 0
	for i := N - 1; i >= 0; i-- {
		c := Min(money/Value[i], Money[i])
		fmt.Println(c, Value[i])
		money = money - Value[i]*c
		num += c
	}

	if money > 0 {
		num = -1
	}

	return num
}

func main() {

	fmt.Println(Pay(522))

}
