package user

import (
	"bufio"
	"testlib"
	"testlib/core"
)

func GetOutputWriter() *bufio.Writer {
	pipe, err := testlib.GetStreamWithError(testlib.OutputToUser)
	if err != nil {
		return nil
	}
	return bufio.NewWriter(pipe)
}

func WriteStringToOutputWithError(str string) error {
	return core.WriteStringToWriterWithError(GetOutputWriter(), str)
}

func WriteStringToOutput(str string) {
	core.WriteStringToWriter(GetOutputWriter(), str)
}
func FlushOutputWithError() error {
	return core.FlushWriterWithError(GetOutputWriter())
}
func FlushOutput() {
	core.FlushWriter(GetOutputWriter())
}
