package db

import (
	"database/sql"
	"fmt"

	"example.com/m/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	fmt.Print(connectionString)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("ConnectionString error ..")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {

	return db

}
