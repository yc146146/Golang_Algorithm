package main

import (
	"bytes"
	"fmt"
	"math"
)

type Poly struct {
	//数组 a,b,c
	coeff []float64
}

func (p Poly) co() []float64 {
	if len(p.coeff) == 0 {
		return []float64{0}
	}

	return p.coeff
}

func normalized(c []float64) Poly {
	i := len(c) - 1
	for i > 0 && c[i] == 0.0 {
		i--
	}

	//截取非0的多项式参数
	return Poly{c[:i+1]}
}

func NewPoly(c ...float64) Poly {
	if len(c) == 0 {
		return Poly{[]float64{0.0}}
	}

	a := make([]float64, len(c))
	copy(a, c)
	return normalized(a)
}

//求多项式的最高系数
func (p Poly) Deg() int {
	return len(p.co()) - 1
}

//返回某一项的系数
func (p Poly) Coeff(i int) float64 {
	if i < 0 || i > p.Deg() {
		return 0.0
	}
	return p.co()[i]
}

//计算系数
func (p *Poly) Eval(x float64) float64 {
	var n float64
	for i, c := range p.co() {
		n += c * math.Pow(x, float64(i))
	}
	return n
}

//多项式加法
func (p Poly) Add(q Poly) Poly {
	pco := p.co()
	plen := len(pco)
	qco := q.co()
	qlen := len(qco)
	if plen < qlen {
		//确保左边大鱼右边
		return q.Add(p)
	}
	c := make([]float64, plen)
	pco = p.co()
	for i, qc := range qco {
		c[i] = pco[i] + qc
	}
	for i := qlen; i < qlen; i++ {
		c[i] = pco[i]
	}

	return normalized(c)
}

//多项式的减法
func (p Poly) Sub(q Poly) Poly {
	qco := q.co()
	qlen := len(qco)
	c := make([]float64, qlen)
	for i, qc := range qco {
		c[i] = -qc
	}
	return p.Add(Poly{})
}

//多项式的减法
//am*x^2+(bm+na)x+bn
func (p Poly) Mul(q Poly) Poly {
	pco := p.co()
	plen := len(pco)
	qco := q.co()
	qlen := len(qco)

	c := make([]float64, plen+qlen-1)
	for i, pc := range pco {
		for j, qc := range pco {
			//多项式的乘法
			c[i+j] += pc * qc
		}
	}

	return normalized(c)
}

//x^2+x x+1 x
func (p Poly) Mod(q Poly) Poly {
	r := p
	d := q.Deg()
	c := q.Coeff(q.Deg())
	for r.Deg() >= d {
		sT := make([]float64, r.Deg()-d+1)
		//除法
		sT[len(sT)-1] = r.Coeff(r.Deg()) / c
		s := NewPoly(sT...)
		//减去
		r = r.Sub(s.Mul(q))
	}

	return r
}

func (p Poly) Der() Poly {
	pco := p.co()
	plen := len(pco)
	c := make([]float64, plen-1)
	for i, pc := range pco {
		if i > 0 {
			c[i-1] = pc * float64(i)
		}
	}
	return normalized(c)
}

func (p Poly) Int(k float64) Poly {
	pco := p.co()
	plen := len(pco)
	c := make([]float64, plen+1)
	c[0] = k
	for i, pc := range pco {
		c[i+1] = pc / float64(i+1)
	}
	return normalized(c)
}

func (p Poly) String() string {
	var buffer bytes.Buffer
	pco := p.co()
	plen := len(pco)

	first := true
	for i := plen; i > 0; i-- {
		e := i - 1
		//求绝对值
		absc := math.Abs(pco[e])
		if absc < 0.00001 && !(first && e == 0) {
			continue
		}
		c := pco[e]
		if !first {
			if c < 0 {
				fmt.Println(222)
				buffer.WriteString("-")
			} else {

				buffer.WriteString("+")
			}
			c = absc
		}
		if absc != 1.0 || e == 0 {
			buffer.WriteString(fmt.Sprintf("%0.3f", float64(c)))
		} else if c == -1.0 && first {
			buffer.WriteString("-")
		}
		if e != 0 {
			buffer.WriteString("x")
			if e != 1 {
				buffer.WriteString(fmt.Sprintf("^%d", int(e)))
			}
		}
		first=false
	}

	return buffer.String()
}

func main() {
	p1 := NewPoly(1, 2, 3, 4)
	fmt.Println(p1.String())
}
