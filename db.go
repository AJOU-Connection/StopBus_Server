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

func addGetIn(r Reserv) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("INSERT INTO GetIn VALUES (?, ?, ?)", r.UserToken, r.RouteID, r.StationID)
	if err != nil { // error exists
		return err
	}
	return nil
}

func deleteGetIn(routeID string, stationID string) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("DELETE FROM GetIn WHERE routeID = ? AND stationID = ?", routeID, stationID)
	if err != nil { // error exists
		return err
	}
	return nil
}

func getGetInCount(routeID string, stationID string) (int, error) {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return -1, err
	}
	defer mysql.Close()

	var getInCount int
	err = mysql.QueryRow("SELECT COUNT(*) FROM GetIn WHERE routeID=? AND stationID=?", routeID, stationID).Scan(&getInCount)
	if err != nil {
		return -1, err
	}

	return getInCount, nil
}

func getGetOutCount(routeID string, stationID string) (int, error) {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return -1, err
	}
	defer mysql.Close()

	var getOutCount int
	err = mysql.QueryRow("SELECT COUNT(*) FROM GetOut WHERE routeID=? AND stationID=?", routeID, stationID).Scan(&getOutCount)
	if err != nil {
		return -1, err
	}

	return getOutCount, nil
}
