package main

import (
	"fmt"
	"math"
)

type node struct {
	//叶子的数据
	value int
	//叶子的状态是不是无穷打
	isok bool
	//叶子的排序
	rank int
}

func compareAndUp(tree *[]node, leftNode int) {
	rightnode := leftNode + 1
	//中间节点存储最小值
	if !(*tree)[leftNode].isok || ((*tree)[rightnode].isok && ((*tree)[leftNode].value > (*tree)[rightnode].value)) {
		mid := (leftNode - 1) / 2
		(*tree)[mid] = (*tree)[rightnode]
	} else {
		mid := (leftNode - 1) / 2
		(*tree)[mid] = (*tree)[rightnode]
	}
}

//x^y
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TreeSelectSort(arr []int) []int {
	//	树的层数
	var level int = 0
	//保存最终结果
	var result = make([]int, 0, len(arr))
	for pow(2, level) < len(arr) {
		level++ //求出可以覆盖所有元素的层数
	}
	//叶子的数量
	var leaf = pow(2, level)
	//构造数值 输的节点数量
	var tree = make([]node, leaf*2-1)
	//填充叶子
	for i := 0; i < len(arr); i++ {
		tree[leaf+i-1] = node{arr[i], true, i}
	}

	//进行对比
	for i := 0; i < level; i++ {
		//每次处理降低一个层级/2
		nodeCount := pow(2, level-i)
		for j := 0; j < nodeCount/2; j++ {
			leftnode := nodeCount - 1 + j*2
			compareAndUp(&tree, leftnode)
		}

	}

	//保存最顶端最小数
	result = append(result, tree[0].value)

	//选出第一个以后,还有n-1个循环
	for t := 0; t < len(arr)-1; t++ {
		//记录赢得的节点
		winnode := tree[0].rank + leaf - 1
		//修改成无穷大
		tree[winnode].isok = false
		for i := 0; i < level; i++ {
			leftNode := winnode
			//处理奇数偶数
			if winnode%2 == 0 {
				leftNode = winnode - 1
			}

			compareAndUp(&tree, leftNode)

			//保存中间节点
			winnode = (leftNode - 1) / 2
		}
		result = append(result, tree[0].value)
		fmt.Println(result)
	}

	return arr

}

func main() {
	var length = 10
	var mymap = make(map[int]int, length)
	var obj []int
	//构造map
	for i := 0; i < length; i++ {
		//map哈希的随机存储
		mymap[i] = i
	}
	for k, _ := range mymap {
		obj = append(obj, k)
	}

	fmt.Println(obj)
	fmt.Println(TreeSelectSort(obj))

}
