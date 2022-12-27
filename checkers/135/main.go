package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func main() {
	defer func() {
		testlib.GetStream(testlib.OutputToUser).Close()
		testlib.GetStream(testlib.InputFromUser).Close()
		//testlib.GetStream(testlib.TaskOutput).Close()
		testlib.GetStream(testlib.TaskInput).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput) // эквивалент
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	n := core.ReadInt(inputStream)
	var cnt int = 0
	//answer := core.ReadInt(outputStream)
	a := make([]int, 0)
	var requestType byte
	var res int
	var userAnswer int
	core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%d\n", n))
	core.FlushWriter(outputToUserStream)
	for i := 0; i < n; i++ {
		a = append(a, core.ReadInt(inputStream))
	}
	for {
		if cnt >= 50 {
			testlib.WriteAnswer("Too many requests", testlib.WA)
			return
		}
		requestType = core.ReadByte(inputFromUserStream)
		// isFind := false
		switch requestType {
		case byte('!'):
			userAnswer = core.ReadInt(inputFromUserStream)
			res = userAnswer - 1
			if res == 0 {
				if a[0] >= a[1] {
					testlib.WriteAnswer("OK", testlib.OK)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("Element isn't a peak one"), testlib.WA)
					return
				}
			} else if res == n-1 {
				if a[n-1] >= a[n-2] {
					testlib.WriteAnswer("OK", testlib.OK)
					return
				} else {
					testlib.WriteAnswer(fmt.Sprintf("Element isn't a peak one"), testlib.WA)
					return
				}
			} else if a[res] >= a[res-1] && a[res] >= a[res+1] {
				testlib.WriteAnswer("OK", testlib.OK)
				return
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Element isn't a peak one"), testlib.WA)
				return
			}
		case byte('?'):
			userAnswer := core.ReadInt(inputFromUserStream)
			if userAnswer > 0 && userAnswer <= n {
				core.WriteStringToWriter(outputToUserStream, fmt.Sprintf("%d\n", a[userAnswer-1]))
				core.FlushWriter(outputToUserStream)
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Index is out of range. User index: %d. Data length: %d", userAnswer, len(a)), testlib.PE)
				return
			}
			break
		default:
			testlib.WriteAnswer(fmt.Sprintf("Invalid request type: %s", string(requestType)), testlib.PE)
			return
		}
		cnt++
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
