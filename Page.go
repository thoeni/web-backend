package main

type Page struct {
	Id			string `json:"id"`
	Summary 	string `json:"summary"`
	Content  	string `json:"content"`
}

type Pages []Page
