package main

type Page struct {
	Id			string `json:"id"`
	Title		string `json:"title"`
	Summary 	string `json:"summary"`
	Content  	string `json:"content"`
}

type Pages []Page
