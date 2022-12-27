package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func dfs(a [][]int, u []int, s int, n int) {
	for i := 0; i < n; i++ {
		if i2b(a[s][i]) && !i2b(u[i]) {
			u[i]++
			dfs(a, u, i, n)
		}
	}
}

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
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n1, n2 := core.ReadInt(outputStream), core.ReadInt(inputFromUserStream)
	if n1 > n2 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer. User output greater than correct"), testlib.WA)
		return
	}
	n, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	aMatrix := createMatrix(n, n)
	s := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			aMatrix[i][j] = core.ReadInt(inputFromUserStream)
			s += aMatrix[i][j]
		}
	}
	if n2 != s/2 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer"), testlib.WA)
		return
	}
	u := make([]int, n)
	for i := 0; i < n; i++ {
		u = append(u, 0)
	}
	u[0]++
	dfs(aMatrix, u, 0, n)
	count := 0
	for i := 0; i < n; i++ {
		if i2b(u[i]) {
			count++
		}
	}
	if count != n {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer"), testlib.WA)
		return
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
