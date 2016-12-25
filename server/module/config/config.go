package config

import (
	"os"
	"encoding/json"
)

type DBSettings struct {
	User string
	Password string
	DbName string
}

type Configuration struct {
	NameApp string
	DbSettings DBSettings
}

func Get() *Configuration{
	fileConf, errConf := os.Open("conf.json")
	defer fileConf.Close()

	if errConf != nil {
		panic(errConf)
	}

	decoder := json.NewDecoder(fileConf)
	conf := new(Configuration)
	errDecode := decoder.Decode(conf)

	if errDecode != nil {
		panic(errDecode)
	}

	return conf
}