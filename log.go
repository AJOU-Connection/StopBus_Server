package main

import (
	"fmt"
	"os"
)

func GetLogFile() *os.File {
	f, err := os.OpenFile(fmt.Sprintf("access_.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
