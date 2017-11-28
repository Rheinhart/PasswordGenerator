package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"pwdGenerator"
	"strconv"
)

var templates *template.Template

const (
	FORMAT_INVALID = 490 //TODO
	INPUT_ERROR    = 491
)

func init() {

	if templates == nil {
		templates = template.Must(template.ParseGlob("src/templates/*.tmpl"))
	}
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

	index := templates.Lookup("index")
	index.ExecuteTemplate(w, "index", nil)

}

/**
 * @api {get} /passwords Generate passwords
 * @apiName GetPasswords
 * @apiGroup Passwords

 * @apiDescription API to get random generated password

 * @apiParam {number} length minimun length of the password
 * @apiParam {number} specials number of special characters of the password
 * @apiParam {number} digits number of digits of the password
 * @apiParam {number} limits how many passwords to be generated

 * @apiExample Example usage
 * curl -i "http://localhost:8080/passwords?length=8&specials=2&digits=2&limits=10"

 * @apiSuccess 200 ok
 * @apiSuccessExample {json} Success-Response:
 * {
 * "status": {
 *   "code": 200,
 *   "message": ""
 *  },
 *  "query": {
 *   "limit": 15,
 *   "length": 8,
 *   "digits": 2,
 *   "specials": 2
 *  },
 *  "passwords":
 *  [
 *   "zeF$;PwC8tK0",
 *   "o0?MMX'9",
 *   "'Kqfl6v5XY)",
 *   "7yzKL\u003e7HC+",
 *   "?NTt0FF5qv\"",
 *   "pdcw$0Y9/",
 *   "uy8Rcr0*o{zZ",
 *   "zp6J:y5w#jJ",
 *   "q0\u003cu)GBM5",
 *  "BOxBxn+X\u003e02",
 *  ]
 * }

 * @apiError 491 input error
 * @apiErrorExample {json} Error-Response:
 * {
 * "code": 491,
 * "message": "Input Error: digit number should not larger than length!"
 * }
 */
func PasswordHandler(w http.ResponseWriter, req *http.Request) {

	query := pwdGenerator.PwdQuery{MaxPwds: 15}

	if req.Method == "GET" {

		if lengths, err := strconv.Atoi(req.URL.Query().Get("length")); err == nil {
			query.MinLength = lengths
		}

		if digits, err := strconv.Atoi(req.URL.Query().Get("digits")); err == nil {
			query.NumDigit = digits
		}

		if specials, err := strconv.Atoi(req.URL.Query().Get("specials")); err == nil {
			query.NumSpecial = specials
		}

		if limits, err := strconv.Atoi(req.URL.Query().Get("limits")); err == nil {
			query.MaxPwds = limits
		}

		g := templates.Lookup("generation")

		if ok, e := query.CheckErrors(); !ok {

			var jsError pwdGenerator.JsStatus
			jsError.Code = INPUT_ERROR
			jsError.Message = e.Error()
			// return json if necessary, here just rend template
			if jsData, err := json.MarshalIndent(jsError, "", " "); err != nil {
				panic(err)
			} else {
				log.Printf("%s\n", string(jsData))
				//w.Write(jsData)
			}
			g.ExecuteTemplate(w, "generation", nil)

		} else {

			generator := pwdGenerator.NewPwdGenerator()
			generator.Query = query
			generator.Code = 200
			generator.GenerateManyPasswords()
			g.ExecuteTemplate(w, "generation", generator)

			// return json if necessary, here just rend template
			if jsData, err := json.MarshalIndent(generator, "", " "); err != nil {
				panic(err)
			} else {
				log.Printf("%s\n", string(jsData))
				//w.Write(jsData)
			}

		}

		m := templates.Lookup("mapping")
		m.ExecuteTemplate(w, "mapping", nil)
	}

}

/**
 * @api {post} /passwords/mapping Convert vowels of password
 * @apiName ReplaceVowel
 * @apiGroup Passwords

 * @apiDescription API to randomly covert vowels to random numbers of the password

 * @apiParam {string} password password contains vowels to be converted


 * @apiExample Example usage
 * curl -i "http://localhost:8080/passwords/mapping?password=@$afGwqwe"

 * @apiSuccess 200 ok
 * @apiSuccessExample {json} Success-Response:
 * {
 * "status": {
 *  "code": 200,
 *  "message": ""
 * },
 *  "oldPassword": "@$afGwqwe",
 *  "newPassword": "@$8fGwqw2"
 * }
 */
func VowelMappingHandler(w http.ResponseWriter, req *http.Request) {

	pp := new(pwdGenerator.PwdPair)

	req.ParseForm()

	if req.Method == "POST" {

		w.Header().Set("Content-type", "text/html")

		pp.OldPassword = req.FormValue("password")

		generator := pwdGenerator.NewPwdGenerator()

		pp.MappingRules = generator.CreateVowelDigitRandomMapping()
		pp.NewPassword = generator.Map(pp.MappingRules, pp.OldPassword)
		pp.Code = 200

		log.Printf("Mapping Rules: %c\n", pp.MappingRules)

		if jsData, err := json.MarshalIndent(pp, "", " "); err != nil {
			panic(err)
		} else {
			log.Printf("%s\n", string(jsData))
			//w.Write(jsData)
		}

		g := templates.Lookup("generation")
		err := g.ExecuteTemplate(w, "generation", nil)
		if err != nil {
			log.Printf("%s\n", err)
		}
		m := templates.Lookup("mapping")
		m.ExecuteTemplate(w, "mapping", pp)
	}
}
