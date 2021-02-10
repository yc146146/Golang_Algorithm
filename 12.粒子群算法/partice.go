package main

import (
	"math"
	"math/rand"
	"time"
)

type Range struct {
	//最小值的集合
	min []float64
	//最大值的集合
	max []float64
}

//新建一个区间
func NewRange(min, max []float64)*Range{
	switch {
	case min==nil:
		panic("min can not be nil")
	case max==nil:
		panic("max can not be nil")
	case len(min)!=len(max):
		panic("min length must be max length")
	}
	return &Range{min,max}
}

type Particle struct {
	//当前的位置
	position []float64

	//当前的速度
	velocity []float64

	//位置的区间
	valuesRange*Range

	//保存全局最优值
	evalValue float64
	//局部最优
	best[]float64

}

//新建粒子群
func NewParticle (position, vecocity []float64, valueRange*Range)*Particle{
	switch {
	case position==nil:
		panic("psoition can not be nil")
	case vecocity==nil:
		panic("vecocity can not be nil")
	case len(position)!=len(vecocity):
		panic("psoition length must be vecocity length")
	}

	cpypos := make([]float64, len(position))
	copy(cpypos, position)

	cpyv := make([]float64, len(vecocity))
	copy(cpyv, position)


	best := make([]float64, len(position))
	copy(best, position)

	return &Particle{cpypos, cpyv,valueRange,math.MaxFloat64,best}

}


func (p *Particle)Position()[]float64{
	return p.position
}

func (p *Particle)Velocity()[]float64{
	return p.velocity
}

func (p *Particle)Range()*Range{
	return p.valuesRange
}


func (p *Particle)EvalValue()float64{
	return p.evalValue
}


//用于计算
func (p *Particle)Best()[]float64{
	if p.best == nil{
		return nil
	}

	cpy := make([]float64, len(p.best))
	return cpy
}

func (p *Particle)Step(f TargetFunc, param *Param, globalBest[]float64){
	switch {
	case f==nil:
		panic("f can not be nil")
	case param==nil:
		panic("param can not be nil")
	case globalBest==nil:
		panic("max can not be nil")
	case len(globalBest)!=len(p.position):
		panic("globalBest length must be position length")
	}

	//老的位置
	Oldposition := make([]float64, len(p.position))
	c1 := param.C1()
	c2 := param.C2()
	//获取数值
	w:=param.W()
	copy(Oldposition,p.position)

	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=range p.position{
		p.position[i]+=p.velocity[i]
		//随机值
		r1,r2 := rnd.Float64(),rnd.Float64()

		p.velocity[i]=w[i]*p.velocity[i]+r1*c1[i]*(p.best[i]-p.position[i])*r2*c2[i]*(globalBest[i]-p.position[i])
	}
	//超越了范围
	if !p.valuesRange.In(p.position){
		copy(p.position, Oldposition)
	}

	//跟新最好的
	p.evalValue = f(p.position)
	//计算最好值
	bestValue := f(p.best)
	if p.evalValue<bestValue{
		copy(p.best, p.position)
	}

}