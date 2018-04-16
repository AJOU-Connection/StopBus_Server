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

// Configs is server config
type Configs struct {
	ServiceKey string   `json:"serviceKey"`
	DB         DBConfig `json:"db"`
}

// DBConfig is database config
type DBConfig struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	IPAddress string `json:"ipAddress"`
	Port      string `json:"port"`
	Name      string `json:"name"`
}

// config is internal setting variable
var config Configs

// server init function - only once excution
func init() {
	file, err := ioutil.ReadFile("./configs/config.json") // read config.json
	if err != nil {                                       // error exists
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config) // store loaded json at config variable
	if err != nil {                     // error exists
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	router := httprouter.New() // create router
	router.GET("/", Index)     // GET Root

	log.Fatal(http.ListenAndServe(":51234", router)) // 51234
}
