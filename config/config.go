package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	Prefix    string

	// Private variables
	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	Prefix    string `json:"prefix"`
}

func Read() error {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	Prefix = config.Prefix

	return nil
}
