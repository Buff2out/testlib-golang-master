package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func main() {
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	userOutputStream := testlib.GetWriter(testlib.OutputToUser)
	userInputStream := testlib.GetReader(testlib.InputFromUser)
	n := core.ReadInt(inputStream)
	var data []int
	core.WriteStringToWriter(userOutputStream, fmt.Sprintf("%d\n", n))
	core.FlushWriter(userOutputStream)
	for i := 0; i < n; i++ {
		data = append(data, core.ReadInt(inputStream))
	}
	m := core.ReadInt(inputStream)
	core.WriteStringToWriter(userOutputStream, fmt.Sprintf("%d\n", m))
	core.FlushWriter(userOutputStream)
	for i := 0; i < m; i++ {
		cnt := 0
		val := core.ReadInt(inputStream)
		core.WriteStringToWriter(userOutputStream, fmt.Sprintf("%d\n", val))
		core.FlushWriter(userOutputStream)
		answer := core.ReadInt(outputStream)
		for {
			if cnt >= 50 {
				testlib.WriteAnswer("Too many requests", testlib.WA)
			}
			requestType := core.ReadByte(userInputStream)
			isFind := false
			switch requestType {
			case byte('!'):
				userAnswer := core.ReadInt(userInputStream)
				if userAnswer != answer {
					testlib.WriteAnswer(fmt.Sprintf("Incorrect user output. User answer: %d. Correct answer: %d", userAnswer, answer), testlib.WA)
				}
				isFind = true
				break
			case byte('?'):
				userAnswer := core.ReadInt(userInputStream)
				if userAnswer > 0 && userAnswer <= n {
					core.WriteStringToWriter(userOutputStream, fmt.Sprintf("%d\n", data[userAnswer-1]))
					core.FlushWriter(userOutputStream)
				} else {
					testlib.WriteAnswer(fmt.Sprintf("Index is out of range. User index: %d. Data length: %d", userAnswer, len(data)), testlib.PE)
				}
				break
			default:
				testlib.WriteAnswer(fmt.Sprintf("Invalid request type: %s", string(requestType)), testlib.PE)
			}
			if isFind {
				break
			}
			cnt++
		}
	}
	testlib.WriteAnswer("Accepted", testlib.OK)
	defer func() {
		testlib.GetStream(testlib.OutputToUser).Close()
		testlib.GetStream(testlib.InputFromUser).Close()
	}()
}
