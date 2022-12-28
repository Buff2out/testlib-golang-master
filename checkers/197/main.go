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

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	f := []int{-2, -2, -1, -1, 1, 1, 2, 2}
	s := []int{1, -1, 2, -2, -2, 2, -1, 1}
	n, m := core.ReadInt(inputStream), core.ReadInt(inputStream)
	c := createMatrix0(n+1, m+1)
	temp := core.ReadInt(inputFromUserStream)
	if temp == -1 {
		if n == 4 && m == 4 {
			testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
			return
		} else {
			testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
			return
		}
	}
	if temp < 1 || temp > n {
		testlib.WriteAnswer(fmt.Sprintf("PE Index out of range"), testlib.PE)
		return
	}
	x, y := temp, core.ReadInt(inputFromUserStream)
	if y < 1 || y > m {
		testlib.WriteAnswer(fmt.Sprintf("PE Index out of range"), testlib.PE)
		return
	}
	c[x][y] = 1
	x2, y2 := x, y
	for i := 0; i < n*m-1; i++ {
		isFind := false
		x1, y1 := core.ReadInt(inputFromUserStream), core.ReadInt(inputFromUserStream)
		if y1 < 1 || y1 > m || x1 < 1 || x1 > n {
			testlib.WriteAnswer(fmt.Sprintf("PE Index out of range"), testlib.PE)
			return
		}
		for j := 0; j < 8; j++ {
			if x1-x == f[j] && y1-y == s[j] {
				isFind = true
				break
			}
		}
		if !isFind {
			testlib.WriteAnswer(fmt.Sprintf("WA Incorrect step : x = %d, y = %d", x1, y1), testlib.WA)
			return
		}
		if c[x1][y1] == 1 && !(x1 == x2 && y1 == y2) {
			testlib.WriteAnswer(fmt.Sprintf("WA Twice"), testlib.WA)
			return
		}
		c[x1][y1] = 1
		x = x1
		y = y1
	}
	testlib.WriteAnswer(fmt.Sprintf("Full solution"), testlib.OK)
}
