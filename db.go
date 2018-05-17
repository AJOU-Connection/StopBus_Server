package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func addUserToken(user User) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("INSERT INTO User VALUES (?, ?, NOW())", user.Token, user.UUID)
	if err != nil { // error exists
		return err
	}
	return nil
}
