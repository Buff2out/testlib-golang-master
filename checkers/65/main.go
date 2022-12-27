package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func createSlice(n int, inputStream *bufio.Reader) ([]int, error) {
	aSlice := make([]int, 0, n)
	var errf error = nil
	for i := 0; i < n; i++ {
		elSlice, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from user stream: %w ", err).Error(), testlib.PE)
			errf = err
			break
		}
		aSlice = append(aSlice, elSlice)
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
	sum, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	aSlice, err := createSlice(n, inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}

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
		if resultData == -1 && userData == -1 {
			break
		}
		// assingnment i1, i2
		i1, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
			return
		} else if 1 < i1 && i1 > n {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
			return
		}
		i2, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
			return
		} else if 1 < i2 && i2 > n {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
			return
		} else if i1 == i2 {
			testlib.WriteAnswer(fmt.Sprintf("First variable equal second"), testlib.WA)
			return
		}
		// end of assingnment i1, i2
		if (aSlice[i1-1] + aSlice[i2-1]) != sum {
			testlib.WriteAnswer(fmt.Sprintf("Incorrect sum.Get  '%d'. Correct answer: '%d", aSlice[i1-1]+aSlice[i2-1], sum), testlib.WA)
			return
		}
		// if userData != resultData {
		// 	testlib.WriteAnswer(fmt.Sprintf("Numbers are not equal. Expected number: '%d'. Got: '%d", resultData, userData), testlib.WA)
		// 	return
		// }
	}
	testlib.WriteAnswer("OK", testlib.OK)
	return
}
