package main

import (
	"fmt"
	"math/rand"
)

func FinsherShuffle(arr []int)[]int{

	newarray := make([]int, 0)

	for {
		Loop:
			length := len(arr)
		if length==0{
			break
		}

		for i := 0; i<=length;i++{
			p := rand.Uint64()%uint64(length)
			fmt.Println(p)
			newarray=append(newarray, arr[p])
			arr = append(arr[0:p], arr[p+1:]...)
			goto Loop
		}
	}

	return newarray

}



func main() {

	arr := []int {1,2,3,4,5,6,7,8,9,0}

	fmt.Println(FinsherShuffle(arr))
}
