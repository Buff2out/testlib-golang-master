package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func f1(x, y float64) float64 {
	return x*x + y*y - 0.1*math.Abs(1-x) - 0.1*math.Abs(1-y)
}

func f2(x, y float64) float64 {
	return 20*math.Abs(x-50)*math.Abs(y-25) + 10*(math.Abs(x-10)*math.Abs(y-10)+math.Abs(x-50))
}

func dist(anX, anY, x, y float64) float64 {
	return math.Sqrt(math.Pow(anX-x, 2) + math.Pow(anY-y, 2))
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
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	// n := core.ReadInt(inputStream)
	var eps float64 = 0.001
	f := core.ReadInt(inputStream)
	k := 0
	// ans := make([]float64, 0)
	var x, y float64
	cnt := 0
	var requestType byte
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		k++
		if cnt > 100000 {
			testlib.WriteAnswer(fmt.Sprintf("Too many requests"), testlib.WA)
			return
		}
		requestType = core.ReadByte(inputFromUserStream)
		x = core.ReadFloat(inputFromUserStream)
		y = core.ReadFloat(inputFromUserStream)
		switch requestType {
		case byte('!'):
			if i2b(f) {
				if math.Abs(x+0.05) < eps && math.Abs(y+0.05) < eps {
					testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("Incorrect output: %.3f %.3f", x, y), testlib.WA)
					return
				}
			} else if math.Abs(50-x) < eps && math.Abs(10-y) < eps {
				testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
				return
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Incorrect output: %.3f %.3f", x, y), testlib.WA)
				return
			}
		case byte('?'):
			if i2b(f) {
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f\n", f1(x, y)))
				core.FlushWriter(outputToUserStream)
			} else {
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f\n", f2(x, y)))
				core.FlushWriter(outputToUserStream)
			}
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Incorrect input command: %s", string(requestType)), testlib.PE)
			return
		}
	}
	testlib.WriteAnswer(fmt.Sprintf("Incorrect: Input after EOF"), testlib.EF)
	return
}
