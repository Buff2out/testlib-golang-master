package main

import (
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func createMatrix(n int, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		// for j := 0; j < m; j++ {
		// 	matrix[i][j] = 0
		// }
	}
	return matrix
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	gMtrx := createMatrix(100, 100)
	inMtrx := createMatrix(100, 100)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			inMtrx[i][j] = core.ReadInt(inputStream)
		}
	}
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		a := core.ReadInt(inputFromUserStream)
		if 1 > a || a > n {
			testlib.WriteAnswer(fmt.Sprintf("Failed to read data from inputFromUserStream, var a < 1 or n < var a"), testlib.WA)
			return
		}
		b := core.ReadInt(inputFromUserStream)
		if 1 > b || b > n {
			testlib.WriteAnswer(fmt.Sprintf("Failed to read data from inputFromUserStream, var b < 1 or n < var b"), testlib.WA)
			return
		}
		a--
		b--
		gMtrx[a][b] = 1
		gMtrx[a][b] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if gMtrx[i][j] != inMtrx[i][j] {
				// msg << "WA wrong cell " << i << " " << j;
				testlib.WriteAnswer(fmt.Sprintf("WA wrong cell %d %d", i, j), testlib.WA)
				return
			}
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
