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
	n1, n2, n := core.ReadInt(inputFromUserStream), core.ReadInt(outputStream), core.ReadInt(inputStream)
	if n1 == n2 && (n1 == 0 || n1 == -1) {
		for {
			if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
				if err != nil {
					testlib.WriteAnswer(err.Error(), testlib.EF)
					return
				}
				break
			}
			testlib.WriteAnswer(fmt.Sprintf(" Too many input data "), testlib.PE)
			return
		}
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	//xor (n1 == -1) ^ (n2 == -1)
	if ((n1 == -1) || (n2 == -1)) && !((n1 == -1) && (n2 == -1)) {
		testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
		return
	}
	if n1 > n2 {
		testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
		return
	}
	a := createMatrix0(n, n)
	path := make([]int, n1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = core.ReadInt(inputStream)
		}
	}
	s := core.ReadInt(inputStream)
	t := core.ReadInt(inputStream)
	for i := 0; i < n1+1; i++ {
		if i == 0 {
			x := core.ReadInt(inputFromUserStream)
			if x != s {
				testlib.WriteAnswer(fmt.Sprintf("WA, x != s"), testlib.WA)
				return
			}
			x--
			path[i] = x
		} else if i == n1 {
			x := core.ReadInt(inputFromUserStream)
			if x != t {
				testlib.WriteAnswer(fmt.Sprintf("WA, x != t"), testlib.WA)
				return
			}
			x--
			path[i] = x
		} else {
			x := core.ReadInt(inputFromUserStream)
			x--
			path[i] = x
		}
	}
	for i := 0; i < len(path)-1; i++ {
		if !i2b(a[path[i]][path[i+1]]) {
			testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
}
