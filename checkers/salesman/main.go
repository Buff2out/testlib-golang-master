package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

type Pair struct {
	first  int
	second int
}

func dist(a, b Pair) float64 {
	return math.Sqrt(math.Pow(float64(a.first-b.first), 2) + math.Pow(float64(a.second-b.second), 2))
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
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	n := core.ReadInt(inputStream)
	answer := core.ReadFloat(outputStream)
	userAnswer := core.ReadFloat(inputFromUserStream)
	vec := make([]Pair, 0)
	for i := 0; i < n; i++ {
		temp := Pair{
			first:  core.ReadInt(inputStream),
			second: core.ReadInt(inputStream),
		}
		vec = append(vec, temp)
	}
	ind := make([]int, n)
	u := make([]int, n)
	for i := 0; i < n; i++ {
		u[i] = 0
	}
	for i := 0; i < n; i++ {
		ind[i] = core.ReadInt(inputFromUserStream)
		if i2b(u[ind[i]-1]) {
			testlib.WriteAnswer(fmt.Sprintf("Value already exist"), testlib.PE)
			return
		}
		u[ind[i]-1]++
	}
	var res float64 = 0.0
	for i := 0; i < n; i++ {
		res += dist(vec[ind[i]-1], vec[ind[(i+1)%n]-1])
	}
	if math.Abs(res-userAnswer) > 0.5 {
		testlib.WriteAnswer(fmt.Sprintf("Incorrect output, User answer: %.3f. Correct output: %.3f", userAnswer, res), testlib.WA)
		return
	}
	if userAnswer-answer > answer*0.1 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Absolute diff: %.3f", userAnswer), testlib.WA)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
}
