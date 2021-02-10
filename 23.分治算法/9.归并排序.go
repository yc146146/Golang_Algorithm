package main

import "fmt"

func InsertSortX(arr [] int)[]int{
	//数组长度
	length := len(arr)
	if length <= 1{
		//一个元素的数组,直接返回
		return arr
	}else{
		//跳过第一个
		for i:=1;i<length;i++{
			backup := arr[i]
			j:=i-1
			for j>=0 && backup<arr[j]{
				//从前往后移动
				arr[j+1] = arr[j]
				j--
			}
			//插入
			arr[j+1] = backup
			fmt.Println(arr)
		}


		return arr
	}
}



func merge(leftarr []int, rightarr []int) []int {
	leftindex := 0  //左边索引
	rightindex := 0 //右边索引
	lastarr := []int {}
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		}else{
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		}
	}


	//把没有结束的归并过来
	for leftindex < len(leftarr){
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}

	for rightindex < len(rightarr){
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr[] int) []int {
	length := len(arr)
	if length <= 1 {
		return arr

	}else if length>1&&length<5{
		return InsertSortX(arr)
	}else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		return merge(leftarr, rightarr)
	}
}

func MergeSort2(arr[] int) []int {
	length := len(arr)
	if length <= 1 {
		return arr

	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		return merge(leftarr, rightarr)
	}
}

func main() {

	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	//fmt.Println(HeapSortMax(arr))
	fmt.Println(MergeSort(arr))

}
