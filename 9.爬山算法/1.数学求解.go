package main

import (
	"fmt"
	"math"
	"math/rand"
)

//函数指针
type Mountain func(x float64) float64

const (
	workers  int     = 100
	max      float64 = 10
	stepSize float64 = 0.001
	down     float64 = -0.001
)

func Climb(done chan float64, f Mountain) {
	seed := rand.Float64() * max
	right := seed + stepSize
	//左右随机试探 限定一定范围之内
	left := seed - stepSize

	for f(seed) < f(right) || f(seed) < f(left) {
		if f(seed) < f(right) {
			//右边
			seed = right
		} else {
			//左边
			seed = left
		}
		right = seed + stepSize
		left = seed - stepSize
	}

	done <- seed
}

func ArrangeWorkers(name string, fn Mountain) {

	//100个管道
	localmaxes := make(chan float64, workers)
	for i := 0; i < workers; i++ {
		//并发100个人爬山
		go Climb(localmaxes, fn)
	}

	//迭代式取出100个数据
	globalMax := float64(-100)
	localMax := float64(-99)
	for i := 0; i < workers; i++ {
		//从管道取出数据
		localMax = <-localmaxes
		if localMax > globalMax {
			//赋值
			globalMax = localMax
		}
	}
	fmt.Println("取出数据", name, globalMax)

	//关闭管道
	defer close(localmaxes)

}

//-1,1
func sin(x float64) float64 {
	return math.Sin(x)
}

func cos(x float64) float64 {
	return math.Cos(x)
}

func tan(x float64) float64 {
	return math.Tan(x)
}

//最快速度找到一个解
func main() {
	fmt.Println(sin(1.3))

	ArrangeWorkers("math.Sin", sin)
	ArrangeWorkers("math.Cos", cos)
	ArrangeWorkers("math.Tan", tan)

	fmt.Println((sin(7.854477651926955)))
	fmt.Println((cos(12.566748149871607)))
	fmt.Println((tan( 10.995501257174165)))

}
