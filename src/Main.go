package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const GOPORT string = "8080"

func main() {
	router := NewRouter()
	fmt.Println("Ready, listening on", GOPORT, "...")
	log.Fatal(http.ListenAndServe(":"+GOPORT, cors.Default().Handler(router)))
}
