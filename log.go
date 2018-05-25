package main

import (
	"fmt"
	"os"
)

func GetLogFile() *os.File {
	logDirPath := "logs"

	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		os.Mkdir(logDirPath, os.ModeDir)
	}

	f, err := os.OpenFile(fmt.Sprintf("%v/access_.log", logDirPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
