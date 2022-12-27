package utils

import (
	"bufio"
	"fmt"
	"testlib/core"
)

// IsStreamFinish Return error message if only one of the streams are finished
func IsStreamFinish(taskOutputStream, userInputStream *bufio.Reader) (bool, error) {
	if core.ReaderIsEmpty(taskOutputStream) && !core.ReaderIsEmpty(userInputStream) {
		return true, fmt.Errorf("Solution has more data than exceeded ")
	}
	if core.ReaderIsEmpty(userInputStream) && !core.ReaderIsEmpty(taskOutputStream) {
		return true, fmt.Errorf("Solution has less data than exceeded ")
	}
	if core.ReaderIsEmpty(taskOutputStream) && core.ReaderIsEmpty(userInputStream) {
		return true, nil
	}
	return false, nil
}
