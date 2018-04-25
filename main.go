package StopBus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Configs 구초체는 서버 설정 정보에 대한 구조체이다.
type Configs struct {
	ServiceKey string   `json:"serviceKey"`
	DB         DBConfig `json:"db"`
}

// DBConfig 구조체는 데이터베이스 설정 정보에 대한 구조체이다.
type DBConfig struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	IPAddress string `json:"ipAddress"`
	Port      string `json:"port"`
	Table     string `json:"table"`
}

// config 변수는 전역변수이다.
var config Configs

// init 함수는 서버가 구동되기 전 초기화 함수이다.
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

// main 함수는 서버의 구동함수이다.
func main() {
	router := httprouter.New() // create router
	router.GET("/", Index)     // GET Root

	log.Fatal(http.ListenAndServe(":51234", router)) // 51234
}