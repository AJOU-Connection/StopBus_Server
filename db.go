package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Token string
	UUID  string
}

func addUserToken(user User) int64 {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		log.Println(err)
		return -1
	}
	defer mysql.Close()

	r, err := mysql.Exec("INSERT INTO User VALUES (?, ?, NOW())", user.Token, user.UUID)
	if err != nil { // error exists
		log.Println(err)
		return -1
	}
	n, _ := r.RowsAffected()

	return n
}
