package main

type Project struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Tech string `json:"tech"`
}

type Projects []Project
