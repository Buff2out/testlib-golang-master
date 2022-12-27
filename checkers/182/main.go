package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n1 := core.ReadInt(inputFromUserStream)
	n2 := core.ReadInt(outputStream)
	if n1 > n2 {
		testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
		return
	}
	if n1 == -1 {
		if n2 == -1 {
			testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
			return
		} else {
			testlib.WriteAnswer("OK", testlib.OK)
			return
		}
	}
	n := core.ReadInt(inputStream)
	m := core.ReadInt(inputStream)
	x1 := core.ReadInt(inputStream)
	y1 := core.ReadInt(inputStream)
	x2 := core.ReadInt(inputStream)
	y2 := core.ReadInt(inputStream)

	x, y := make([]int, 0), make([]int, 0)

	x3, y3 := 0, 0
	for i := 0; i < n1+1; i++ {
		x3 = core.ReadInt(inputFromUserStream)
		y3 = core.ReadInt(inputFromUserStream)
		x = append(x, x3)
		y = append(y, y3)
	}
	if x[0] != x1 || y[0] != y1 || x[len(x)-1] != x2 || y[len(y)-1] != y2 {
		testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
		return
	}
	dx, dy := 0, 0
	for i := 0; i < len(x)-1; i++ {
		dx, dy = x[i+1]-x[i], y[i+1]-y[i]
		if (!(dx == 2 && dy == 1) && !(dx == 2 && dy == -1) && !(dx == 1 && dy == 2) && !(dx == 1 && dy == -2) && !(dx == -1 && dy == 2) && !(dx == -1 && dy == -2) && !(dx == -2 && dy == -1) && !(dx == -2 && dy == 1)) || x[i] > n || x[i] == 0 || x[i+1] > n || x[i+1] == 0 || y[i] > m || y[i] == 0 || y[i+1] > m || y[i+1] == 0 {
			testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
			return
		}
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
