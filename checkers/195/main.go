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

var g = createMatrix0(100, 100)

func main() {

	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n := core.ReadInt(inputStream)
	m := 0
	a := createMatrix0(100, 100)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = core.ReadInt(inputStream)
			if i2b(a[i][j]) {
				m++
			}
		}
	}
	m /= 2
	p := core.ReadInt(inputFromUserStream)
	anss := core.ReadInt(outputStream)
	if anss == -1 && p != -1 {
		testlib.WriteAnswer(fmt.Sprintf("-1"), testlib.WA)
		return
	}
	if anss == -1 && p == -1 {
		testlib.WriteAnswer("OK", testlib.OK)
		return
	}
	p--
	f := p
	for i := 0; i < m; i++ {
		r := core.ReadInt(inputFromUserStream)
		r--
		if a[p][r] == 0 {
			testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
			return
		}
		if i2b(g[r][p]) {
			testlib.WriteAnswer(fmt.Sprintf("WA3"), testlib.WA)
			return
		}
		g[r][p] = 1
		g[p][r] = 1
		p = r
	}
	if f != p {
		testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
		return
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
