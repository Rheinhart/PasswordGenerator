package main

import (
	//"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"utils"
	//"log"
	//"net/http"
	//"strings"
	//"text/template"
	//"strings"
	"time"
	//"bytes"
	"fmt"
	"pwdGenerator"
)

const maxPwds = 20

func main() {

	defer utils.Track(time.Now())
	fmt.Print("sss")
	pwdGen := pwdGenerator.NewPwdGenerator()
	pwdGenerator.GenerateManyPasswords(*pwdGen, maxPwds)
	//for _ ,v:=range list {
	//	fmt.Println(v)
	//}

}
