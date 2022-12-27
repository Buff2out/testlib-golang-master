package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

func func1(a, b int, x, y float64) float64 {
	temp1, temp2 := x+float64(a), y+float64(b)
	return temp1*temp1 + temp2*temp2
}

func dx(a int, x float64) float64 {
	return 2 * (x + float64(a))
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
	a := core.ReadFloat(inputStream)
	b := core.ReadFloat(inputStream)
	// ans := make([]float64, 0)
	var anX float64 = float64(-a)
	var anY float64 = float64(-b)
	var x, y float64
	cnt := 0
	var requestType byte
	for {
		requestType = core.ReadByte(inputFromUserStream)
		if cnt > 1000000000 {
			testlib.WriteAnswer(fmt.Sprintf("Too many requests"), testlib.WA)
			return
		}
		switch requestType {
		case byte('!'):
			for i := 0; i < n; i++ {
				x = core.ReadFloat(inputFromUserStream)
				y = core.ReadFloat(inputFromUserStream)
				if dist(anX, anY, x, y) > 0.1 {
					testlib.WriteAnswer(fmt.Sprintf("Incorrect user output. User answer: %.3f, %.3f. abs2: %.3f", x, y], dist(anX, anY, x, y)), testlib.WA)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
					return
				}
			}
		case byte('?'):
			x = core.ReadFloat(inputFromUserStream)
			y = core.ReadFloat(inputFromUserStream)
			core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%.20f %.20f %.20f\n", func1(a,b,x,y), dx(a,x), dx(b,y)))
			core.FlushWriter(outputToUserStream)
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Wrong query: %s", string(requestType)), testlib.PE)
			return
		}
	}
}
