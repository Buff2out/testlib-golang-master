package testlib

import (
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Metadata string `json:"metadata"`
	Input    string `json:"input"`
	Output   string `json:"output"`
}

func (c *Configuration) InitFromFile() {
	cfg, err := ini.Load("checker.conf")
	if err != nil {
		WriteAnswer(fmt.Sprintf("Failed to parse checker config. Err: %s", err.Error()), CF)
		os.Exit(1)
	}
	c.Metadata = cfg.Section("DEFAULT").Key("metadata").String()
	c.Input = cfg.Section("DEFAULT").Key("input").String()
	c.Output = cfg.Section("DEFAULT").Key("output").String()
}
func (c *Configuration) InitFromPipe() error {
	stream, err := GetStreamWithError(CheckConfig)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	_ = stream.Close()
	return nil
}
