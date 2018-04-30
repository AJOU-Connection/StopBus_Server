package StopBus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// configuration is a structure that specifies the contents of config.json.
type configuration struct {
	ServiceKey string `json:"serviceKey"`
}

// config is a variable that stores configuration information.
var config configuration

//init is an initialization function.
func init() {
	file, err := ioutil.ReadFile("./configs/config.json") // read config.json
	if err != nil {                                       // error exists
		log.Printf("[ERROR] %v\n", err)
		return
	}

	err = json.Unmarshal(file, &config) // store loaded json at config variable
	if err != nil {                     // error exists
		log.Printf("[ERROR] %v\n", err)
		return
	}
}

// main is the main function.
func main() {
	router := httprouter.New() // create router
	router.GET("/", Index)     // GET Root

	log.Fatal(http.ListenAndServe(":51234", router)) // 51234
}
