package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func createBoolSlice(n int, b bool, inputStream *bufio.Reader) []bool {
	aSlice := make([]bool, 0, n)
	for i := 0; i < n; i++ {
		aSlice = append(aSlice, b)
	}
	return aSlice
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	// n, err := core.ReadIntWithError(inputStream)
	// if err != nil {
	// 	testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
	// 	return
	// }
	n := core.ReadInt(inputStream)
	w := core.ReadInt(inputStream)
	aSlice := make([]int, 0)
	bSlice := make([]int, 0)
	cSlice := createBoolSlice(n, false, inputStream)
	for i := 0; i < n; i++ {
		aSlice = append(aSlice, core.ReadInt(inputStream))
	}
	for i := 0; i < n; i++ {
		bSlice = append(bSlice, core.ReadInt(inputStream))
	}
	rescnt := 0
	for {
		if flag, err := utils.IsStreamFinish(outputStream, outputStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		x, err := core.ReadIntWithError(outputStream)
		if err != nil || x < 1 || x > n {
			testlib.WriteAnswer(fmt.Errorf("error to parse or x < 1 || x > n error: %w, x: %d ", err, x).Error(), testlib.PE)
			return
		}
		rescnt += aSlice[x-1]
	}
	oufcnt := 0
	oufw := 0
	for {
		if flag, err := utils.IsStreamFinish(inputFromUserStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		x, err := core.ReadIntWithError(outputStream)
		if err != nil || x < 1 || x > n {
			testlib.WriteAnswer(fmt.Errorf("error to parse or x < 1 || x > n error: %w, x: %d ", err, x).Error(), testlib.PE)
			return
		}
		if cSlice[x-1] {
			testlib.WriteAnswer(fmt.Sprintf("Same ingot"), testlib.WA)
			return
		}
		cSlice[x-1] = true
		oufcnt += aSlice[x-1]
		oufw += bSlice[x-1]
		rescnt += aSlice[x-1]
		if oufw > w {
			testlib.WriteAnswer(fmt.Sprintf("Rucksack overflow"), testlib.WA)
			return
		}
	}
	if oufcnt < rescnt {
		testlib.WriteAnswer(fmt.Sprintf("Wrong answer"), testlib.WA)
		return
	}
	if oufcnt == rescnt {
		testlib.WriteAnswer("Full solution", testlib.OK)
		return
	}
	if oufcnt > rescnt {
		testlib.WriteAnswer(fmt.Errorf("Contestant output better than jury's").Error(), testlib.EF)
		return
	}
}
