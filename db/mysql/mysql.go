package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MySQL *sql.DB

func Init() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DB"))
	if err != nil {
		log.Panic("Connect MySQL failed, err: ", err)
	}

	fmt.Println("Connect MySQL success")

	MySQL = db
}
