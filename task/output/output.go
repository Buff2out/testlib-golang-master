package output

import (
	"bufio"
	"testlib"
	"testlib/core"
)

var reader *bufio.Reader

func init() {
	pipe, err := testlib.GetStreamWithError(testlib.TaskOutput)
	if err != nil {
		reader = nil
	}
	reader = bufio.NewReader(pipe)
}

func GetInputReader() *bufio.Reader {
	return reader
}

func ReadIntWithError() (int, error) {
	return core.ReadIntWithError(GetInputReader())
}
func ReadInt() int {
	return core.ReadInt(GetInputReader())
}

func ReadByteWithError() (byte, error) {
	return core.ReadByteWithError(GetInputReader())
}
func ReadByte() byte {
	return core.ReadByte(GetInputReader())
}

func ReadWordWithError() (string, error) {
	return core.ReadWordWithError(GetInputReader())
}
func ReadWord() string {
	return core.ReadWord(GetInputReader())
}

func ReadFloatWithError() (float64, error) {
	return core.ReadFloatWithError(GetInputReader())
}
func ReadFloat() float64 {
	return core.ReadFloat(GetInputReader())
}

func ReadBoolWithError() (bool, error) {
	return core.ReadBoolWithError(GetInputReader())
}
func ReadBool() bool {
	return core.ReadBool(GetInputReader())
}

func ReadLineWithError(ignoreEmptyLines bool) (string, error) {
	return core.ReadLineWithError(GetInputReader(), ignoreEmptyLines)
}
func ReadLine(ignoreEmptyLines bool) string {
	return core.ReadLine(GetInputReader(), ignoreEmptyLines)
}
