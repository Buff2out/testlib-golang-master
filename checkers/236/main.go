package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

type Pair struct {
	first  int
	second int
}

func dist(a, b Pair) float64 {
	return math.Sqrt(math.Pow(float64(a.first-b.first), 2) + math.Pow(float64(a.second-b.second), 2))
}

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func createMatrix0(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func checkValueOutOfSegmen(val, a, b int) bool {
	if val < a || val > b {
		return true
	} else {
		return false
	}
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()

	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	dots := make([]Pair, 0)
	//var cnt float64 = 0.0
	n := core.ReadInt(inputStream)
	k := core.ReadInt(inputStream)
	var reference float64 = 0.0
	var result float64 = 0.0
	var cluster_sum float64 = 0.0
	cluster := 0
	var x float64 = 0.0
	var y float64 = 0.0
	num := 0
	var tmp float64 = 0.0
	for i := 0; i < n; i++ {
		temp := Pair{
			first:  core.ReadInt(inputStream),
			second: core.ReadInt(inputStream),
		}
		dots = append(dots, temp)
	}
	reference = core.ReadFloat(outputStream)
	result = core.ReadFloat(inputFromUserStream)
	if result > reference*1.3 {
		testlib.WriteAnswer(fmt.Sprintf("Not correct clasterisation"), testlib.WA)
		return
	}
	// double cluster_sum = 0; already initialized
	for i := 0; i < k; i++ {
		x = core.ReadFloat(inputFromUserStream)
		y = core.ReadFloat(inputFromUserStream)
		num = core.ReadInt(inputFromUserStream)
		tmp = 0.0
		for j := 0; j < num; j++ {
			cluster = core.ReadInt(inputFromUserStream) - 1
			tmp += math.Abs(math.Sqrt(math.Pow(x-float64(dots[cluster].first), 2) + math.Pow(y-float64(dots[cluster].second), 2)))
		}
		cluster_sum += tmp
	}
	if math.Abs(result-cluster_sum) < 10 {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("WTF. Result: %.3f, Clister sum: %.3f", result, cluster_sum), testlib.WA)
}
