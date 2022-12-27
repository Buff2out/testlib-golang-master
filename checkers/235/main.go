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
	var counter int = 0
	a1 := 0
	b1 := 0
	a2 := 0
	b2 := 0
	// var x int
	// var xStr, yStr string
	for i := 0; i < 5; i++ {
		a1 = core.ReadInt(inputStream)
		b1 = core.ReadInt(inputStream)
		a2 = core.ReadInt(inputStream)
		b2 = core.ReadInt(inputStream)
		if a1 == a2 && b1 == b2 {
			counter++
		}
	}
	if counter < 4 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Only %d of 5 was right", counter), testlib.WA)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
}
