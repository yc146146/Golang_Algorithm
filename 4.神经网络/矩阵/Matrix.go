package main


import (
	"errors"
	"fmt"
)

//
type Matrix struct {
	rows int

	cols int

	//i*strp+j
	Elements []float64
	//行的偏移量
	step int
}

func MakeMatrix(Elements []float64, rows, cols int) *Matrix {
	A := new(Matrix)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.Elements = Elements
	return A
}

//返回多少行
func (A *Matrix) CountRows() int {
	return A.rows
}

func (A *Matrix) CountCols() int {
	return A.cols
}

func (A *Matrix) GetElem(i, j int) float64 {
	return A.Elements[i*A.step+j]
}

//设置数据
func (A *Matrix) SetElem(i, j int, v float64) {
	if i*A.step + j > len(A.Elements)-1{
		return
	}
	A.Elements[i*A.step+j] = v
}

//深度拷贝 对角线
func (A *Matrix) diagonalCopy() []float64 {
	diag := make([]float64, A.cols)
	for i := 0; i < len(diag); i++ {
		diag[i] = A.GetElem(i, i)
	}
	return diag
}

//拷贝
func (A *Matrix) copy() *Matrix {
	B := new(Matrix)
	B.rows = A.rows
	B.cols = A.cols
	B.step = A.step
	B.Elements = make([]float64, A.cols*A.rows)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Elements[i*A.step+j] = A.GetElem(i,j)
		}
	}
	return B
}

func (A *Matrix) trace() float64 {
	var tr float64 = 0
	for i := 0; i < A.cols; i++ {
		//叠加对角线
		tr += A.GetElem(i, i)
	}
	return tr
}

//矩阵加法
func (A *Matrix) add(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("矩阵大小不一样不能加")
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElem(i, j, A.GetElem(i, j)+B.GetElem(i, j))
		}
	}

	return nil
}

//矩阵减法 A+=B
func (A *Matrix) sub(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("矩阵大小不一样不能加")
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElem(i, j, A.GetElem(i, j)-B.GetElem(i, j))
		}
	}

	return nil
}

//矩阵乘法 与整数
func (A *Matrix) scale(a float64) error {

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElem(i, j, A.GetElem(i, j)*a)
		}
	}

	return nil
}

func Add(A *Matrix, B *Matrix) *Matrix {
	if A.cols != B.cols && A.rows != B.rows {
		return nil
	}

	res := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			res.SetElem(i, j, A.GetElem(i, j)+B.GetElem(i, j))
		}
	}
	return res
}

func Sub(A *Matrix, B *Matrix) *Matrix {
	if A.cols != B.cols && A.rows != B.rows {
		return nil
	}
	res := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			res.SetElem(i, j, A.GetElem(i, j)-B.GetElem(i, j))
		}
	}
	return res

}

func Mutiply(A *Matrix, B *Matrix) *Matrix {
	res := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.GetElem(i, k) * B.GetElem(k, j)
			}
			res.SetElem(i, j, sum)
		}
	}
	return res

}

func main() {
	a := []float64{1,2,3,4,5,6}
	A:=MakeMatrix(a,3,2)
	fmt.Println(A.cols,A.rows)

	b := []float64{11,22,33,34,35,36}
	B:=MakeMatrix(b,3,2)
	fmt.Println(B.cols,B.rows)
	fmt.Println(A.GetElem(0,0))
	fmt.Println(A.GetElem(1,0))
	fmt.Println(A.GetElem(2,0))
	C := Add(A,B)
	fmt.Println(C.rows, C.cols)
	D := Mutiply(A,A)
	fmt.Println(D.cols, D.rows)
}
