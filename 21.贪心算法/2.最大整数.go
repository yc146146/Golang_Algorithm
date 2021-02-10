package main

import (
	"fmt"
	"strconv"
)

func GetMax(arr[]int)int{
	max := arr[0]
	for i:=1;i<=len(arr)-1;i++{
		if arr[i]>max{
			max=arr[i]
		}
	}
	return max
}

//取得一个整数的位数
func GetLength(num int)int{
	var length=0
	for i:=num;i!=0;i/=10{
		length++
	}
	return length
}

func GetMaxN(arr[] int)[]int{

	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if strconv.Itoa(arr[i]) < strconv.Itoa(arr[j]){
				arr[i],arr[j]=arr[j],arr[i]
			}
		}
	}


	return arr
}


func main() {
	arr:= []int{127,8,19,4,13}
	arr=GetMaxN(arr)
	fmt.Println(arr)
	fmt.Println(GetLength(GetMax(arr)))
}

