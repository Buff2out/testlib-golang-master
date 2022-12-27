package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
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

func findMaxKey(s2 map[int]int) int {
	maxKey := -1
	for key, _ := range s2 {
		if maxKey < key {
			maxKey = key
		}
	}
	return maxKey
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n1, n2 := core.ReadInt(outputStream), core.ReadInt(inputFromUserStream)
	if n1 != n2 {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Incorrect first value"), testlib.WA)
	}
	nn1, nn2 := 0, 0
	s := make(map[int]int)
	for i := 0; i < n1; i++ {
		nn1, nn2 = core.ReadInt(outputStream), core.ReadInt(inputFromUserStream)
		if nn1 != nn2 {
			testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Incorrect value"), testlib.WA)
			return
		}
		for j := 0; j < nn1; j++ {
			s[core.ReadInt(outputStream)] = 1
		}
		maxKey := findMaxKey(s)
		for j := 0; j < nn1; j++ {
			if s[core.ReadInt(inputFromUserStream)] == s[maxKey] {
				testlib.WriteAnswer(fmt.Sprintf("Wrong answer. Incorrect value"), testlib.WA)
				return
			}
		}
		// setting all values from set s as 0 (removing from set)
		for key, _ := range s {
			s[key] = 0
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
