package Query

import (
	"database/sql"
	"ginexample/db/mysql"
	"log"
)

func RunQuery(query string) (*sql.Rows, error) {
	rows, err := mysql.MySQL.Query(query)
	if err != nil {
		log.Println("Query failed, err: ", err)
		return nil, err
	}

	return rows, nil
}

func RunQueryCount(query string) (int, error) {
	var count int
	err := mysql.MySQL.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Query failed, err: ", err)
		return 0, err
	}

	return count, nil
}
