package main

import "fmt"

//20%25 20
func gcdx(x,y int)int{
	var tmp int
	for {
		tmp = x%y
		fmt.Println(tmp)
		if tmp >0 {
			x = y
			y = tmp
		}else{
			return y
		}
	}
}

func gcd(x,y int) int{
	var n int
	if x>y{
		n=y
	}else{
		n=y
	}

	for i:=n;i>=1;i--{
		if x%i==0&&y%i==0{
			return i
		}
	}
	return 1

}

func main() {
	x,y := 20,25

	fmt.Println(gcdx(x,y))
	fmt.Println(gcd(x,y))
	fmt.Println(x*y/gcdx(x,y))

}
