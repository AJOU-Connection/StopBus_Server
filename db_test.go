package main

import (
	"testing"
)

func TestQuery(t *testing.T) {
	db := Database{"mysql", {config.DB.User, config.DB.Password, config.DB.IPAddress, config.DB.Port, config.DB.Name}}
	db.Query("SELECT * FROM Route")
	//fmt.Println(db.Query("SELECT * FROM Route"))
}
