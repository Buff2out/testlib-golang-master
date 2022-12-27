package testlib

import (
	"encoding/json"
	"fmt"
)

var Separators []byte
var Config Configuration

func init() {
	Separators = append(Separators, ' ', '\n', '\r', '\t')
	err := Config.InitFromPipe()
	if err != nil {
		Config.InitFromFile()
	}
}

type CheckerResponse struct {
	Verdict Verdict
	Message string
}

func WriteAnswer(msg string, verdict Verdict) {
	response := CheckerResponse{Message: msg, Verdict: verdict}
	pipe, err := GetStreamWithError(Answer)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(pipe).Encode(response)
	if err != nil {
		errorPipe, _ := GetStreamWithError(Error)
		_, _ = errorPipe.WriteString(fmt.Sprintf("Failed to write answer. Answer: %v. Error: %s", response, err.Error()))
	}
	_ = pipe.Close()
}
