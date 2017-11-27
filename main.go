package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pwdGenerator"
	"strconv"
)

var templates *template.Template

func init() {

	if templates == nil {
		templates = template.Must(template.ParseGlob("src/templates/*.tmpl"))
	}
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/passwords", PasswordHandler)
	http.HandleFunc("/passwords/mapping", VowelMappingHandler)

	log.Println("Random Passwor Generator starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

	index := templates.Lookup("index")
	index.ExecuteTemplate(w, "index", nil)

}

func PasswordHandler(w http.ResponseWriter, req *http.Request) {

	q := pwdGenerator.PwdQuery{MaxPwds: 15}
	if req.Method == "GET" {
		if lengths, err := strconv.Atoi(req.URL.Query().Get("length")); err == nil {
			q.MinLength = lengths
		}

		if digits, err := strconv.Atoi(req.URL.Query().Get("digits")); err == nil {
			q.NumDigit = digits
		}

		if specials, err := strconv.Atoi(req.URL.Query().Get("specials")); err == nil {
			q.NumSpecial = specials
		}

		if limits, err := strconv.Atoi(req.URL.Query().Get("limits")); err == nil {
			q.MaxPwds = limits
		}

		g := templates.Lookup("generation")

		if q.NumDigit < 0 || q.MinLength <= 0 || q.NumSpecial < 0 || q.NumDigit > q.MinLength || q.NumSpecial > q.MinLength || q.MinLength-q.NumSpecial-q.NumDigit < 0 {

			log.Printf("input rules error")
			g.ExecuteTemplate(w, "generation", nil)

		} else {

			generator := pwdGenerator.NewPwdGenerator()
			generator.GenerateManyPasswords(q)
			g.ExecuteTemplate(w, "generation", generator)

			jsList, _ := json.MarshalIndent(generator, "", " ")
			log.Printf("%s\n", string(jsList))

			// return json if necessary
			//if err := json.NewEncoder(w).Encode(generator.Passwords); err != nil {
			//	panic(err)
			//}

		}

		m := templates.Lookup("mapping")
		m.ExecuteTemplate(w, "mapping", nil)
	}

}

func VowelMappingHandler(w http.ResponseWriter, req *http.Request) {

	pp := new(pwdGenerator.PwdPair)

	req.ParseForm()

	if req.Method == "POST" {

		w.Header().Set("Content-type", "text/html")

		pp.OldPassword = req.FormValue("password")

		generator := pwdGenerator.NewPwdGenerator()

		pp.MappingRules = generator.CreateVowelDigitRandomMapping()
		pp.NewPassword = generator.Map(pp.MappingRules, pp.OldPassword)

		log.Printf("Mapping Rules: %c\n", pp.MappingRules)

		jsPair, _ := json.MarshalIndent(pp, "", " ")
		log.Printf("%s\n", string(jsPair))
		//if err := json.NewEncoder(w).Encode(pp); err != nil {
		//	panic(err)
		//	}

		g := templates.Lookup("generation")
		err := g.ExecuteTemplate(w, "generation", nil)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
		}
		m := templates.Lookup("mapping")
		m.ExecuteTemplate(w, "mapping", pp)
	}
}
