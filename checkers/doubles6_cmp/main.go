package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func almostEqual(userData, resultData, float64EqualityThreshold float64) bool {
	return math.Abs(userData-resultData) > float64EqualityThreshold
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	float64EqualityThreshold := 1e-6
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
		userData, err := core.ReadFloatWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse float from user stream: %w ", err).Error(), testlib.PE)
			return
		}
		resultData, err := core.ReadFloatWithError(outputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from system answer: %w ", err).Error(), testlib.EF)
			return
		}
		if almostEqual(userData, resultData, float64EqualityThreshold) {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are not equal. Expected number: '%f'. Got: '%f", resultData, userData), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
	return
}
