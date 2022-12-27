package main

import (
	"bufio"
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func createMultiset(n int, inputStream *bufio.Reader) (map[int]int, error) {
	aMultiset := make(map[int]int)
	var errf error = nil
	for i := 0; i < n; i++ {
		el, err := core.ReadIntWithError(inputStream)
		if err != nil {
			testlib.WriteAnswer(fmt.Errorf("Failed to parse int from user stream: %w ", err).Error(), testlib.PE)
			errf = err
			break
		}
		value, isMapContainsKey := aMultiset[el]
		if isMapContainsKey {
			aMultiset[el] = 1
		} else {
			aMultiset[el] = value + 1
		}
	}
	return aMultiset, errf
}

//troubled checker

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
	aMultiset, err := createMultiset(n, inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	m, err := core.ReadIntWithError(inputStream)
	if err != nil {
		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from inputStream: %w ", err).Error(), testlib.PE)
		return
	}
	bMultiset := make(map[int]int)
	cnt, cnt1 := 0, 0
	i1 := 0
	for {
		if flag, err := utils.IsStreamFinish(outputStream, inputFromUserStream); flag {
			if err != nil {
				testlib.WriteAnswer(err.Error(), testlib.WA)
				return
			}
			break
		}
		// for j1 := 0; j1 < m; j1++ {
		//     x, err := core.ReadIntWithError(outputStream)
		// 	if err != nil {
		// 		testlib.WriteAnswer(fmt.Errorf("Failed to parse int from output Stream: %w ", err).Error(), testlib.PE)
		// 		return
		// 	}
		// }
		cnt++
		i1++
		for j2 := 0; j2 < m; j2++ {
			// cntmap[x]++;
			x, err := core.ReadIntWithError(inputFromUserStream)
			if err != nil {
				testlib.WriteAnswer(fmt.Errorf("Failed to parse int from user stream: %w ", err).Error(), testlib.PE)
				return
			}
			value, isMapContainsKey := bMultiset[x]
			if isMapContainsKey {
				bMultiset[x] = 1
			} else {
				bMultiset[x] = value + 1
			}
			//end of cntmap[x]++;
			if bMultiset[x] > aMultiset[x] {
				testlib.WriteAnswer(fmt.Sprintf("Map: '%d'; Count: '%d'; Number: '%d'", bMultiset[x], aMultiset[x], x), testlib.WA)
				return
			}
		}
		cnt1++
	}
	if cnt1 < cnt {
		testlib.WriteAnswer(fmt.Sprintf("Users count less than correct"), testlib.WA)
		return
	}
	if cnt1 == cnt {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	if cnt1 > cnt {
		testlib.WriteAnswer(fmt.Sprintf("Users count greater than correct"), testlib.WA)
		return
	}
	return
}
