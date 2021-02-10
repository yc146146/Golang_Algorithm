package main


func getN(N int)int{
	if N == 1{
		return 1
	}else if N == 2{
		return 2
	}else{
		//链表 红黑树 线索树
		return getN(N-1)+getN(N-2)
	}
}



func main() {

}
