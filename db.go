package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func addUserToken(token string) int64 {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		log.Println(err)
		return -1
	}
	defer mysql.Close()

	r, err := mysql.Exec("INSERT INTO User VALUES (?, NOW())", token)
	if err != nil { // error exists
		log.Println(err)
		return -1
	}
	n, _ := r.RowsAffected()

	return n
}
