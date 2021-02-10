package main

import "fmt"

//控制皇后问题
const Num  = 8
//次数
var count int = 1

var queens[Num][Num] int


func show(){
	fmt.Printf("第%d种写法\n", count)
	for i:=0;i<Num;i++{
		for j:=0;j<Num;j++{
			if queens[i][j]==1{
				fmt.Printf(" %2s","△")
			}else{
				fmt.Printf(" %2s","□")
			}
		}
		fmt.Println()
	}
}

func setQueue(row, col int)bool{
	//第一个放入
	if row == 0{
		queens[row][col]=1
		return true
	}

	for i:=0;i<Num;i++{
		//列有一个为1，无法放下
		if queens[row][i]==1{
			return false
		}
	}

	for i:=0;i<Num;i++{
		//列有一个为1，无法放下
		if queens[i][col]==1{
			return false
		}
	}

	for i,j:=row,col;i<Num && j<Num;i,j=i+1,j+1{
		//对角线无法放下
		if queens[i][j]==1{
			return false
		}
	}

	for i,j:=row,col;i>=0 && j>=0;i,j=i-1,j-1{
		//对角线无法放下
		if queens[i][j]==1{
			return false
		}
	}

	for i,j:=row,col;i<Num && j>=0;i,j=i+1,j-1{
		//对角线无法放下
		if queens[i][j]==1{
			return false
		}
	}

	for i,j:=row,col;i>=0 && j<Num;i,j=i-1,j+1{
		//对角线无法放下
		if queens[i][j]==1{
			return false
		}
	}

	queens[row][col]=1
	return true
}

func solveQueue(row int){
	if row==Num{
		show()
		count++
		return
	}

	for i:=0;i<Num;i++{
		if setQueue(row, i){
			//循环
			solveQueue(row + 1)
		}
		queens[row][i]=0
	}
}


func main() {

	solveQueue(0)
}
