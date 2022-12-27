package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func createSlice(n int, inputStream *bufio.Reader) ([]int64, error) {
	aSlice := make([]int64, 0, n)
	var errf error = nil
	for i := 0; i < n; i++ {
		elSlice, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
			errf = err
			break
		}
		aSlice = append(aSlice, int64(elSlice))
	}
	return aSlice, errf
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	aSlice, err := createSlice(n, inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	var cnt1, cnt2 int64
	for {
		if flag, err := utils.IsStreamFinish(outputStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		ind, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
			return
		} else if ind < 1 && ind > n {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
			return
		}
		for i := 0; i < ind-1; i++ {
			cnt1 += aSlice[i]
		}
		for i := ind; i < n; i++ {
			cnt2 += aSlice[i]
		}
		if cnt1 == cnt2 {
			testlib.WriteAnswer("OK", testlib.OK)
			break
		} else {
			testlib.WriteAnswer(fmt.Sprintf("First part: '%d'; Second part: '%d", cnt1, cnt2), testlib.WA)
			return
		}
	}
	return
}
