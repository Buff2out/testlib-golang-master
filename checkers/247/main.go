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
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	x1, y1 := core.ReadInt(inputStream), core.ReadInt(inputStream)
	x2, y2 := core.ReadInt(inputStream), core.ReadInt(inputStream)
	diffy := y2 - y1
	diffx := x2 - x1
	ans11 := Pair{
		first:  x1 + diffy,
		second: y1 - diffx,
	}
	ans21 := Pair{
		first:  x1 - diffy,
		second: y1 + diffx,
	}
	ans12 := Pair{
		first:  x2 + diffy,
		second: y2 - diffx,
	}
	ans22 := Pair{
		first:  x2 - diffy,
		second: y2 + diffx,
	}
	ans1 := Pair{
		first:  core.ReadInt(inputFromUserStream),
		second: core.ReadInt(inputFromUserStream),
	}
	ans2 := Pair{
		first:  core.ReadInt(inputFromUserStream),
		second: core.ReadInt(inputFromUserStream),
	}
	if ((ans1 == ans11 && ans2 == ans12) || (ans1 == ans12 && ans2 == ans11)) || ((ans1 == ans22 && ans2 == ans21) || (ans2 == ans21 && ans1 == ans22)) {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("WA ans1: %d, %d; ans2: %d, %d; ans11: %d, %d; ans12: %d, %d; ans21: %d, %d; ans22: %d, %d", ans1.first, ans1.second, ans2.first, ans2.second, ans11.first, ans11.second, ans12.first, ans12.second, ans21.first, ans21.second, ans22.first, ans22.second), testlib.WA)
}
