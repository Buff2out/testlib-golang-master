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
	var correctAnswer int = core.ReadInt(outputStream)
	var userLength int = core.ReadInt(inputFromUserStream)
	if correctAnswer < userLength {
		testlib.WriteAnswer(fmt.Sprintf("Merlin said more"), testlib.WA)
		return
	}
	if correctAnswer == 0 && userLength == -1 {
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
		testlib.WriteAnswer(fmt.Sprintf(" Too many input data "), testlib.WA)
		return
	}
	if correctAnswer == userLength && correctAnswer == 0 || correctAnswer == userLength && correctAnswer == -1 {
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
		testlib.WriteAnswer(fmt.Sprintf("Full solution "), testlib.OK)
		return
	}
	var n int = core.ReadInt(inputStream)
	graph := createMatrix0(1, n)
	for i := 0; i < n; i++ {
		graph[0][i] = core.ReadInt(inputStream)
	}
	var start int = core.ReadInt(inputStream)
	var finish int = core.ReadInt(inputStream)
	userPath := make([]int, 0)
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.EF)
				return
			}
			break
		}
		temp := core.ReadInt(inputFromUserStream)
		if checkValueOutOfSegmen(temp, 1, n) {
			testlib.WriteAnswer(fmt.Sprintf(" user value out of segment "), testlib.WA)
			return
		}
		userPath = append(userPath, temp)
	}
	if userPath[0] != start || userPath[len(userPath)-1] != finish {
		testlib.WriteAnswer(fmt.Sprintf(" Incorrect path "), testlib.WA)
		return
	}
	calculatedUserPath := 0
	for i := 0; i < len(userPath)-1; i++ {
		if graph[userPath[i]-1][userPath[i+1]-1] == 0 {
			testlib.WriteAnswer(fmt.Sprintf(" Step on not existing way "), testlib.WA)
			return
		}
		calculatedUserPath++
	}
	if calculatedUserPath > correctAnswer {
		testlib.WriteAnswer(fmt.Sprintf(" Incorrect path length calculation "), testlib.WA)
		return
	} else if calculatedUserPath < correctAnswer {
		testlib.WriteAnswer(fmt.Sprintf("WTF?"), testlib.EF)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("Full solution"), testlib.OK)
}
