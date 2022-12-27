package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func createMatrix0(n int, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func checkValueOutOfSegmen(val int, a int, b int) bool {
	if val < a || val > b {
		return true
	} else {
		return false
	}
}

var w = createMatrix0(1000, 1000)
var m = createMatrix0(1000, 1000)

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n := core.ReadInt(inputStream)
	so := core.ReadInt(inputFromUserStream)
	sa := core.ReadInt(outputStream)
	// g := createMatrix0(n, n)
	uw := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			w[i][j] = core.ReadInt(inputStream)
		}
		uw[i] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[i][j] = core.ReadInt(inputStream)
		}
	}
	s := 0
	k := 0
	for i := 0; i < n; i++ {
		k = core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(k, 1, n) {
			testlib.WriteAnswer(fmt.Sprintf("out of segment [1,n]"), testlib.WA)
			return
		}
		k--
		if i2b(uw[k]) {
			testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
			return
		}
		uw[k]++
		s += w[i][k]
		s += m[k][i]
	}
	if s != so {
		testlib.WriteAnswer(fmt.Sprintf("WA3"), testlib.WA)
		return
	}
	if s != sa {
		testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
		return
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
