package main

type Contact struct {
	Id			string `json:"id"`
	Type		string `json:"type"`
	Contact 	string `json:"contact"`
}

type Contacts []Contact
