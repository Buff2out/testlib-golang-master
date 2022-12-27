package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testlib"
	"testlib/core"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func createMatrix0(n int, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func checkValueOutOfSegmen(val int, a int, b int) bool {
	if val < a || val > b {
		return true
	} else {
		return false
	}
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
		_ = testlib.GetStream(testlib.OutputToUser).Close()
		_ = testlib.GetStream(testlib.TaskInput).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	// outputToUserStream := testlib.GetWriter(testlib.OutputToUser)

	n := core.ReadInt(inputStream)
	mp := make(map[int]int)
	keys_mp := make([]int, 0)
	for i := 0; i < n; i++ {
		varType := core.ReadInt(inputStream)
		if varType == 1 {
			key := core.ReadInt(inputStream)
			val := core.ReadInt(inputStream)
			mp[key] += val
			keys_mp = append(keys_mp, key)
			keys_mp = removeDuplicateValues(keys_mp)
			sort.Slice(keys_mp, func(i2, j int) bool {
				return keys_mp[i2] < keys_mp[j]
			})
			continue
		}
		_ = core.ReadInt(inputStream)
		a := make([]int, 0)
		b := make([]int, 0)
		aCheckDupl := make(map[int]bool)
		bCheckDupl := make(map[int]bool)
		lineA := strings.Split(core.ReadLine(outputStream, true), " ")
		lineB := strings.Split(core.ReadLine(inputFromUserStream, true), " ")
		for i3 := 0; i3 < len(lineA); i3++ {
			toA, _ := strconv.Atoi(lineA[i3])
			if aCheckDupl[toA] {
				testlib.WriteAnswer(fmt.Sprintf("Some key was printed more than once"), testlib.EF)
				return
			}
			aCheckDupl[toA] = true
			a = append(a, toA)
		}
		for i3 := 0; i3 < len(lineB); i3++ {
			toB, _ := strconv.Atoi(lineB[i3])
			if bCheckDupl[toB] {
				testlib.WriteAnswer(fmt.Sprintf("Some key was printed more than once"), testlib.WA)
				return
			}
			bCheckDupl[toB] = true
			b = append(b, toB)
		}
		sort.Slice(a, func(i2, j int) bool {
			return a[i2] < a[j]
		})
		sort.Slice(b, func(i2, j int) bool {
			return b[i2] < b[j]
		})
		if len(a) != len(b) {
			testlib.WriteAnswer(fmt.Sprintf("Size of keys is not corrected"), testlib.WA)
			return
		}
		for i3 := 0; i3 < len(a); i3++ {
			if a[i3] == keys_mp[len(keys_mp)-1] {
				testlib.WriteAnswer(fmt.Sprintf("EF check that all keys are the keys that was already given, not arbitary ones IS FAILED ."), testlib.EF)
				return
			}
			if b[i3] == keys_mp[len(keys_mp)-1] {
				testlib.WriteAnswer(fmt.Sprintf("WA check that all keys are the keys that was already given, not arbitary ones IS FAILED ."), testlib.WA)
				return
			}
		}
		// transform
		for j1 := 0; j1 < len(a); j1++ {
			temp := mp[a[j1]]
			a[j1] = temp
		}
		for j1 := 0; j1 < len(b); j1++ {
			temp := mp[b[j1]]
			b[j1] = temp
		}
		sort.Slice(a, func(i2, j int) bool {
			return a[i2] < a[j]
		})
		sort.Slice(b, func(i2, j int) bool {
			return b[i2] < b[j]
		})
		for j1 := 0; j1 < len(b); j1++ {
			if b[j1] != a[j1] {
				testlib.WriteAnswer(fmt.Sprintf("Solution not optimal."), testlib.WA)
				return
			}
		}
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
	return
}
