package main

import (
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func count(i byte) (int, bool) {
	switch i {
	case byte('L'):
		return 1, false
		break
	case byte('R'):
		return 3, false
		break
	case byte('U'):
		return 7, false
		break
	case byte('D'):
		return 15, false
		break
	default:
		return 0, true
	}
	return 0, true
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
	//inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	counterAnswer := 0
	counterUser := 0
	i1 := 0
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		x := core.ReadWord(outputStream)
		y := core.ReadWord(inputFromUserStream)
		for j := 0; j < len(x); j++ {
			temp, err := count(byte(x[j]))
			if err {
				testlib.WriteAnswer(fmt.Sprintf(" Incorrect input"), testlib.PE)
				return
			}
			counterAnswer += temp
		}
		for j := 0; j < len(x); j++ {
			temp, err := count(byte(y[j]))
			if err {
				testlib.WriteAnswer(fmt.Sprintf(" Incorrect input"), testlib.PE)
				return
			}
			counterAnswer += temp
		}
	}
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		testlib.WriteAnswer(fmt.Sprintf("Incorrect path.Your path is too long "), testlib.WA)
		return
	}
	for {
		if flag, err := utils.IsStreamFinish(outputStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		testlib.WriteAnswer(fmt.Sprintf("Incorrect path.Your path is too short "), testlib.WA)
		return
	}
	if counterAnswer != counterUser {
		testlib.WriteAnswer(fmt.Sprintf("Incorrect path.Path is not available "), testlib.WA)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf(" Ok. %d  tokens equal.", i1), testlib.OK)
}
