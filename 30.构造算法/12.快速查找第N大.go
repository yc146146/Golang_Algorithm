package main

import "fmt"


func findKlargest(arr[] int, k int)int{
	return findKlargestgo(arr, 0, len(arr)-1, k)
}

func findKlargestgo(arr []int, left int, right int, k int)int{
	if left >= right{
		return arr[left]
	}

	query := partition(arr, left, right)
	if query + 1 == k{
		return arr[query]
	}

	return -1
}


func main() {
	arr := []int {3,9,2,8,1,7,4,6,5,10}

	fmt.Println(arr)
}
