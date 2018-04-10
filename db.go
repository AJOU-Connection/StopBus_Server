package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DBType    string
	User      string
	Password  string
	IPAddress string
	Port      string
	Name      string
}

type RouteSchema struct {
	ID            string
	Number        string
	Type          string
	UpFirstTime   string
	UpLastTime    string
	DownFirstTime string
	DownLastTime  string
}

func (db Database) Query(query string) (routes []RouteSchema) {
	// sql.DB 객체 생성
	mysql, err := sql.Open(db.DBType, db.User+":"+db.Password+"@tcp("+db.IPAddress+":"+db.Port+")/"+db.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	rows, err := mysql.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	routes = make([]RouteSchema, 10, 100)
	index := 0
	for rows.Next() {
		err := rows.Scan(&routes[index].ID,
			&routes[index].Number,
			&routes[index].Type,
			&routes[index].UpFirstTime,
			&routes[index].UpLastTime,
			&routes[index].DownFirstTime,
			&routes[index].DownLastTime)
		if err != nil {
			log.Fatal(err)
		}
		index++
	}
	return
}
