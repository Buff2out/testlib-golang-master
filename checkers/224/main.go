package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

func f1(x, y float64) float64 {
	return x*x + y*y - 0.1*math.Abs(1-x) - 0.1*math.Abs(1-y)
}

func fun(a, b, c, d, x float64) float64 {
	return x*x*x*a + x*x*b + x*c + d
}

func diff(a, b, c, d, x float64) float64 {
	return 3*x*x*a + 2*x*b + c
}

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
		_ = testlib.GetStream(testlib.OutputToUser).Close()
		_ = testlib.GetStream(testlib.TaskInput).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	a, b := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	c, d := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	an := core.ReadFloat(inputStream)
	var x float64
	an *= -1
	cnt := 0
	var requestType byte
	for {
		if cnt >= 50 {
			testlib.WriteAnswer(fmt.Sprintf("Too many requests"), testlib.WA)
			return
		}
		requestType = core.ReadByte(inputFromUserStream)
		switch requestType {
		case byte('!'):
			x = core.ReadFloat(inputFromUserStream)
			if math.Min(math.Min(math.Abs(x-an), math.Abs(fun(a, b, c, d, x)-fun(a, b, c, d, an))), math.Abs(diff(a, b, c, d, x)-diff(a, b, c, d, an))) > 1e-6 {
				testlib.WriteAnswer(fmt.Sprintf("Wrong answer. x: %.3f ans: %.3f abs: %.f", x, an, math.Abs(x-an)), testlib.WA)
				return
			} else {
				testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
				return
			}
			break
		case byte('?'):
			x = core.ReadFloat(inputFromUserStream)
			core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%20.f %20.f\n", fun(a, b, c, d, x), diff(a, b, c, d, x)))
			core.FlushWriter(outputToUserStream)
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Incorrect input command: %s", string(requestType)), testlib.PE)
			return
		}
		cnt++
	}
	testlib.WriteAnswer(fmt.Sprintf("Incorrect: Input after EOF"), testlib.EF)
	return
}
