package main

import (
	"log"
	"net/http"
	"fmt"
	"net"
	"os"
)

const GOPORT string = "8080"

func main() {
	router := NewRouter()
	fmt.Println("Ready, listening on", GOPORT,"...")
	var mariadb_container = os.Getenv("MARIADB_CONTAINER");
	containerIp, err := net.LookupIP(mariadb_container)
	if err != nil {
		panic(err.Error()) // TODO proper error handling instead of panic in your app
	}
	fmt.Println("Database container IP is:", containerIp)
	log.Fatal(http.ListenAndServe(":"+GOPORT, router))
}
