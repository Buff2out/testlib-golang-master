package main

import (
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	for {
		if flag, err := utils.IsStreamFinish(outputStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		userData, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from user stream: %w ", err).Error(), testlib.PE)
			return
		}
		resultData, err := core.ReadIntWithError(outputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from system answer: %w ", err).Error(), testlib.EF)
			return
		}
		if userData != resultData {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are not equal. Expected number: '%d'. Got: '%d", resultData, userData), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
	return
}
