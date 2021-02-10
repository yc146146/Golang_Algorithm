package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	//任务数量100
	Tasknum = 100
	//计算节点为10
	NodeNum = 10
	TasktimeLengthMin=10
	//任务最大时间
	TasktimeLengthMax=100
	NodeSpeedMin=10
	//最大时间
	NodeSpeedMax=100

	//迭代次数
	IteratorNum=10000
	//染色体
	ChromosomeNum=20
	//复制染色体的比例
	CopyPrecent=0.3
	//生物编译组合
	CrossoverNum=ChromosomeNum*(1-CopyPrecent)
)

func randomIntSlice(length, min, max int)[]int{
	m := make([]int, length)
	for i:=0;i<length;i++{
		m[i]=rand.Intn(max-min)+min
	}
	return m
}




var tasks []int
var nodes []int

//生成新一代染色体
func createGeneration(chromosomeMatrix [][]int, selectProbility []float64)[][]int{
	//新建矩阵
	newchromosomeMatrix := make([][]int, ChromosomeNum)
	//第一代
	if chromosomeMatrix != nil{
		for i := 0; i < ChromosomeNum; i++ {
			newchromosomeMatrix[i] = make([]int, Tasknum)
			for j:=0;j<Tasknum;j++{
				//复制初始化
				newchromosomeMatrix[i][j]=rand.Intn(NodeNum)
			}
		}
		return newchromosomeMatrix
	}


	//交叉
	//newchromosomeMatrix =
		crossover(chromosomeMatrix, selectProbility)
	//变异
	//newchromosomeMatrix =
		mutation(chromosomeMatrix)

	//复制
	newchromosomeMatrix=copy(newchromosomeMatrix,chromosomeMatrix,selectProbility)
	return newchromosomeMatrix

}


//交叉
func crossover(chromosomeMatrix [][]int, selectProbility []float64)(newchromosomeMatrix[][]int){
	if chromosomeMatrix == nil{
		return
	}
	for i:=0;i<CrossoverNum;i++{
		//提取富足
		Chromosomedie := chromosomeMatrix[rws(selectProbility)]
		Chromosomeba := chromosomeMatrix[rws(selectProbility)]
	//	设置随机值
		index := rand.Intn(Tasknum)
		var chromoSon []int
		chromoSon=append(chromoSon,Chromosomedie[:index]...)
		//染色体叠加
		chromoSon=append(chromoSon,Chromosomeba[:index]...)
		newchromosomeMatrix = append(chromosomeMatrix, chromoSon)
	}
	return
}


//生物变异
func mutation(chromosomeMatrix [][]int)(newchromosomeMatrix[][]int){
	if chromosomeMatrix == nil{
		return
	}
	index:= rand.Intn(CrossoverNum)
	taskindex := rand.Intn(Tasknum)
	nodeIndex := rand.Intn(NodeNum)
	chromosomeMatrix[index][taskindex] = nodeIndex
	return chromosomeMatrix
}

//复制 其中一部分复制老的染色体
func copy(chromosomeMatrix [][]int,OldchromosomeMatrix [][]int, selectProbility []float64)(newchromosomeMatrix[][]int){
	if chromosomeMatrix == nil{
		return
	}
	indexs:= maxn(selectProbility,ChromosomeNum-CrossoverNum)
	for _,i:=range indexs{
		chromosomeMatrix = append(chromosomeMatrix,OldchromosomeMatrix[i])
	}
	return chromosomeMatrix
}

//找到适应度最高的
func maxn(selectProbability []float64, n int)(indexs []int){
	indexs = make([]int,0,0)
	//开普
	m := make(map[float64]int)
	for k,v := range selectProbability{
		m[v] = k
	}
	//保存切片
	var keys[]float64
	for k:=range m{
		keys=append(keys,k)
	}
	sort.Float64s(keys)
	if len(keys)!=0{
		for i:=0;i<n;i++{
			indexs = append(indexs,m[keys[len(keys)-i-1]])
		}
	}


	return
}
//轮盘赌博
func rws(selectProbability []float64)(index int){
	sum := 0.0
	r := rand.Float64()
	for index < len(selectProbability){
		sum += selectProbability[index]
		if sum >= r{
			break
		}
		index++
	}
	return
}

//计算时间
func calctime(chromosomeMatrix [][]int)float64{
	min := 0.0
	for i:=0;i<len(chromosomeMatrix);i++{
		sum := 0.0
		for j:=0;j<len(chromosomeMatrix[0]);j++{
			nodeindex := chromosomeMatrix[i][j]
			sum += float64(tasks[j])/float64(nodes[nodeindex])
		}
		if min == 0.0 ||sum<min{
			min = sum
		}
	}
	return min
}
//计算适应度
func calAdapt(chromosomeMatrix [][]int)(selectProbability []float64){
	//数组适应度
	var adaptability[]float64
	for i:=0;i<len(chromosomeMatrix);i++{
		sum := 0.0
		for j:=0;j<len(chromosomeMatrix[0]);j++{
			nodeindex := chromosomeMatrix[i][j]
			sum += float64(tasks[j])/float64(nodes[nodeindex])
		}
		//计算比例
		adaptability = append(adaptability,1.0/sum)
	}
	//计算基数
	total := 0.0
	for _,v := range adaptability{
		total  +=v

	}
	//计算选择概率
	for _,v :=range adaptability{
		selectProbability = append(selectProbability, v/total)
	}
	return
}


//开始迭代
func gosearch(IteratorNum, ChromosomeNum int){
	chromoMatrix := createGeneration(nil, nil)
	fmt.Println("第0代计算消耗时间", calctime(chromoMatrix))
	for i:=0;i<IteratorNum;i++{
		selectPro := calAdapt(chromoMatrix)
		chromoMatrix=createGeneration(chromoMatrix, selectPro)

		fmt.Printf("第%d代计算消耗时间%f\n",i+1,calctime(chromoMatrix))
		fmt.Println(chromoMatrix)
	}
}

func main() {
	tasks = randomIntSlice(Tasknum, TasktimeLengthMin,TasktimeLengthMax)
	nodes = randomIntSlice(NodeNum, NodeSpeedMin,NodeSpeedMax)
	gosearch(IteratorNum, ChromosomeNum)
}




