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

	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	pageId := vars["pageId"]

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM page WHERE id = ?")
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var response Page
	err = stmtOut.QueryRow(pageId).Scan(&response.Id, &response.Title, &response.Summary, &response.Content)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ContactsResource(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	queryParamsMap := r.URL.Query()
	contactType := queryParamsMap["type"]
	rows := QueryContact(db, contactType)
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
	}
	//	Execute query
	rows, err := db.Query("SELECT * FROM contact")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func AllTestimonialResource(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	//	Execute query
	rows, err := db.Query("SELECT * FROM testimonial")
	if err != nil {
		log.Fatal(err)
	}
	var response Testimonials

	for rows.Next() {
		var testimonial Testimonial
		err := rows.Scan(&testimonial.Name, &testimonial.Text, &testimonial.Date)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, testimonial)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func TestimonialResource(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var response Testimonials

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM testimonial where name = ?")
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var testimonial Testimonial
	//	Execute query
	err = stmtOut.QueryRow(name).Scan(&testimonial.Name, &testimonial.Text, &testimonial.Date)
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}

	response = append(response, testimonial)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func AllProjectResource(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	//	Execute query
	rows, err := db.Query("SELECT * FROM project")
	if err != nil {
		log.Fatal(err)
	}
	var response Projects

	for rows.Next() {
		var project Project
		err := rows.Scan(&project.Name, &project.Text, &project.Tech)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, project)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ProjectResource(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var response Projects

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM project where name = ?")
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var project Project
	//	Execute query
	err = stmtOut.QueryRow(name).Scan(&project.Name, &project.Text, &project.Tech)
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}

	response = append(response, project)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}