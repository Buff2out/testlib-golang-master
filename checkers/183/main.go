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

	n1, n2 := core.ReadInt(inputFromUserStream), core.ReadInt(outputStream)

	if n1 == 0 && n2 == 0 {
		testlib.WriteAnswer("OK", testlib.OK)
		return
	}
	if n1 > n2 {
		testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
		return
	}
	if n1 == -1 && n2 == -1 {
		testlib.WriteAnswer("OK", testlib.OK)
		return
	}

	//n, m := 8, 8
	x, y, x5, y5 := core.ReadInt(inputStream), core.ReadInt(inputStream), core.ReadInt(inputStream), core.ReadInt(inputStream)
	x1, y1, x2, y2 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	x1, y1, x2, y2 = append(x1, x), append(y1, y), append(x2, x5), append(y2, y5)
	x3, y3, x4, y4 := 0, 0, 0, 0
	for i := 0; i < n1; i++ {
		x3, y3 = core.ReadInt(inputFromUserStream), core.ReadInt(inputFromUserStream)
		x4, y4 = core.ReadInt(inputFromUserStream), core.ReadInt(inputFromUserStream)
		x1, y1 = append(x1, x3), append(y1, y3)
		x2, y2 = append(x2, x4), append(y2, y4)
	}
	if x1[len(x1)-1] != x2[len(x2)-1] || y1[len(y1)-1] != y2[len(y2)-1] {
		testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
		return
	}
	dx1, dy1 := 0, 0
	dx2, dy2 := 0, 0
	dx, dy := make([]int, 0), make([]int, 0)
	f1, f2 := false, false
	for i := 0; i < len(x1)-1; i++ {
		dx1, dy1 = x1[i+1]-x1[i], y1[i+1]-y1[i]
		dx2, dy2 = x2[i+1]-x2[i], y2[i+1]-y2[i]
		dx = append(dx, 1, 1, -1, -1, 2, 2, -2, -2)
		dy = append(dy, 2, -2, 2, -2, 1, -1, 1, -1)
		f1, f2 = false, false
		for j := 0; j < 8; j++ {
			if dx1 == dx[j] && dy1 == dy[j] {
				f1 = true
			}
			if dx2 == dx[j] && dy2 == dy[j] {
				f2 = true
			}
		}
		if !f1 || !f2 {
			testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
			return
		}
		// cleaning slice:
		dx = dx[:0]
		dy = dy[:0]
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
