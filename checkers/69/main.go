package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func create2Slices(n int, inputStream *bufio.Reader) ([]int, []int, error) {
	aSlice1 := make([]int, 0, n)
	aSlice2 := make([]int, 0, n)
	var errf error = nil
	for i := 0; i < n; i++ {
		elSlice, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
			errf = err
			break
		}
		aSlice1 = append(aSlice1, elSlice)
		aSlice2 = append(aSlice1, elSlice)
	}
	return aSlice1, aSlice2, errf
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
	aSlice, bSlice, err := create2Slices(n, inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from user stream: %w ", err).Error(), testlib.PE)
		return
	}
	cnt1, cnt2, max1, max2 := 0, 0, 0, 0
	for {
		if flag, err := utils.IsStreamFinish(outputStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		ind2, err := core.ReadIntWithError(outputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from outputStream: %w ", err).Error(), testlib.PE)
			return
		}
		ind1, err := core.ReadIntWithError(inputFromUserStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to read data from inputFromUserStream: %w ", err).Error(), testlib.EF)
			return
		} else if ind1 < 1 && ind1 > n {
			testlib.WriteAnswer(fmt.Sprintf("Numbers are out of range"), testlib.WA)
			return
		}
		aSlice[ind1-1] = 1
		bSlice[ind2-1] = 1
		for i := 0; i < n; i++ {
			if aSlice[i] == 0 {
				max1 = Max(max1, cnt1)
				cnt1 = 0
			} else {
				cnt1++
			}
		}
		max1 = Max(max1, cnt1)
		for i := 0; i < n; i++ {
			if bSlice[i] == 0 {
				max2 = Max(max2, cnt2)
				cnt2 = 0
			} else {
				cnt2++
			}
		}
		max2 = Max(max2, cnt2)
		if max1 < max2 {
			testlib.WriteAnswer(fmt.Sprintf("User max value  '%d' less than correct max value '%d", max1, max2), testlib.WA)
			return
		}
		if max1 == max2 {
			testlib.WriteAnswer("OK", testlib.OK)
			break
		}
		if max1 > max2 {
			testlib.WriteAnswer(fmt.Sprintf("STUDENT SOLUTION RESULT '%d' BETTER THAN SYSTEM '%d", max1, max2), testlib.EF)
			return
		}
	}
	return
}
