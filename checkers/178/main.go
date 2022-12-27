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

	n1, n2, x := 0, 0, 0
	s1, s2 := make(map[int]int), make(map[int]int)

	for {
		if flag, err := utils.IsStreamFinish(outputStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
		}
		x = core.ReadInt(outputStream)
		n1++
		s1[x] = 1
	}
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}

		}
		x = core.ReadInt(inputFromUserStream)
		n2++
		s2[x] = 1
	}
	if n1 != n2 {
		testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
		return
	}
	maxKey := findMaxKey(s2)
	for key, _ := range s1 {
		if s2[key] == s2[maxKey] {
			testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("Full solution", testlib.OK)
}
