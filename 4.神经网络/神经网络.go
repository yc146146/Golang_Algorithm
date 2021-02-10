package main

import (
	"fmt"
	"math"
	"math/rand"
)

type DG struct {
	//编号
	data float64
	//数据类型
	classname string
}

//生成随机数
func random(a, b float64) float64 {
	return (b-a)*rand.Float64() + a
}

func matrix(I, J int) [][]float64 {
	m := make([][]float64, I)
	for i := 0; i < I; i++ {
		m[i] = make([]float64, J)
	}
	return m
}

//数学计算
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

//数学计算反向
func disigmoid(y float64) float64 {
	return y * (1 - y)
}

func Vector(I int, fill float64) []float64 {
	v := make([]float64, I)
	for i := 0; i < I; i++ {
		v[i] = fill
	}
	return v
}

//神经网络结构体
type NeuralNetwork struct {
	//输入层 中间层 输出层 节点的数量
	NIputs, NHiddens, NOutputs int
	//是否线性回归
	Regression bool
	//	输入数据 输出数据 中间数据
	InputActivations, HiddenActivations, OutputActivations []float64
	//神经网络
	Contexts [][]float64
	//输出权重
	InputWeights, OutputWeights [][]float64
	//输入输出的改变
	InputChanges, OutputChanges [][]float64

	result []DG
}

func (nn *NeuralNetwork) Init(inputs, hiddens, outputs int, result []DG) {

	nn.result=result
	//跳过第一个
	nn.NIputs = inputs + 1
	nn.NHiddens = hiddens + 1
	//输出
	nn.NOutputs = outputs

	nn.InputActivations = Vector(nn.NIputs, 1.0)
	nn.HiddenActivations = Vector(nn.NHiddens, 1.0)
	nn.OutputActivations = Vector(nn.NOutputs, 1.0)

	nn.InputWeights = matrix(nn.NIputs, nn.NHiddens)
	nn.OutputWeights = matrix(nn.NHiddens, nn.NOutputs)

	//填充矩阵 按照随机数
	for i := 0; i < nn.NIputs; i++ {
		for j := 0; j < nn.NHiddens; j++ {
			nn.InputWeights[i][j] = random(-1, 1)
		}
	}

	for i := 0; i < nn.NHiddens; i++ {
		for j := 0; j < nn.NOutputs; j++ {
			nn.OutputWeights[i][j] = random(-1, 1)
		}
	}

	nn.InputChanges = matrix(nn.NIputs, nn.NHiddens)
	nn.OutputChanges = matrix(nn.NHiddens, nn.NOutputs)

}

//神经网络设置上下文
func (nn *NeuralNetwork) SetContext(nContexts int, initValues [][]float64) {
	if initValues == nil {
		initValues = make([][]float64, nContexts)
	}
	for i := 0; i < nContexts; i++ {
		//填充中间数
		initValues[i] = Vector(nn.NHiddens, 0.5)
	}
	nn.Contexts = initValues
}

//反向传播神经网络
func (nn *NeuralNetwork) BackPropagate(targets []float64, lRate, mFactor float64) float64 {
	if len(targets) != nn.NOutputs {
		panic("error, 目标数据有误")
	}
	//填充数据
	outputDeltas := Vector(nn.NOutputs, 0.0)
	for i := 0; i < nn.NOutputs; i++ {
		//计算损失值
		outputDeltas[i] = disigmoid(nn.OutputActivations[i]) * (targets[i] - nn.OutputActivations[i])
	}
	hiddenDelas := Vector(nn.NHiddens, 0.0)
	for i := 0; i < nn.NHiddens; i++ {
		var e float64
		for j := 0; j < nn.NOutputs; j++ {
			//计算收益
			e += outputDeltas[j] * nn.OutputWeights[i][j]
		}
		//反向推理
		hiddenDelas[i] = disigmoid(nn.HiddenActivations[i]) * e
	}

	for i := 0; i < nn.NHiddens; i++ {
		for j := 0; j < nn.NOutputs; j++ {
			change := outputDeltas[j] * nn.HiddenActivations[i]
			//计算权重
			nn.OutputWeights[i][j] = nn.OutputWeights[i][j] + lRate*change + mFactor*nn.OutputChanges[i][j]
			//保存修改的值
			nn.OutputChanges[i][j] = change
		}
	}

	for i := 0; i < nn.NIputs; i++ {
		for j := 0; j < nn.NHiddens; j++ {
			change := hiddenDelas[j] * nn.InputActivations[i]
			//计算权重
			nn.InputChanges[i][j] = nn.InputWeights[i][j] + lRate*change + mFactor*nn.InputChanges[i][j]
			//保存修改的值
			nn.InputChanges[i][j] = change
		}
	}
	//计算距离
	var e float64
	for i := 0; i < len(targets); i++ {
		e += 0.5 * math.Pow(targets[i]-nn.OutputActivations[i], 2)
	}

	return e

}

