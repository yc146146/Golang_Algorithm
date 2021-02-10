package main

import (
	"fmt"
	"math"
)

//判断质数 2 3 5 7

func CheckisZ(num int)bool{
	if num < 2{
		return false
	}else if num == 2|| num ==3{
		return true
	}else{
		for i:=2;i<=int(math.Sqrt(float64(num)));i++{
			if num%i==0{
				return false
			}
		}
		return true
	}


}


func Makedata(num int){
	for i := 2; i < num; i++ {
		if num%i == 0 {
			if CheckisZ(i){
				fmt.Printf("%d*",i)
				num/=i
			}else{
				num/=i
				Makedata(i)
				fmt.Printf("*")

			}

		}
	}
	if CheckisZ(num){
		fmt.Printf("%d*",num)
	}else{
		Makedata(num)
	}


}


func main() {
	num := 180
	Makedata(num)


	//fmt.Println(CheckisZ(6))

}
