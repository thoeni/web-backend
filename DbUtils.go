package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"fmt"
)


func ConnectDB() *sql.DB  {
	//	Read ENV variable
	var mariadb_pwd = os.Getenv("MARIADB_PWD");
	var mariadb_container = os.Getenv("MARIADB_CONTAINER");
	fmt.Println("Connecting to -> root:"+mariadb_pwd+"@tcp("+mariadb_container+":3306)/DBWEB")
	db, err := sql.Open("mysql", "root:"+mariadb_pwd+"@tcp("+mariadb_container+":3306)/DBWEB")
	fmt.Println("Connected to: tcp("+mariadb_container+":3306)/DBWEB")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