//训练神经网络
func (nn *NeuralNetwork) Train(partens [][][]float64, iterations int, lRate, mFactor float64, debug bool) []float64 {
	//错误集合
	errors := make([]float64, iterations)
	for i := 0; i < iterations; i++ {
		var e float64
		for _, p := range partens {
			//更新数据
			nn.Update(p[0])
			//反向传播
			tmp := nn.BackPropagate(p[1], lRate, mFactor)
			//叠加
			e += tmp
		}
		errors[i] = e
		if debug && i%1000 == 0 {
			//每1000次 显示进度
			fmt.Println(i, e)
		}
	}
	return errors
}

//更新数据 预测
func (nn *NeuralNetwork) Update(inputs []float64) []float64 {
	if len(inputs) != nn.NIputs-1 {
		panic("错误的输入")
	}

	for i := 0; i < nn.NIputs-1; i++ {
		//输入数据
		nn.InputActivations[i] = inputs[i]
	}

	for i := 0; i < nn.NHiddens-1; i++ {
		var sum float64
		for j := 0; j < nn.NIputs; j++ {
			sum += nn.InputActivations[j] * nn.InputWeights[j][i]
		}

		//计算出 上下文
		for k := 0; k < len(nn.Contexts); k++ {
			for j := 0; j < nn.NHiddens-1; j++ {
				//计算sum
				sum += nn.Contexts[k][j]
			}
		}
		//增加权重系数
		nn.HiddenActivations[i] = sigmoid(sum)
	}

	if len(nn.Contexts) > 0 {
		for i := len(nn.Contexts) - 1; i > 0; i-- {
			//从后往前叠加
			nn.Contexts[i] = nn.Contexts[i-1]
		}
		nn.Contexts[0] = nn.HiddenActivations
	}
	for i := 0; i < nn.NOutputs; i++ {
		var sum float64
		for j := 0; j < nn.NHiddens; j++ {
			sum += nn.HiddenActivations[j] * nn.OutputWeights[j][i]
		}
		nn.OutputActivations[i] = sigmoid(sum)
	}

	return nn.OutputActivations

}

//评估神经网络,测试
func (nn *NeuralNetwork) Test(partens [][][]float64) {
	for _, p := range partens {
		fmt.Println(p[0], "->", nn.Update(p[0]), ":", p[1])
	}
}

func (nn *NeuralNetwork)GetClassname(pat []float64)string{
	predict := nn.Update(pat)
	//	0.02

	miniabs := float64(math.MaxInt32)
	mini := -1
	for i:=0;i<len(nn.result);i++{
		nowabs := math.Abs( nn.result[i].data-predict[0])
		if miniabs > nowabs{
			miniabs = nowabs
			mini = i
		}
	}
	return nn.result[mini].classname


	return ""

}

func main() {

	rand.Seed(0)
	dgs := []DG{{0.,"屌丝"},{1.0,"高富"}}
	//1代表高富帅 0屌丝
	pattern := [][][]float64{
		{{160, 88}, {0}},
		{{182, 191}, {1}},
		{{179, 182}, {1}},
		{{159, 81}, {0}},
		//{{162, 82}, {0}},
		//{{159, 77}, {0}},
	}

	nn := &NeuralNetwork{}
	//初始化
	nn.Init(2, 2, 1,dgs)
	//训练
	nn.Train(pattern, 4000, 0.6, 0.4, false)
	nn.Test(pattern)

	inputs := []float64{163, 88}
	fmt.Println(nn.Update(inputs))

	inputs = []float64{180, 182}
	fmt.Println(nn.Update(inputs))

	fmt.Println(nn.GetClassname([]float64{180, 182}))
	fmt.Println(nn.GetClassname([]float64{159, 83}))
}
