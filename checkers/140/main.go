package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	// inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	out := core.ReadWord(inputFromUserStream)
	ans := core.ReadWord(outputStream)

	for out[len(out)-1] == 32 {
		out = out[0 : len(out)-1]
	}
	for ans[len(ans)-1] == 32 {
		ans = ans[0 : len(ans)-1]
	}
	if ans == out {
		testlib.WriteAnswer("Full solution", testlib.OK)
	} else {
		testlib.WriteAnswer(fmt.Sprintf("Incorrect user output. User answer: %d. Correct answer: %d", out, ans), testlib.WA)
	}
}
