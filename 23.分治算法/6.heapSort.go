package main

import "fmt"

func HeapSortMax(arr []int, length int)[]int{
	//length := len(arr)
	if length <=1{
		return arr
	}else{
		//深度 n 2*n+1 2*n+2
		depth := length/2-1
		for i:=depth;i>=0;i--{
			//假定最大的在i的位置
			topmax := i
			leftchild:=2*i+1
			rightchild:=2*i+2
			if leftchild<=length-1 && arr[leftchild]>arr[topmax]{
				topmax = leftchild //如果左边比我大 记录最大的
			}
			if rightchild<=length-1 && arr[rightchild]>arr[topmax]{
				topmax = rightchild
			}
			//确保i的值就是最大
			if topmax !=i {
				arr[i], arr[topmax] = arr[topmax],arr[i]
			}
		}
	}
	return arr
}

func HeapSort(arr[] int)[]int{
	length := len(arr)
	for i:=0;i<length;i++{
		lastmesslen := length-i
		HeapSortMax(arr, lastmesslen)
		if i<length{
			arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
		}
		fmt.Println("ex:",arr)
	}

	return arr
}


func main() {
	arr := []int {11,9,2,8,3,7,4,6,5,10}
	//fmt.Println(HeapSortMax(arr))
	fmt.Println(HeapSort(arr))
}

