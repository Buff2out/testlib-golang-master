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

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	a := make([]bool, 0)
	b := make([]string, 0)
	for {
		if flag, err := utils.IsStreamFinish(inputStream, inputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		b = append(b, core.ReadLine(inputStream, true))
	}
	cnt1 := 0
	x := false
	for {
		if flag, err := utils.IsStreamFinish(outputStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		x = i2b(core.ReadInt(outputStream))
		if x {
			cnt1++
		}
		a = append(a, x)
	}
	cnt := 0
	xint := 0
	y := 0
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		xint = core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(xint, 1, len(a)) {
			testlib.WriteAnswer(fmt.Sprintf("out of segment xint(%d) [1,len(a)(%d)]", xint, len(a)), testlib.WA)
			return
		}
		y = core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(y, 1, len(b[xint-1])-13) {
			testlib.WriteAnswer(fmt.Sprintf("out of segment y(%d) [1,len(a)(%d)]", y, len(b[xint-1])-13), testlib.WA)
			return
		}
		if !a[xint-1] {
			testlib.WriteAnswer(fmt.Sprintf("Merlin was silent"), testlib.WA)
			return
		}
		if b[xint-1][y-1:14] != "Avada-ke-davra" {
			testlib.WriteAnswer(fmt.Sprintf("Merlin didn't say that"), testlib.WA)
			return
		}
		cnt++
	}
	if cnt == cnt1 {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
	} else {
		testlib.WriteAnswer(fmt.Sprintf("Merlin said more"), testlib.WA)
	}
}
