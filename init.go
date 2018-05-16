package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

// configuration is a structure that specifies the contents of config.json.
type configuration struct {
	ServiceKey string   `json:"serviceKey"`
	Database   Database `json:"database"`
}

type Database struct {
	User    string `json:"user"`
	Passwd  string `json:"passwd"`
	IP_addr string `json:"ip_addr"`
	Port    string `json:"port"`
	DBname  string `json:"dbname"`
}

// config is a variable that stores configuration information.
var config configuration

// init is an initialization function.
func init() {
	err := setUpConfig()
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
}

func setUpConfig() error {
	if _, err := os.Stat("./configs"); err != nil {
		return errors.New("directory not exists: ./configs")
	}

	file, err := ioutil.ReadFile("./configs/config.json") // read config.json
	if err != nil {
		return errors.New("config file not exists: ./configs/config.json")
	}

	err = json.Unmarshal(file, &config) // store loaded json at config variable
	if err != nil {
		return errors.New("invalid JSON file: ./configs/config.json")
	}
	return nil
}
