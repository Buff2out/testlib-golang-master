package main

import (
	"fmt"
	"math"
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
		//_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
		//_ = testlib.GetStream(testlib.OutputToUser).Close()
		_ = testlib.GetStream(testlib.TaskInput).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	a, b := core.ReadInt(inputStream), core.ReadInt(inputStream)
	c, d := core.ReadInt(inputStream), core.ReadInt(inputStream)
	var x float64 = core.ReadFloat(inputFromUserStream)
	if math.Abs(float64(a)*x*x*x+float64(b)*x*x+float64(c)*x+float64(d)) < 0.000001 {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("Incorrect user output"), testlib.WA)
	return
}
