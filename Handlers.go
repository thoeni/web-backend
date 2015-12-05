package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Golang API!")
}

func PageResource(w http.ResponseWriter, r *http.Request) {

	var db *sql.DB = ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	surveyId := vars["pageId"]

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM page WHERE id = ?")
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var response Page
	err = stmtOut.QueryRow(surveyId).Scan(&response.Id, &response.Summary, &response.Content)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}