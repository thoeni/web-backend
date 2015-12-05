package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
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
	err = stmtOut.QueryRow(surveyId).Scan(&response.Id, &response.Title, &response.Summary, &response.Content)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ContactsResource(w http.ResponseWriter, r *http.Request) {

	var db *sql.DB = ConnectDB()
	defer db.Close()

	var queryParamsMap map[string][]string = r.URL.Query()
	var contactType = queryParamsMap["type"]
	var rows = QueryContact(db, contactType)
	var response Contacts

	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.Id, &contact.Type, &contact.Contact)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, contact)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func QueryContact(db *sql.DB, contactType []string) *sql.Rows {
	if (contactType != nil) {
		// Prepare statement for reading data
		stmtOut, err := db.Prepare("SELECT * FROM contact where type = ?")
		if err != nil {
			panic(err.Error()) // TODO proper error handling instead of panic in your app
		}
		defer stmtOut.Close()
		//	Execute query
		rows, err := stmtOut.Query(contactType[0])
		if err != nil {
			panic(err.Error()) // TODO proper error handling instead of panic in your app
		}
		return rows
	} else {
		//	Execute query
		rows, err := db.Query("SELECT * FROM contact")
		if err != nil {
			log.Fatal(err)
		}
		return rows
	}
}