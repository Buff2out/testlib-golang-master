package user

import (
	"bufio"
	"testlib"
	"testlib/core"
)

func GetInputReader() *bufio.Reader {
	pipe, err := testlib.GetStreamWithError(testlib.InputFromUser)
	if err != nil {
		return nil
	}
	return bufio.NewReader(pipe)
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
