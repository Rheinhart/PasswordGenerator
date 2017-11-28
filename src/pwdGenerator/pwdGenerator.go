// Package pwdGenerator password generator package
package pwdGenerator

import (
	"bytes"
	"errors"
	"log"
	"utils"
)

// charBytes using self defined type
type charBytes []byte

type CharBytes interface {
	RandomGetChars(num int) (out charBytes)
}

//Predefined Character Set
var (
	mixLetter      = charBytes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	mixedVowel     = charBytes("AEIOUaeiou")
	mixedConsonant = charBytes("BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz")
	digit          = charBytes("0123456789")
	specChar       = charBytes("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
)

// RandomGetChars get random characters from charBytes
func (charSet charBytes) RandomGetChars(num int) (out charBytes) {

	var buff bytes.Buffer
	for i := 0; i < num; i++ {
		index := utils.Random(0, len(charSet))
		buff.WriteByte(charSet[index])
		out = buff.Bytes()
	}

	return
}

// PwdQuery get parameters from client
type PwdQuery struct {
	MaxPwds    int `json:"limit"`
	MinLength  int `json:"length"`
	NumDigit   int `json:"digits"`
	NumSpecial int `json:"specials"`
}

// JsStatus json error message
type JsStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// PwdPair post method data from client
type PwdPair struct {
	JsStatus     `json:"status"`
	MappingRules map[byte]byte `json:"-"` // random vowel to digit
	OldPassword  string        `json:"oldPassword"`
	NewPassword  string        `json:"newPassword"`
}

// passwordGenerator to generate one or many random passwords
type passwordGenerator struct {
	JsStatus  `json:"status"`
	Query     PwdQuery `json:"query"`
	Passwords []string `json:"passwords"`
}

// PwdGenerable interface
type PwdGenerable interface {
	Generate() string
	Map(mapping map[byte]byte, old string) string
	CreateVowelDigitRandomMapping() map[byte]byte
}

//NewPwdGenerator passwordGenerator factory
func NewPwdGenerator() *passwordGenerator {
	return new(passwordGenerator)
}

// CreateVowelDigitRandomMapping random generate mapping rules from vowel to digit
func (p *passwordGenerator) CreateVowelDigitRandomMapping() map[byte]byte {

	var mapping = make(map[byte]byte)
	rDigits := utils.Shuffle(digit)
	rVowels := utils.Shuffle(mixedVowel)
	for i, v := range rVowels {
		mapping[v] = rDigits[i]
	}

	return mapping
}

//Map mapping vowel to digit according the rules
func (p *passwordGenerator) Map(rules map[byte]byte, old string) string {

	var buff bytes.Buffer
	for i := 0; i < len(old); i++ {
		if v, ok := rules[byte(old[i])]; ok {
			buff.WriteByte(v)
		} else {
			buff.WriteByte(old[i])
		}
	}

	return buff.String()
}

// Generate generate one password
func (p *passwordGenerator) Generate() (pwd string) {

	distance := 5 //maximum length = minmum length + distance

	length := utils.Random(p.Query.MinLength, p.Query.MinLength+distance) // random get the password length bwtween min and max
	rest := length - p.Query.NumSpecial - p.Query.NumDigit

	var letters, digits, specials charBytes
	var allChars bytes.Buffer

	letters = mixLetter.RandomGetChars(rest)
	if _, err := allChars.Write(letters); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	digits = digit.RandomGetChars(p.Query.NumDigit)
	if _, err := allChars.Write(digits); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	specials = specChar.RandomGetChars(p.Query.NumSpecial)
	if _, err := allChars.Write(specials); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	pwd = string(utils.Shuffle(allChars.Bytes())) // shuffle all characters in the password

	return
}

// GenerateManyPasswords generate many passwords
func (p *passwordGenerator) GenerateManyPasswords() {

	var pwd string

	for i := 0; i < p.Query.MaxPwds; i++ {
		pwd = p.Generate()
		p.Passwords = append(p.Passwords, pwd)
	}

}

// CheckErrors check query errors
func (q PwdQuery) CheckErrors() (bool, error) {

	if q.MinLength <= 0 {
		err := errors.New("Input Error: length can not smaller than 0!")
		return false, err
	}
	if q.NumSpecial < 0 {
		err := errors.New("Input Error: special number should not smaller than 0!")
		return false, err
	}
	if q.NumDigit < 0 {
		err := errors.New("Input Error: digit number should not smaller than 0!")
		return false, err
	}
	if q.NumDigit > q.MinLength {
		err := errors.New("Input Error: digit number should not larger than length!")
		return false, err
	}
	if q.NumSpecial > q.MinLength {
		err := errors.New("Input Error: special number should not larger than length!")
		return false, err
	}

	if q.MinLength-q.NumSpecial-q.NumDigit < 0 {
		err := errors.New("Input Error:  special number plus digit number should not larger than length!")
		return false, err

	} else {
		return true, nil
	}
}
