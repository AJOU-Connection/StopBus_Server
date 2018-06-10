package main

import "log"

// main is the main function.
func main() {
	go fakeBus.Run()

	server := Server{}
	err := server.Run(51234)
	if err != nil {
		log.Fatalf("main error: %v", err)
	}

}
