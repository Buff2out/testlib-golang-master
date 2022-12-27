packege main

import (
	"fmt"
	"testlib"
	"testlib/core"
	"testlib/utils"
)

func main() {
	defer func() {
		testlib.GetStream(testlib.TaskOutput).Close()
		testlib.GetStream(testlib.InputFromUser).Close()
	}()
	testlib.WriteAnswer("OK", testlib.OK)
}