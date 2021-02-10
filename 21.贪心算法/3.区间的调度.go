package main

import (
	"fmt"
	"sort"
)

const N = 5

var Start = []int{1, 2, 4, 6, 8}
var End = []int{3, 5, 7, 9, 10}

type pair struct {
	start int
	end   int
}

type pairs [N]pair

var mypairs pairs

func (my pairs) Len() int {
	return len(mypairs)
}

func (my pairs) Swap(i, j int) {
	mypairs[i], mypairs[j] = mypairs[j], mypairs[i]
}

func (my pairs) Less(i, j int) bool {
	return mypairs[i].end < mypairs[j].end
}

func Select() int {
	for i := 0; i < N; i++ {
		//保存
		mypairs[i].start = Start[i]
		mypairs[i].end = End[i]
	}

	sort.Sort(mypairs)

	//排序
	count := 0
	t := 0
	for i := 0; i < N; i++ {
		if t < mypairs[i].start {
			fmt.Println(i)
			count++
			t = mypairs[i].end
		}
	}
	return count
}

func main() {
	fmt.Println(Select())
}
