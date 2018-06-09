package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Run(port int) error {
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), Handler())
	if err != nil {
		log.Printf("StopBus Server Failed... :%v\n", port)
		return err
	}

	log.Printf("StopBus Server Run! :%v\n", port)
	return nil
}
