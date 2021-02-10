package main

import "fmt"

func Sqrt(x float64)float64{
	const E = 0.000001
	z := float64(1)
	for z*z-x > E || z*z-x < -E {
		z = (z+x/z)/2
	}

	return z

}

func main() {
	fmt.Println(Sqrt(2))
}