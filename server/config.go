package server

import (
	"encoding/json"
	"fmt"
	"go-web-example/server/log"
	"io/ioutil"
)

type Config struct {
	Mode        string
	StaticDir   string
	TemplateDir string
	MessageDir  string
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
	return config.Mode
}

func GetStaticDir() string {
	return config.StaticDir
}

func GetTemplateDir() string {
	return config.TemplateDir
}

func GetMessageDir() string {
	return config.MessageDir
}
