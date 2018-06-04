package main

import (
	"fmt"
	"os"
	"time"
)

const logDirPath = "logs"

func GetLogFile() *os.File {
	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		os.Mkdir(logDirPath, os.ModeDir)
	}

	f, err := os.OpenFile(fmt.Sprintf("%v/access.log", logDirPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func ErrorLogger(err error) {
	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		os.Mkdir(logDirPath, os.ModeDir)
	}

	f, err := os.OpenFile(fmt.Sprintf("%v/error.log", logDirPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	f.WriteString(fmt.Sprintf("%v %v\n", getCurrentTime(), err.Error()))

	return
}

func getCurrentTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", t.Year()%100, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
