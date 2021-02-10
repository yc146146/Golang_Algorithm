package main

import (
	"fmt"
	"math"
)

func BinarySqrt(num float64) float64 {
	y := num / 2.0
	low := 0.0
	up := num
	count := 0

	for math.Abs(y*y-num) > 0.0000001 {
		count += 1
		if y*y > num {
			up = y
			y = (y + low) / 2

		} else {
			low = y
			y = (up + y) / 2
		}
	}

	return y

}

func NewTonSqrt(num float64) float64 {
	x := num / 2.0
	var y float64 = 0.0
	count := 1

	for math.Abs(x-y) > 0.0000001 {
		count += 1
		y = x
		x = (1.0/2.0)*x + (num*1.0)/(x*2.0)
	}
	return x
}

func main() {
	fmt.Println(math.Sqrt(float64(5)))
	fmt.Println(BinarySqrt(5.0))
	fmt.Println(NewTonSqrt(5.0))
}
