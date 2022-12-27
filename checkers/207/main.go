package main

import (
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
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

var m = createMatrix0(1000, 1000)

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	// inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	a := createMatrix0(100, 100)
	cnt := 0
	core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%d\n", 0))
	core.FlushWriter(outputToUserStream)
	for {
		if flag, err := utils.IsStreamFinish(outputStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		x := core.ReadInt(outputStream)
		y := core.ReadInt(outputStream)
		if x != y {
			a[x][y] = 1
			a[y][x] = 1
			cnt++
		}
	}
	var cnt1 float64 = 0
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		x := core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(x, 1, 100) {
			testlib.WriteAnswer(fmt.Sprintf("out of segment x [1,100]"), testlib.WA)
			return
		}
		y := core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(y, 1, 100) {
			testlib.WriteAnswer(fmt.Sprintf("out of segment y [1,100]"), testlib.WA)
			return
		}
		x--
		y--
		if a[x][y] == 1 {
			cnt1++
		} else {
			cnt1 -= 0.5
		}
	}
	if cnt1/float64(cnt) > 0.5 {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	} else {
		testlib.WriteAnswer(fmt.Sprintf("WA ", cnt1*100/float64(cnt)), testlib.WA)
		return
	}
}
