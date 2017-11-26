package main

import (
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"utils"
	//"log"
	//"net/http"
	//"strings"
	//"text/template"
	//"strings"
	//"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"pwdGenerator"
)

const maxPwds = 20

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/passwords", Index)
	log.Println("Random Passwor Generator starting...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Print("sss")
	pwdGen := pwdGenerator.NewPwdGenerator()
	pwdGenerator.GenerateManyPasswords(*pwdGen, maxPwds)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
