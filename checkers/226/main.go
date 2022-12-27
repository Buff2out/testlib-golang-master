package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

func func1(x float64) float64 {
	_, cosRes := math.Sincos(0.5 * x)
	return 2 * cosRes
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
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	n := core.ReadInt(inputStream)
	a := core.ReadFloat(inputStream)
	b := core.ReadFloat(inputStream)
	ans := make([]float64, 0)
	cnt := 0
	var requestType byte
	var r float64
	var x float64
	for i := 0; i < n; i++ {
		ans = append(ans, core.ReadFloat(inputStream))
	}
	core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f %.20f %.20f\n", n, a, b))
	core.FlushWriter(outputToUserStream)

	for {
		cnt++
		requestType = core.ReadByte(inputFromUserStream)
		switch requestType {
		case byte('!'):
			for i := 0; i < n; i++ {
				r = core.ReadFloat(inputFromUserStream)
				if math.Abs(r-ans[i]) > 0.25 {
					testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Abs: %.3f %.3f number: %d", x, ans[i], i), testlib.WA)
					return
				}
				testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
				return
			}
		case byte('?'):
			x = core.ReadFloat(inputFromUserStream)
			core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f\n", func1(x)))
			core.FlushWriter(outputToUserStream)
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Incorrect input command: %s", string(requestType)), testlib.PE)
			return
		}
	}
}
