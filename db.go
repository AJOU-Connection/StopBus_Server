package StopBus

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database 구조체는 데이터베이스와 관련된 함수를 가질 구조체이다.
type Database struct{}

// RouteSchema 구조체는 Route 테이블의 Column을 담은 구조체이다.
type RouteSchema struct {
	ID            string
	Number        string
	Type          string
	UpFirstTime   string
	UpLastTime    string
	DownFirstTime string
	DownLastTime  string
}

// StationSchema 구조체는 Station 테이블의 Column을 담은 구조체이다.
type StationSchema struct {
	ID        string
	number    string
	name      string
	direction string
}

func configCheck() (ret bool) {
	ret = true

	if (DBConfig{}) == config.DB {
		log.Println("[ERROR] config.DB not exists.")
		ret = false
	}

	return
}

// Query 함수는 데이터베이스로 쿼리를 보낸 결과를 반환하는 함수이다.
func (db Database) Query(query string) (result []RouteSchema, queryError error) {
	if configCheck() == false {
		queryError = errors.New("Query() error")
		return
	}

	// sql.DB 객체 생성
	mysql, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.IPAddress+":"+config.DB.Port+")/"+config.DB.Table)

	if err != nil { // error exists
		log.Printf("[ERROR] %v\n", err)
		queryError = errors.New("Query() error")
		return
	}
	defer mysql.Close()

	rows, err := mysql.Query(query)

	if err != nil { // error exists
		log.Printf("[ERROR] %v\n", err)
		queryError = errors.New("Query() error")
		return
	}
	defer rows.Close()
	var rowData RouteSchema
	for rows.Next() {
		err := rows.Scan(&rowData.ID,
			&rowData.Number,
			&rowData.Type,
			&rowData.UpFirstTime,
			&rowData.UpLastTime,
			&rowData.DownFirstTime,
			&rowData.DownLastTime)
		if err != nil {
			log.Printf("[ERROR] %v\n", err)
			queryError = errors.New("Query() error")
			return
		}
		result = append(result, rowData)
	}
	return
}