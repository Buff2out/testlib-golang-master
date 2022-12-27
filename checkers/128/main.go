package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

func f1(x, y float64) float64 {
	return x*x + y*y - 0.1*math.Abs(1-x) - 0.1*math.Abs(1-y)
}

func f2(x, y float64) float64 {
	return 20*math.Abs(x-50)*math.Abs(y-25) + 10*(math.Abs(x-10)*math.Abs(y-10)+math.Abs(x-50))
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
		_ = testlib.GetStream(testlib.OutputToUser).Close()
		_ = testlib.GetStream(testlib.TaskInput).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	n := core.ReadInt(inputStream)
	a := make([]int, 0)
	cnt := 0
	answer := 0
	userAnswer := 0
	var requestType byte
	core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%d\n", n))
	core.FlushWriter(outputToUserStream)
	for i := 0; i < n; i++ {
		a = append(a, core.ReadInt(inputStream))
	}
	answer = core.ReadInt(outputStream)
	for {
		// if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
		// 	if err != nil {
		// 		testlib.WriteAnswer(err.Error(), testlib.EF)
		// 		return
		// 	}
		// 	break
		// }
		if cnt >= 50 {
			testlib.WriteAnswer(fmt.Sprintf("Too many requests"), testlib.WA)
			return
		}
		requestType = core.ReadByte(inputFromUserStream)
		switch requestType {
		case byte('!'):
			userAnswer = core.ReadInt(inputFromUserStream)
			if userAnswer != answer {
				testlib.WriteAnswer(fmt.Sprintf("Incorrect user output. User answer: %d. Correct answer: %d", userAnswer, answer), testlib.WA)
				return
			}
			testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
			return
			break
		case byte('?'):
			userAnswer = core.ReadInt(inputFromUserStream)
			if checkValueOutOfSegmen(userAnswer, 1, n) {
				testlib.WriteAnswer(fmt.Sprintf("User value %d Out of segment 1 and %d", userAnswer, n), testlib.WA)
				return
			}
			if userAnswer > 0 && userAnswer <= n {
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%d\n", a[userAnswer-1]))
				core.FlushWriter(outputToUserStream)
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Wrong index "), testlib.PE)
				return
			}
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Incorrect input command: %s", string(requestType)), testlib.PE)
			return
		}
		cnt++
	}
	testlib.WriteAnswer(fmt.Sprintf("Incorrect: Input after EOF"), testlib.EF)
	return
}
