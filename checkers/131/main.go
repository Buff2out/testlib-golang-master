package main

import (
	"bufio"
	"fmt"
	"io"

	//"strings"
	"testlib"
	"testlib/core"
)

func createSlice(n int, inputStream *bufio.Reader) ([]string, error) {
	aSlice := make([]string, 0, n)
	var errf error = nil
	for i := 0; i < n; i++ {
		elSlice, err := core.ReadWordWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
			errf = err
			break
		}
		aSlice = append(aSlice, elSlice)
	}
	return aSlice, errf
}

func findStringInLinearTime(n int, aSlice []string, target string) int {
	indx := -1
	for i := 0; i < n; i++ {
		if aSlice[i] == target {
			indx = i
			break
		}
	}
	return indx
}

func main() {
	defer func() {
		testlib.GetStream(testlib.TaskOutput).Close()
		testlib.GetStream(testlib.InputFromUser).Close()
	}()
	//flag := false
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	n, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.EF)
		return
	}
	aSlice, err := createSlice(n, inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	m, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.EF)
		return
	}
	for i := 0; i < m; i++ {
		target := ""
		switch line, err := core.ReadWordWithError(inputStream); err {
		case nil:
			target = line
			break

		case io.EOF:
			target = line
			//flag = true
			break
		default:
			testlib.WriteAnswer(err.Error(), testlib.EF)
		}
		answ, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputFromUserStream: %w ", err).Error(), testlib.PE)
			return
		}
		if (answ < 1 || answ > n) && (answ != -1) {
			testlib.WriteAnswer(fmt.Errorf("Presentation error. Incorrect index").Error(), testlib.PE)
			return
		}
		if (answ == -1) && (aSlice[findStringInLinearTime(n, aSlice, target)] != aSlice[len(aSlice)-1]) {
			testlib.WriteAnswer(fmt.Errorf("Presentation error. Can't find word with this index").Error(), testlib.PE)
			return
		}
		if answ == -1 {
			continue
		}
		if aSlice[answ-1] != target {
			testlib.WriteAnswer(fmt.Errorf("Incorrect user output").Error(), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
