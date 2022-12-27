package main

import (
	"fmt"
	"math"
	"sort"
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

var eps float64 = 5 * 1e-3

func subtrPairFloats64(a, b PairFloat64) PairFloat64 {
	c := PairFloat64{
		first:  a.first - b.first,
		second: a.second - b.second,
	}
	return c
}

func sumPairFloats64(a, b PairFloat64) PairFloat64 {
	c := PairFloat64{
		first:  a.first + b.first,
		second: a.second + b.second,
	}
	return c
}

func quotientPairFloats64(a, b PairFloat64) PairFloat64 {
	c := PairFloat64{
		first:  a.first / b.first,
		second: a.second / b.second,
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

func CheckBySinusTheoreme(sortedSideLengths, sortedAngles []float64) bool {
	//flag := false
	var tempValue float64 = sortedSideLengths[0] / math.Sin(sortedAngles[0]*math.Pi/180)
	for i := 1; i < 3; i++ {
		if math.Abs(sortedSideLengths[i]/math.Sin(sortedAngles[i]*math.Pi/180)-tempValue) > eps {
			return false
		}
	}
	return true
}

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func f2b(f_0 float64) bool {
	var eps float64 = 1e-6
	if -eps < f_0 && f_0 < eps {
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
		_ = testlib.GetStream(testlib.TaskInput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	inputStream := testlib.GetReader(testlib.TaskInput)
	//outputStream := testlib.GetReader(testlib.TaskOutput)
	inputFromUserStream := testlib.GetReader(testlib.InputFromUser)
	//outputToUserStream := testlib.GetWriter(testlib.OutputToUser)
	lena, anglea, angleb := core.ReadFloat(inputStream), core.ReadFloat(inputStream), core.ReadFloat(inputStream)
	anglec := 180.0 - anglea - angleb
	// anglea, angleb := core.ReadFloat(inputStream), core.ReadFloat(inputStream)
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
	//t := make([]PairFloat64, 0)
	sortedAngles := make([]float64, 0)
	sortedSideLengths := make([]float64, 0)
	gammaAngleIndex := 0
	sortedAngles = append(sortedAngles, anglea)
	sortedAngles = append(sortedAngles, angleb)
	sortedAngles = append(sortedAngles, anglec)
	sort.Slice(sortedAngles, func(i, j int) bool {
		return sortedAngles[i] < sortedAngles[j]
	})
	sortedSideLengths = append(sortedSideLengths, lenx(subtrPairFloats64(b, a)))
	sortedSideLengths = append(sortedSideLengths, lenx(subtrPairFloats64(c, a)))
	sortedSideLengths = append(sortedSideLengths, lenx(subtrPairFloats64(c, b)))
	sort.Slice(sortedSideLengths, func(i, j int) bool {
		return sortedSideLengths[i] < sortedSideLengths[j]
	})

	for i := 0; i < 3; i++ {
		if sortedAngles[i] == anglec {
			gammaAngleIndex = i
			break
		}
	}
	if !CheckBySinusTheoreme(sortedSideLengths, sortedAngles) {
		testlib.WriteAnswer(fmt.Sprintf(" Wrong answer.Does not match the sinus theorem "), testlib.OK)
		return
	}
	var notMedianAngleIndex int = (gammaAngleIndex + 1) % 3
	var medianSideLength float64 = sortedSideLengths[gammaAngleIndex] / 2
	var lastAngleIndex = (notMedianAngleIndex + 1) % 3

	if math.Abs((medianSideLength*medianSideLength+sortedSideLengths[lastAngleIndex]*sortedSideLengths[lastAngleIndex]-lena*lena)/(2*medianSideLength*sortedSideLengths[lastAngleIndex])-math.Cos(sortedAngles[notMedianAngleIndex]*math.Pi/180)) < eps {
		testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
		return
	}
	testlib.WriteAnswer(fmt.Sprintf("Wrong answer.Does not match the cosinus theorem "), testlib.WA)
}
