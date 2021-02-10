package main

import "math"

//解决问题 设定三个参数
type Param struct {
	w []float64
	c1 []float64
	c2 []float64
}

//新建一个区间
func NewParam(w, c1, c2 []float64)*Param{
	switch {
	case w==nil:
		panic("w can not be nil")
	case c1==nil:
		panic("c1 can not be nil")
	case c2==nil:
		panic("c2 can not be nil")
	case len(w)!=len(c1) || len(w)!=len(c2):
		panic("c1 length must be c2 length")
	}
	return &Param{w,c1,c2}
}

//整体判断所有元素是否在一个区间内
func (r * Range)In(vector []float64)bool{
	if vector==nil{
		panic("vector can not be nil")
	}

	if len(vector)!=len(r.min){
		panic("length not equal")
	}

	for i:=range vector{
		if vector[i]<r.min[i]||vector[i]>r.max[i]{
			return false
		}
	}

	return true
}

//提取最大
func (r * Range)Min()[]float64{
	if r.min==nil{
		return nil
	}
	cpy:=make([]float64, len(r.min))
	copy(cpy,r.min)
	return cpy
}

func (r * Range)Max()[]float64{
	if r.max==nil{
		return nil
	}
	cpy:=make([]float64, len(r.max))
	copy(cpy,r.max)
	return cpy
}

func(p*Param)W()[]float64{
	return p.w
}

func(p*Param)C1()[]float64{
	return p.c1
}

func(p*Param)C2()[]float64{
	return p.c2
}

//函数指针
type TargetFunc func(vector []float64)float64

type Solver struct {
	//函数指针
	f TargetFunc
	//例子群
	particles []*Particle
	//参数
	param *Param
	best []float64
}

//新建一个问题解决类工具对象
func NewSolver(f TargetFunc, particles []*Particle,param *Param)*Solver{
	if particles == nil{
		panic("particles无法为空")
	}else if len(particles)<=0{
		panic("数量错误")
	}

	var bestValue float64
	var best []float64

	for _, p := range particles{
		if p==nil{
			continue
		}

		if best==nil{
			best = p.position
			bestValue=f(best)
		}else if f(p.position)<bestValue{
			copy(best, p.position)
			bestValue=f(best)
		}
	}
	return &Solver{f,particles,param,best}
}

func (s * Solver)TargetFunc()TargetFunc{
	return s.f
}

func (s * Solver)Particles()[]*Particle{
	return s.particles
}

func (s * Solver)Param()*Param{
	return s.param
}

func (s * Solver)Best()[]float64{
	if s.best == nil{
		return nil
	}

	cpy := make([]float64, len(s.best))
	copy(cpy, s.best)
	return cpy
}

func (s * Solver)Step(){
	//最优值
	bestValue := s.f(s.best)
	for _,p :=range s.particles{
		if p==nil{
			continue
		}
		if s.best == nil{
			s.best = p.Best()
			bestValue=s.f(s.best)
		}
		p.Step(s.f, s.param, s.best)
		if p.EvalValue()< bestValue{
			copy(s.best, p.best)
			//保存最优
			bestValue=p.EvalValue()
		}
	}
}

func (s * Solver)Run(errorValue float64, maxcount uint){
	pre := math.MaxFloat64
	for i:=uint(0);i<maxcount;i++{
		s.Step()
		//取得速度
		v:=s.f(s.best)
		if math.Abs(pre-v) <= errorValue{
			break
		}
		//继续循环
		pre = v
	}
}