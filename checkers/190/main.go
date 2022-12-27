package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

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

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n := core.ReadInt(inputStream)
	g := createMatrix0(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = core.ReadInt(inputStream)
		}
	}
	res := core.ReadWord(inputFromUserStream)
	ans := core.ReadWord(outputStream)
	if res != ans {
		testlib.WriteAnswer(fmt.Sprintf("NOPE"), testlib.WA)
		return
	}
	if res == "NO" {
		testlib.WriteAnswer("OK", testlib.OK)
		return
	}
	cnt1 := core.ReadInt(inputFromUserStream)
	if checkValueOutOfSegmen(cnt1, 2, 2*n) {
		testlib.WriteAnswer(fmt.Sprintf("Value out of segment %d and %d", 2, 2*n), testlib.WA)
		return
	}
	x := core.ReadInt(inputFromUserStream)
	if checkValueOutOfSegmen(x, 1, n) {
		testlib.WriteAnswer(fmt.Sprintf("Value out of segment %d and %d", 1, n), testlib.WA)
		return
	}
	cnt := 0
	y := 0
	for i := 0; i < cnt1-1; i++ {
		y = core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(y, 1, n) {
			testlib.WriteAnswer(fmt.Sprintf("Value out of segment %d and %d", 1, n), testlib.WA)
			return
		}
		if g[x-1][y-1] == 100000 {
			testlib.WriteAnswer(fmt.Sprintf("WA two unconnected vertices  %d %d", x, y), testlib.WA)
			return
		}
		cnt += g[x-1][y-1]
		x = y
	}
	if cnt >= 0 {
		testlib.WriteAnswer(fmt.Sprintf("ABOVE ZERO"), testlib.WA)
		return
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
