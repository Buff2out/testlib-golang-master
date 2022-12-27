package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

// type Pair struct {
// 	first  int
// 	second int
// }

// func dist(a, b Pair) float64 {
// 	return math.Sqrt(math.Pow(float64(a.first-b.first), 2) + math.Pow(float64(a.second-b.second), 2))
// }

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
	var cnt float64 = 0.0
	var x int
	var xStr, yStr string
	for i := 0; i < 3; i++ {
		x = core.ReadInt(inputStream)
	}
	for i := 0; i < x; i++ {
		xStr = core.ReadWord(outputStream)
		yStr = core.ReadWord(inputFromUserStream)
		if xStr == yStr {
			cnt++
		}
	}
	if cnt/float64(x) > 0.8 {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	} else {
		testlib.WriteAnswer(fmt.Sprintf("WA %.3f", cnt/50), testlib.WA)
		return
	}
}
