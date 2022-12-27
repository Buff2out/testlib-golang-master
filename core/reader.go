package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testlib"
)

var emptyLineError = fmt.Errorf("Invalid data. String is empty. ")

func ReadIntWithError(reader *bufio.Reader) (int, error) {
	word, err := ReadWordWithError(reader)
	if err != nil {
		if err == io.EOF {
			return 0, err
		}
		return 0, fmt.Errorf("Invalid data. Can't read data. Error: %w ", err)
	}
	result, err := strconv.Atoi(word)
	if err != nil {
		return 0, fmt.Errorf("Invalid data. Can't read int. Get: %s. Error: %w ", word, err)
	}
	return result, nil
}
func ReadInt(reader *bufio.Reader) int {
	result, err := ReadIntWithError(reader)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReadByteWithError(reader *bufio.Reader) (byte, error) {
	return reader.ReadByte()
}
func ReadByte(reader *bufio.Reader) byte {
	result, err := ReadByteWithError(reader)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReadWordWithError(reader *bufio.Reader) (string, error) {
	resultString := ""
	for {
		byteData, err := ReadByteWithError(reader)
		emptyLine := false
		if err != nil {
			if err == io.EOF {
				if len(resultString) == 0 {
					return "", err
				}
				return resultString, nil
			}
			return resultString, err
		}
		for _, separator := range testlib.Separators {
			if byteData == separator {
				if len(resultString) == 0 {
					emptyLine = true
					break
				}
				return resultString, nil
			}
		}
		if emptyLine {
			continue
		}
		resultString += string(byteData)
	}
}
func ReadWord(reader *bufio.Reader) string {
	result, err := ReadWordWithError(reader)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReadLineWithError(reader *bufio.Reader, ignoreEmptyLines bool) (string, error) {
	resultString := ""
	for {
		byteData, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				if len(resultString) == 0 {
					return resultString, err
				}
				return resultString, nil
			}
			return resultString, err
		}
		//Ignore carriage return symbol
		if byteData == '\r' {
			continue
		}
		if byteData == '\n' {
			if len(resultString) == 0 {
				if ignoreEmptyLines {
					continue
				}
				return resultString, emptyLineError
			}
			return resultString, nil
		}
		resultString += string(byteData)
	}
}
func ReadLine(reader *bufio.Reader, ignoreEmptyLines bool) string {
	result, err := ReadLineWithError(reader, ignoreEmptyLines)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReadFloatWithError(reader *bufio.Reader) (float64, error) {
	word, err := ReadWordWithError(reader)
	if err != nil {
		return 0, fmt.Errorf("Invalid data. Can't read data. Error: %w ", err)
	}
	result, err := strconv.ParseFloat(word, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid data. Can't read float. Get: %s. Error: %w ", word, err)
	}
	return result, nil
}
func ReadFloat(reader *bufio.Reader) float64 {
	result, err := ReadFloatWithError(reader)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReadBoolWithError(reader *bufio.Reader) (bool, error) {
	word, err := ReadWordWithError(reader)
	if err != nil {
		return false, fmt.Errorf("Invalid data. Can't read data. Error: %w ", err)
	}
	result, err := strconv.ParseBool(word)
	if err != nil {
		return false, fmt.Errorf("Invalid data. Can't read bool. Get: %s. Error: %w ", word, err)
	}
	return result, nil
}
func ReadBool(reader *bufio.Reader) bool {
	result, err := ReadBoolWithError(reader)
	if err != nil {
		testlib.WriteAnswer(err.Error(), testlib.PE)
		os.Exit(0)
	}
	return result
}

func ReaderIsEmpty(reader *bufio.Reader) bool {
	_, err := reader.Peek(1)
	return err != nil
}
