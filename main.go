package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Configs struct {
	ServiceKey string `json:"serviceKey"`
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	User string `json:"user"`
	Password string `json:"password"`
	IPAddress string `json:"ipAddress"`
	Port string `json:"port"`
	Name string `json:"name"`
}

var config Configs

func init() {
	file, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":51234", router))
}
