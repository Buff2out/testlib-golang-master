package core

import (
	"bufio"
	"fmt"
	"os"
	"testlib"
)

func WriteStringToWriterWithError(writer *bufio.Writer, str string) error {
	_, err := writer.WriteString(str)
	return err
}
func WriteStringToWriter(writer *bufio.Writer, str string) {
	err := WriteStringToWriterWithError(writer, str)
	if err != nil {
		testlib.WriteAnswer(fmt.Sprintf("Fail on write to writer. Err: %s", err.Error()), testlib.CF)
		os.Exit(2)
	}
}
func FlushWriterWithError(writer *bufio.Writer) error {
	return writer.Flush()
}
func FlushWriter(writer *bufio.Writer) {
	err := FlushWriterWithError(writer)
	if err != nil {
		testlib.WriteAnswer(fmt.Sprintf("Failed to flush writer. Err: %s ", err.Error()), testlib.CF)
		os.Exit(2)
	}
}
