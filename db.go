package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type GetType int

const (
	GetIn  GetType = 1
	GetOff GetType = 2
)

func addUserToken(user User) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("INSERT INTO User VALUES (?,?,NOW()) ON DUPLICATE KEY UPDATE token=?, UUID=?, registration_date=NOW()", user.Token, user.UUID, user.Token, user.UUID)
	if err != nil { // error exists
		return err
	}
	return nil
}

func addDriverStop(stopInput StopInput, getType GetType) (bool, error) {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return false, err
	}
	defer mysql.Close()

	var gt string
	var getInfo GetInfo
	if getType == GetIn {
		gt = "getIn"
		getInfo.IsGetIn = true
	} else {
		gt = "getOut"
		getInfo.IsGetOff = true
	}

	rows, err := mysql.Exec("INSERT INTO DriverStop VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE "+gt+"=?", stopInput.RouteID, stopInput.StationID, getInfo.IsGetIn, getInfo.IsGetOff, true)
	if err != nil { // error exists
		return false, err
	}

	n, err := rows.RowsAffected()
	if err != nil { // error exists
		return false, err
	}
	if n == 1 {
		return true, nil
	}

	return false, nil
}

func addGetIn(r Reserv) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("INSERT INTO GetIn VALUES (?, ?, ?)", r.UUID, r.RouteID, r.StationID)
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

func deleteDriverStop(routeID string, stationID string) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("DELETE FROM DriverStop WHERE routeID = ? AND stationID = ?", routeID, stationID)
	if err != nil { // error exists
		return err
	}
	return nil
}

func getTokenFromUUID(UUID string) (string, error) {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return "", err
	}
	defer mysql.Close()

	var token string
	err = mysql.QueryRow("SELECT token FROM User WHERE UUID=?", UUID).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func addStaDirect(stationID string, direct string) error {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return err
	}
	defer mysql.Close()

	_, err = mysql.Exec("INSERT INTO StationDirect VALUES (?,?) ON DUPLICATE KEY UPDATE direct=?", stationID, direct, direct)
	if err != nil { // error exists
		return err
	}
	return nil
}

func getStaDirect(stationID string) string {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return ""
	}
	defer mysql.Close()

	var direct string
	err = mysql.QueryRow("SELECT direct FROM StationDirect WHERE stationID=?", stationID).Scan(&direct)
	if err != nil {
		return ""
	}

	return direct
}

func getGetInUserTokens(routeID string, stationID string) ([]string, error) {
	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return nil, err
	}
	defer mysql.Close()

	tokens := []string{}
	var tempToken string
	rows, err := mysql.Query("SELECT token FROM User INNER JOIN GetIn ON User.UUID = GetIn.UUID AND (routeID = ? AND stationID=?)", routeID, stationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tempToken)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tempToken)
	}

	return tokens, nil
}

func getGetCount(routeID string, stationID string) (GetInfo, error) {
	var getInfo GetInfo

	mysql, err := sql.Open("mysql", config.Database.User+":"+config.Database.Passwd+"@tcp("+config.Database.IP_addr+":"+config.Database.Port+")/"+config.Database.DBname)
	if err != nil { // error exists
		return getInfo, err
	}
	defer mysql.Close()

	err = mysql.QueryRow("SELECT getIn, getOut FROM DriverStop WHERE routeID=? AND stationID=?", routeID, stationID).Scan(&getInfo.IsGetIn, &getInfo.IsGetOff)
	if err != nil {
		return getInfo, err
	}

	return getInfo, nil
}
