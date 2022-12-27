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
type PairFloat64 struct {
	first  float64
	second float64
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
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	x1, y1 := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	x2, y2 := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	sumx := x1 + x2
	sumy := y1 + y2
	ans11 := PairFloat64{
		first:  sumx/2 + (y1-y2)/2,
		second: sumy/2 + (x2-x1)/2,
	}
	ans12 := PairFloat64{
		first:  sumx/2 + (y2-y1)/2,
		second: sumy/2 + (x1-x2)/2,
	}
	ans1 := PairFloat64{
		first:  core.ReadFloat(inputFromUserStream),
		second: core.ReadFloat(inputFromUserStream),
	}
	ans2 := PairFloat64{
		first:  core.ReadFloat(inputFromUserStream),
		second: core.ReadFloat(inputFromUserStream),
	}
	if (ans1 == ans11 && ans2 == ans12) || (ans1 == ans12 && ans2 == ans11) {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("WA ans1: %d, %d; ans2: %d, %d; ans11: %d, %d; ans12: %d, %d;", ans1.first, ans1.second, ans2.first, ans2.second, ans11.first, ans11.second, ans12.first, ans12.second), testlib.WA)
}
