package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func SetupDatabase() {
	var err error
	dbConn, err = sql.Open("mysql", "LksmA9Hp9h:IpXiDGbkwT@tcp(remotemysql.com:3306)/LksmA9Hp9h?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *sql.DB {
	return dbConn
}
