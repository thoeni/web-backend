package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDB() *sql.DB {
	//	Read ENV variable
	var mariadb_pwd = os.Getenv("MARIADB_PWD")
	var mariadb_container = os.Getenv("MARIADB_CONTAINER")
	db, err := sql.Open("mysql", "root:"+mariadb_pwd+"@tcp("+mariadb_container+":3306)/DBWEB")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
