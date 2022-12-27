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
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n := core.ReadInt(inputStream)
	start := core.ReadInt(inputStream)
	g := createMatrix0(n, n)
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = core.ReadInt(inputStream)
		}
		visited[i] = 0
	}
	x := core.ReadInt(inputFromUserStream)
	if x != start {
		testlib.WriteAnswer(fmt.Sprintf("WA x(%d) != start(%d)", x, start), testlib.WA)
		return
	}
	y := 0
	// тут был visited, запихал его инициализацию в предыдущий цикл
	for i := 0; i < n; i++ {
		y = core.ReadInt(inputFromUserStream)
		if !(i2b(g[x-1][y-1])) {
			testlib.WriteAnswer(fmt.Sprintf("WA two unconnected vertices %d %d", x, y), testlib.WA)
			return
		}
		visited[y-1]++
		x = y
	}
	if y != start {
		testlib.WriteAnswer(fmt.Sprintf("WA y(%d) != start(%d)", y, start), testlib.WA)
		return
	}
	for i := 0; i < n; i++ {
		if visited[i] != 1 {
			testlib.WriteAnswer(fmt.Sprintf("vertice visited not 1 time"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
