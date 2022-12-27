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

	a := createMatrix0(10000, 10000)
	n := core.ReadInt(inputStream)
	m := core.ReadInt(inputStream)
	px1 := core.ReadInt(inputStream)
	py1 := core.ReadInt(inputStream)
	px2 := core.ReadInt(inputStream)
	py2 := core.ReadInt(inputStream)
	var dx, dy int
	var ln int
	var opCount int
	eps := 3
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = core.ReadInt(inputStream)
		}
	}
	dx = core.ReadInt(inputFromUserStream)
	dy = core.ReadInt(inputFromUserStream)
	if dx != px1 || dy != py1 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong start position"), testlib.WA)
		return
	}
	for i := 0; i < ln-1; i++ {
		dx = core.ReadInt(inputFromUserStream)
		dy = core.ReadInt(inputFromUserStream)
		if (dx <= m && dx >= 0 && dy <= n && dy >= 0) && a[dy-1][dx-1] == 0 && ((dy == py1+1 && dx == px1) || (dy == py1-1 && dx == px1) || (dy == py1 && dx == px1+1) || (dy == py1 && dx == px1-1)) {
			px1 = dx
			py1 = dy
		} else {
			testlib.WriteAnswer(fmt.Sprintf("Impossible movement %d, %d, %d, %d", px1, py1, dx, dy), testlib.WA)
			return
		}
	}
	if px1 != px2 || py1 != py2 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong end position"), testlib.WA)
		return
	}
	opCount = core.ReadInt(outputStream)
	//eps already initialized
	if ln-opCount > eps {
		testlib.WriteAnswer(fmt.Sprintf("Too long way"), testlib.WA)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
}
