package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

func func1(a, b, x float64) float64 {
	temp := x + a
	return temp*temp + b
}

func func2(a, b, c, x float64) float64 {
	sinRes, _ := math.Sincos(x)
	return sinRes + math.Abs(x+a) - math.Abs(x+b) + math.Abs(x+c)
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

	test := core.ReadInt(inputStream)
	cnt := 0
	var a float64 = 0.0
	var b float64 = 0.0
	var c float64 = 0.0
	var an float64 = 0.0
	var x float64 = 0.0
	var requestType byte = 0.0
	if test > 10 {
		a = core.ReadFloat(inputStream)
		b = core.ReadFloat(inputStream)
		c = core.ReadFloat(inputStream)

		if math.Abs(a-b) > math.Abs(b-c) {
			an = -a
		} else {
			an = -c
		}
		for {
			requestType = core.ReadByte(inputFromUserStream)
			if cnt > 1000000000 {
				testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
				return
			}
			switch requestType {
			case byte('!'):
				x = core.ReadFloat(inputFromUserStream)
				if (math.Abs(x-an) > 0.25) && ((math.Abs(func2(a, b, c, x)) - func2(a, b, c, an)) > 1.9) {
					testlib.WriteAnswer(fmt.Sprintf("WA x: %.3f, ans: %.3f, abs1: %.3f, abs2: %.3f", x, an, math.Abs(x-an), math.Abs(func2(a, b, c, an)-func2(a, b, c, x))), testlib.WA)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
					return
				}
				break
			case byte('?'):
				x = core.ReadFloat(inputFromUserStream)
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f\n", func2(a, b, c, x)))
				core.FlushWriter(outputToUserStream)
				break
			default:
				testlib.WriteAnswer(fmt.Sprintf("Wrong query: %d", string(requestType)), testlib.PE)
				return
			}
			cnt++
		}
	} else {
		a = core.ReadFloat(inputStream)
		b = core.ReadFloat(inputStream)
		an = -a
		for {
			if cnt > 1000000000 {
				testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
				return
			}
			switch requestType {
			case byte('!'):
				x = core.ReadFloat(inputFromUserStream)
				if math.Abs(x-an) > 0.25 {
					testlib.WriteAnswer(fmt.Sprintf("WA x: %.3f, ans: %.3f", x, an, math.Abs(x-an)), testlib.WA)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
					return
				}
				break
			case byte('?'):
				x = core.ReadFloat(inputFromUserStream)
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f\n", func1(a, b, x)))
				core.FlushWriter(outputToUserStream)
				break
			default:
				testlib.WriteAnswer(fmt.Sprintf("Wrong query: %d", string(requestType)), testlib.PE)
				return
			}
			cnt++
		}
	}
}
