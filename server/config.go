package server

import (
	"encoding/json"
	"fmt"
	"go-web-example/server/log"
	"io/ioutil"
)

type Config struct {
	mode string
}

var config Config

func InitConfig(mode string) {
	filename := fmt.Sprintf("./resources/config-%s.json", mode)
	configJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalfln("Failed to read %s. Exit Program. %v", filename, err)
	}

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		log.Fatalfln("Failed to %v", err)
	}
}

func GetMode() string {
	return config.mode
}
