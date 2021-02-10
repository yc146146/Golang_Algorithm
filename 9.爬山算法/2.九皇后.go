package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type HillClimbingSolver interface {
	New (size int)HillClimbingSolver
	Successor()HillClimbingSolver
	Objective()bool
}

type Queen struct {
	board []int
}


func MakeQueen(size int)Queen{
	qq := Queen{make([]int, size)}
	for i:=0;i<size;i++{
		qq.board[i]=i
	}
	qq.MixBoard()
	return qq
}

func (q*Queen)RandInt()int{
	return rand.Intn(len(q.board))
}

func (q*Queen)SwapTwo(){
	first:=q.RandInt()
	second := q.RandInt()
	q.board[first],q.board[second]=q.board[second],q.board[first]
}

func (q*Queen)MixBoard(){
	for i:=0;i<len(q.board);i++{
		q.SwapTwo()
	}
}

//返回棋盘的宽度
func (q*Queen)BoardSize()int{
	return len(q.board)
}


//深拷贝棋盘
func (q*Queen)duplicate()Queen{
	newBorad := make([]int, len(q.board))
	for i:=0;i<len(q.board);i++{
		newBorad[i]=q.board[i]
	}
	return Queen{newBorad}
}


//判断皇后之间是否威胁,判断是否对角斜
func (q *Queen)IsDanger(first int, second int)bool{
	return q.board[first]-first==q.board[second]-second||q.board[first]+first==q.board[second]+second||q.board[first]==q.board[second]
}

//威胁的数量
func (q *Queen)Heuristic()int{
	threads:=0
	for i:=0;i<len(q.board);i++{
		for j:=i+1;j<len(q.board);j++{
			if q.IsDanger(i,j){
				threads++
			}
		}
	}
	return threads
}



func (q Queen)New (size int)HillClimbingSolver{
	return MakeQueen(size)
}

func (q Queen) successor()Queen{
	listSize := len(q.board)*2
	//当前的危险
	cur := q.Heuristic()

	for i:=0;i<listSize;i++{
		new:=q.duplicate()
		new.SwapTwo()
		if new.Heuristic()<=cur{
			return new
		}
	}

	return q
}

func (q Queen)Successor()HillClimbingSolver{
	return q.successor()
}
func (q Queen)Objective()bool{
	return q.Heuristic()==0
}



//解决爬山的标砖接口方法
func SolveWithHillCliming(size int,h HillClimbingSolver)HillClimbingSolver{
	current := h.New(size)
	for {
		for i:=0;i<size*5;i++{
			//当前成功
			successor := current.Successor()
			//无线必进
			if !reflect.DeepEqual(successor,current){
				//无限迭代循环
				current=successor
			}else{
				break
			}
		}

		if current.Objective(){
			return current
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	size := 8
	fmt.Println(size)

	res := SolveWithHillCliming(size,MakeQueen(size))

	result := Queen(res.(Queen))
	fmt.Println(result)
	fmt.Println(result.Heuristic())
}