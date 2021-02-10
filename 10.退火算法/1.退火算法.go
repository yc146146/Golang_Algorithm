package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//随机实数
func RandomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func f(x1, x2 float64) float64 {
	sin := math.Sin(x1)
	cos := math.Cos(x2)
	sqrt := math.Sqrt(math.Pow(x1, 2)+math.Pow(x2, 2)) / math.Pi
	exp := math.Exp(math.Abs(1 - (sqrt)))

	return -math.Abs(sin * cos * exp)

}

//指数
func e(delD, temp float64) float64 {
	return math.Exp(-delD / temp)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var T float64 = 100
	var coolingRate float64 = 0.11382
	var x1 float64 = RandomFloat(-10, 10)
	var x2 float64 = RandomFloat(-10, 10)
	var initalState float64 = f(x1, x2)
	var currentState float64 = initalState
	var bestsofar float64 = currentState
	for T >= 0.11382 {
		x1 = RandomFloat(-10, 10)
		x2 = RandomFloat(-10, 10)
		var newState float64 = f(x1, x2)
		var delD float64 = newState - currentState
		if delD < 0 {
			currentState = newState
			bestsofar=newState
			T=T-coolingRate
		} else if delD > 0 {
			var e float64=e(delD,T)
			var R float64=RandomFloat(0,1)
			if R<e{
				newState=currentState
				bestsofar=newState
				T=T-coolingRate
			}
		}
	}
	fmt.Println(currentState)
	fmt.Println(bestsofar)
}
