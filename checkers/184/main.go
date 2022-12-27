package main

import (
	"fmt"
	"testlib"
	"testlib/core"
)

type node struct {
	x1 int
	y1 int
	x2 int
	y2 int
	l  bool
}

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)

	n1, n2 := core.ReadInt(inputFromUserStream), core.ReadInt(outputStream)

	if n1 > n2 {
		testlib.WriteAnswer(fmt.Sprintf("WA"), testlib.WA)
		return
	}
	if n1 == -1 && n2 == -1 {
		testlib.WriteAnswer("OK", testlib.OK)
		return
	}

	// n, m := 8, 8
	x1, y1, x2, y2 := core.ReadInt(inputStream), core.ReadInt(inputStream), core.ReadInt(inputStream), core.ReadInt(inputStream)
	now := node{
		x1: x1,
		y1: y1,
		x2: x2,
		l:  true,
	}
	f := false
	ff := false
	l3, x3, y3 := 0, 0, 0
	dx, dy := make([]int, 0), make([]int, 0)
	for i := 0; i < n1; i++ {
		l3, x3, y3 = core.ReadInt(inputFromUserStream), core.ReadInt(inputFromUserStream), core.ReadInt(inputFromUserStream)
		l3--
		if l3 == 1 {
			if B2i(now.l) == l3 {
				testlib.WriteAnswer(fmt.Sprintf("WA1"), testlib.WA)
				return
			}
			dx = append(dx, 2, 2, -2, -2, 1, 1, -1, -1)
			dy = append(dy, -1, 1, -1, 1, -2, 2, -2, 2)
			ff = false
			for i := 0; i < 8; i++ {
				if now.x2-x3 == dx[i] && now.y2-y3 == dy[i] {
					ff = true
				}
			}
			if ff == false {
				testlib.WriteAnswer(fmt.Sprintf("WA2"), testlib.WA)
				return
			}
		} else {
			if B2i(now.l) == l3 {
				testlib.WriteAnswer(fmt.Sprintf("WA3"), testlib.WA)
				return
			}
			dx = append(dx, 2, 2, -2, -2, 1, 1, -1, -1)
			dy = append(dy, -1, 1, -1, 1, -2, 2, -2, 2)
			ff = false
			for i := 0; i < 8; i++ {
				if now.x1-x3 == dx[i] && now.y1-y3 == dy[i] {
					ff = true
				}
			}
			if B2i(ff) == 0 {
				testlib.WriteAnswer(fmt.Sprintf("WA4"), testlib.WA)
				return
			}
		}
		if l3 == 1 {
			now.x2 = x3
			now.y2 = y3
		} else {
			now.x1 = x3
			now.y1 = y3
		}
		now.l = !(now.l)
		// cleaning slice:
		dx = dx[:0]
		dy = dy[:0]
	}
	if B2i(f) == 1 {
		testlib.WriteAnswer(fmt.Sprintf("WA5"), testlib.WA)
		return
	}
	if now.x1 != x2 || now.y1 != y2 || now.x2 != x1 || now.y2 != y1 {
		testlib.WriteAnswer(fmt.Sprintf("WA6"), testlib.WA)
		return
	}
	testlib.WriteAnswer("OK", testlib.OK)
}
