package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
)

func createSlice(n int, inputStream *bufio.Reader) ([]int, error) {
	aSlice := make([]int, 0, n)
	var errf error = nil
	for i := 0; i < n; i++ {
		elSlice, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
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
	//outputStream := testlib.GetReader(testlib.TaskOutput)
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
	m, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	for i := 0; i < m; i++ {
		index, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputFromUserStream: %w ", err).Error(), testlib.PE)
			return
		}
		target, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
			return
		}
		answ, err := core.ReadIntWithError(outputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
			return
		}
		if answ == index {
			continue
		}
		if index < 0 || index > n {
			testlib.WriteAnswer(fmt.Sprintf("Incorrect user output"), testlib.WA)
			return
		}
		if aSlice[index-1] != target {
			testlib.WriteAnswer(fmt.Sprintf("Incorrect user output"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
