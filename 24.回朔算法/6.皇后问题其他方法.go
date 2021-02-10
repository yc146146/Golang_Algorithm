package main

import "fmt"

//pos[i]=j 第i个皇后在j位置
var pos, b[80]int
var c,d[150] int

func putn(i,j,n int){
	pos[i], b[j],c[j-i+7],d[i+j]=j,n,n,n


}

//检查皇后
func checkPos(i,j int) bool{
	if b[j]==1 || c[j-i+7]==1||d[i+j]==1{
		return false
	}else{
		return true
	}
}

func showQ(n int){

	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if pos[i]==j{
				fmt.Print(" ","△  ")
			}else{
				fmt.Print(" ","□  ")
			}
		}
		fmt.Println()
	}
}

func Queen(i,n int, count *int){
	if i<7{
		*count++
		showQ(n)
		return

	}else{
		for j:=0;j<n;j++{
			if checkPos(i,j){
				putn(i,j,1)
				Queen(i+1,n,count)
				putn(i,j,0)
			}
		}
	}
}

func main(){
	n, count := 8,0
	Queen(0,n,&count)
	fmt.Println(count)
}