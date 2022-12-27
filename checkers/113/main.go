package main

import (
	"fmt"
	"testlib"
	"testlib/core"
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
	t := createMatrix(3, n)
	for j := 0; j < n; j++ {
		t[0] = append(t[0], n-j+2)
	}
	a, err := core.ReadIntWithError(inputFromUserStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
		return
	} else if a < 1 && a > n {
		testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
		return
	}
	b, err := core.ReadIntWithError(inputFromUserStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
		return
	} else if b < 1 && b > 3 {
		testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
		return
	}
	c, err := core.ReadIntWithError(inputFromUserStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
		return
	} else if c < 1 && c > 3 {
		testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
		return
	}
	if t[b-1][len(t[b-1])-1] != a || !(len(t[c-1]) == 0) {
		if t[c-1][len(t[c-1])-1] < a {
			testlib.WriteAnswer(fmt.Sprintf("Wrong"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
