package main

import (
	"fmt"
	"io"
	"strings"
	"testlib"
	"testlib/core"
)

func main() {
	flag := false
	//inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	for {
		if flag {
			break
		}
		userLine := ""
		switch line, err := core.ReadLineWithError(inputFromUserStream, true); err {
		case nil:
			userLine = line
			break

		case io.EOF:
			userLine = line
			flag = true
			break
		default:
			testlib.WriteAnswer(err.Error(), testlib.PE)
		}
		switch resultLine, err := core.ReadLineWithError(outputStream, true); err {
		case nil:
			if strings.TrimSpace(resultLine) == strings.TrimSpace(userLine) {
				break
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Strings are not equal. Expected line: '%s'. Got: '%s'", resultLine, userLine), testlib.WA)
			}
		case io.EOF:
			if strings.TrimSpace(resultLine) == strings.TrimSpace(userLine) {
				flag = true
				break
			} else {
				testlib.WriteAnswer(fmt.Sprintf("Strings are not equal. Expected line: '%s'. Got: '%s'", resultLine, userLine), testlib.WA)
			}
		default:
			testlib.WriteAnswer(err.Error(), testlib.PE)
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
	defer func() {
		testlib.GetStream(testlib.TaskOutput).Close()
		testlib.GetStream(testlib.InputFromUser).Close()
	}()
}
