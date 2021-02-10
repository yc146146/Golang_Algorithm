package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Pos[]float64

func NewPos(x, y float64)Pos{
	return Pos([]float64{x,y})
}

func (p Pos)X()float64{
	return p[0]
}

func (p Pos)Y()float64{
	return p[1]
}



func main() {
	//	x+y-3=0
	//	2x+5y-9=0
	//	x=2 y=1

	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("x+y-3=0")
	fmt.Println("2x+5y-9=0")

	f := func(v []float64)float64{
		p:=Pos(v)
		return math.Pow(p.X()+p.Y()-3,2)+math.Pow(2*p.X()+5*p.Y()-9, 2)
	}

	//构建1000个粒子
	const count = 1000

	//开辟数组
	paricles := make([]*Particle, count)

	for i:=range paricles{
		position := NewPos(float64(rnd.Intn(20)-10), float64(rnd.Intn(20)-10))
		velocity := NewPos(float64(rnd.Intn(20)-10), float64(rnd.Intn(20)-10))

		min := NewPos(-10, -10)
		max := NewPos(10,10)
		valueRange := NewRange(min, max)
		//初始化
		paricles[i] = NewParticle(position, velocity, valueRange)
	}

	w := NewPos(0.9,0.9)
	c1 := NewPos(0.9,0.9)
	c2 := NewPos(0.9,0.9)

	param := NewParam(w,c1,c2)
	//解决对象
	slover := NewSolver(f, paricles, param)
	const STEP = 1000

	slover.Run(0.00001, STEP)
	best := Pos(slover.Best())
	fmt.Println("x=",best.X(),"y=", best.Y())


}
