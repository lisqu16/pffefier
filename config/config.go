package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     	string
	Prefix    	string
	DBuser 			string
	DBpassword 	string
	DBname			string

	config *configStruct
)

type configStruct struct {
	Token     	string `json:"token"`
	Prefix    	string `json:"prefix"`
	DBuser 			string `json:"dbuser"`
	DBpassword 	string `json:"dbpass"`
	DBname			string `json:"dbname"`
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
	DBuser = config.DBuser
	DBpassword = config.DBpassword
	DBname = config.DBname

	return nil
}
