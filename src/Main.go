package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/rs/cors"
)

const GOPORT string = "8080"

func main() {
	router := NewRouter()
	fmt.Println("Ready, listening on", GOPORT,"...")
	log.Fatal(http.ListenAndServe(":"+GOPORT, cors.Default().Handler(router)))
}
