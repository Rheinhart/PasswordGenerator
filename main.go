package main

import (
	"handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/passwords", handler.PasswordHandler)
	http.HandleFunc("/passwords/mapping", handler.VowelMappingHandler)

	log.Println("Passwor Generator Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
