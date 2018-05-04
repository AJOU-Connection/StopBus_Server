package StopBus

import (
	"log"
	"net/http"
)

// main is the main function.
func main() {
	err := setUpConfig()
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	err = http.ListenAndServe(":51234", Handler())
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
}
