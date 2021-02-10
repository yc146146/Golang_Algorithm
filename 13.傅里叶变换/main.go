package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func isPowerofTwo(n int)bool{
	for n&1==0 && n>1{
		n>>=1
	}
	return n==1
}

//虚数 傅里叶变换
func NativeDFT(input []complex128)[]complex128{
	//求出长度
	n:=len(input)
	//开辟内存
	res := make([]complex128,n,n)

	for k:=0;k<n;k++{
		wk := cmplx.Rect(1, 2*math.Pi*float64(k)/float64(n))
		w:=complex(1,0)
		for i:=0;i<n;i++{
			//叠加波形计算
			res[k] += input[i]*w
			w*=wk
		}
	}

	return res
}










func main() {

	signal := []complex128{29,172,313,13,233,213,21323,44,5,55,532,12}

	for i,X := range NativeDFT(signal){
		x := signal[i]
		xr := real(x)
		Xr := real(X)
		Xi := imag(X)
		fmt.Printf("%d  %f  %f  %f\n", i,xr,Xr,Xi)
	}



}
