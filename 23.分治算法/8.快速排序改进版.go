package main

import (
	"fmt"
	"math/rand"
)

func SortForMerge(arr[]int ,left int,right int){
	for i:=left;i<=right;i++{
		//备份数据
		temp := arr[i]
		var j int
		//定位
		for j=i;j>left && arr[j-1]>temp;j--{
			//数据往后移动
			arr[j] = arr[j-1]
		}
		//插入
		arr[j] = temp

	}
}

func swap(arr[]int, i int, j int){
	arr[i], arr[j] = arr[j], arr[i]
}

//递归的快速【爱旭
func QuickSortXX(arr[]int ,left int,right int){
	//数组剩下3个数，直接插入排序
	if right - left < 3{
		SortForMerge(arr, left, right)
	}else{
		//	随机最后一个数字, 放在第一个位置
		swap(arr, left, rand.Int()%(right-left+1)+left)
		//坐标数组
		vdata := arr[left]
		lt := left
		gt := right + 1
		i := left+1
		for i<gt{
			if arr[i]<vdata{
				swap(arr, i, lt+1)
				lt++
				i++
			}else if arr[i]>vdata{
				swap(arr, i, gt-1)
				gt--

			}else{
				i++
			}
		}
		//交换头部的位置
		swap(arr, left, lt)
		//递归处理小于那一段
		QuickSortXX(arr, left, lt - 1)
		//递归处理大于那一段
		QuickSortXX(arr, gt,  right)
	}
}

//快速排序的核心程序
func QuickSortPlus(arr []int){
	QuickSortXX(arr, 0, len(arr)-1)
}


func main() {
	arr := []int{3,9,2,8,1,7,4,6,5,10}
	fmt.Println("未排序", arr)
	QuickSortPlus(arr)
	fmt.Println("已排序", arr)
}
