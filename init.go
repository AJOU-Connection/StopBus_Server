package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// configuration is a structure that specifies the contents of config.json.
type configuration struct {
	ServiceKey string   `json:"serviceKey"`
	ServerKey  string   `json:"serverKey"`
	Database   Database `json:"database"`
}

type Database struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	IPAddr string `json:"ip_addr"`
	Port   string `json:"port"`
	DBname string `json:"dbname"`
}

// config is a variable that stores configuration information.
var config configuration

// init is an initialization function.
func init() {
	setUpConfig()
	setUpUsingEnv()
}

func setUpConfig() {
	if _, err := os.Stat("./configs"); err != nil {
		return
	}

	file, err := ioutil.ReadFile("./configs/config.json") // read config.json
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config) // store loaded json at config variable
	if err != nil {
		return
	}
}

func setUpUsingEnv() {
	if config.ServiceKey == "" {
		config.ServiceKey = os.Getenv("serviceKey")
	}

	if config.ServerKey == "" {
		config.ServerKey = os.Getenv("serverKey")
	}

	if config.Database == (Database{}) {
		config.Database = Database{
			os.Getenv("databaseUser"),
			os.Getenv("databasePasswd"),
			os.Getenv("databaseIpAddr"),
			os.Getenv("databasePort"),
			os.Getenv("databaseDbname"),
		}
	}
}
