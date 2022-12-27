package testlib

import (
	"bufio"
	"fmt"
	"os"
)

type StreamType string

const (
	InputFromUser StreamType = "STDIN"
	OutputToUser  StreamType = "STDOUT"
	MetaData      StreamType = "META"
	TaskInput     StreamType = "INPUT"
	TaskOutput    StreamType = "OUTPUT"
	Answer        StreamType = "ANSWER"
	Error         StreamType = "STDERR"
	CheckConfig   StreamType = "CONFIG"
)

func GetStreamWithError(streamType StreamType) (*os.File, error) {
	switch streamType {
	case InputFromUser:
		return os.Stdin, nil
	case OutputToUser:
		return os.Stdout, nil
	case Answer:
		return os.NewFile(uintptr(3), "answer"), nil
	case MetaData:
		return os.Open(Config.Metadata)
	case TaskInput:
		return os.Open(Config.Input)
	case TaskOutput:
		return os.Open(Config.Output)
	case Error:
		return os.Stderr, nil
	case CheckConfig:
		return os.NewFile(uintptr(4), "config"), nil
	default:
		panic("NOT IMPLEMENTED")
	}
	return nil, nil
}
func GetStream(streamType StreamType) *os.File {
	stream, err := GetStreamWithError(streamType)
	if err != nil {
		WriteAnswer(fmt.Sprintf("Failed to get stream %s. Err: %s", streamType, err.Error()), CF)
		os.Exit(2)
	}
	return stream
}

func ConvertToReader(stream *os.File) *bufio.Reader {
	return bufio.NewReader(stream)
}
func ConvertToWriter(stream *os.File) *bufio.Writer {
	return bufio.NewWriter(stream)
}

func GetReader(streamType StreamType) *bufio.Reader {
	return ConvertToReader(GetStream(streamType))
}
func GetWriter(streamType StreamType) *bufio.Writer {
	return ConvertToWriter(GetStream(streamType))
}
