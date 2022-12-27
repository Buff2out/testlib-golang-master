package meta

import (
	"bufio"
	"os"
	"testlib"
	"testlib/core"
)

var reader *bufio.Reader
var file *os.File

func GetInputReader() *bufio.Reader {
	if reader == nil {
		pipe := GetFile()
		if pipe == nil {
			return nil
		}
		return bufio.NewReader(pipe)
	} else {
		return reader
	}
}
func GetPath() string {
	return testlib.Config.Metadata
}
func GetFile() *os.File {
	if file == nil {
		pipe, err := testlib.GetStreamWithError(testlib.MetaData)
		if err != nil {
			return nil
		}
		return pipe
	} else {
		return file
	}
}

func ReadIntWithError() (int, error) {
	return core.ReadIntWithError(GetInputReader())
}
func ReadInt() int {
	return core.ReadInt(GetInputReader())
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
