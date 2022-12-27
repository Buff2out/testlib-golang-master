package main

import (
	"fmt"
	"math"
	"testlib"
	"testlib/core"
)

type Pair struct {
	first  int
	second int
}
type PairFloat64 struct {
	first  float64
	second float64
}

func subtrPairFloats64(a, b PairFloat64) PairFloat64 {
	c := PairFloat64{
		first:  a.first - b.first,
		second: a.second - b.second,
	}
	return c
}

func dist(a, b PairFloat64) float64 {
	return math.Sqrt(math.Pow((a.first-b.first), 2) + math.Pow((a.second-b.second), 2))
}

func lenx(x PairFloat64) float64 {
	return math.Sqrt(x.first*x.first + x.second*x.second)
}

func angleaa(a, b PairFloat64) float64 {
	return math.Acos((a.first*b.first + a.second*b.second) / (lenx(a) * lenx(b)))
}

func ar_tr(a, b PairFloat64) float64 {
	return a.first*b.second - b.first*a.second
}

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func createMatrix0(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func checkValueOutOfSegmen(val, a, b int) bool {
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
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	var eps float64 = 1e-6
	lena, lenb := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	angle := core.ReadFloat(inputStream)
	a := PairFloat64{
		first:  core.ReadFloat(inputFromUserStream),
		second: core.ReadFloat(inputFromUserStream),
	}
	b := PairFloat64{
		first:  core.ReadFloat(inputFromUserStream),
		second: core.ReadFloat(inputFromUserStream),
	}
	c := PairFloat64{
		first:  core.ReadFloat(inputFromUserStream),
		second: core.ReadFloat(inputFromUserStream),
	}
	t := make([]PairFloat64, 0)
	t = append(t, a)
	t = append(t, b)
	t = append(t, c)
	for i := 0; i < 3; i++ {
		first := subtrPairFloats64(t[(i+1)%3], t[i])
		second := subtrPairFloats64(t[(i+2)%3], t[i])
		if ((math.Abs(lenx(first)-lena) < eps && math.Abs(lenx(second)-lenb) < eps) || (math.Abs(lenx(first)-lenb) < eps && math.Abs(lenx(second)-lena) < eps)) && math.Abs(angleaa(first, second)*180/math.Pi-angle) < eps {
			testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
			return
		}
	}
	testlib.WriteAnswer(fmt.Sprintf("Wrong answer"), testlib.WA)
}
