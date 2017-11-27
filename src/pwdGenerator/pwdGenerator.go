package pwdGenerator

import (
	"bytes"
	"log"
	"utils"
)

//using self defined type
type charBytes []byte

type CharBytes interface {
	RandomGetChars(num int) (out charBytes)
}

var (
	mixLetter      = charBytes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	mixedVowel     = charBytes("AEIOUaeiou")
	mixedConsonant = charBytes("BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz")
	digit          = charBytes("0123456789")
	specChar       = charBytes("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
)

func (charSet charBytes) RandomGetChars(num int) (out charBytes) {

	var buff bytes.Buffer
	for i := 0; i < num; i++ {
		index := utils.Random(0, len(charSet))
		buff.WriteByte(charSet[index])
		out = buff.Bytes()
	}

	return
}

type PwdQuery struct {
	MaxPwds    int `json: "limit"`
	MinLength  int `json: "length"`
	NumDigit   int `json: "digits"`
	NumSpecial int `json: "specials"`
}

type PwdPair struct {
	MappingRules map[byte]byte `json: "rules"`
	OldPassword  string        `json: "oldPassword"`
	NewPassword  string        `json: "nldPassword"`
}

type passwordGenerator struct {
	PwdQuery  `json: "query"`
	Passwords []string `json: "passwords"`
}

type PwdGenerable interface {
	Generate() string
	Map(mapping map[byte]byte, old string) string
	CreateVowelDigitRandomMapping() map[byte]byte
	GetPasswords() []string
}

func (p *passwordGenerator) GetPasswords() []string {
	return p.Passwords
}

//pwd generator factory
func NewPwdGenerator() *passwordGenerator {
	return new(passwordGenerator)
}

func (p *passwordGenerator) CreateVowelDigitRandomMapping() map[byte]byte {

	var mapping = make(map[byte]byte)
	rDigits := utils.Shuffle(digit)
	rVowels := utils.Shuffle(mixedVowel)
	for i, v := range rVowels {
		mapping[v] = rDigits[i]
	}

	//fmt.Printf("%c\n", mapping)
	return mapping
}

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

func (p *passwordGenerator) Generate(q PwdQuery) (pwd string) {

	p.PwdQuery = q
	distance := 5

	length := utils.Random(p.MinLength, p.MinLength+distance)
	restLegnth := length - p.NumSpecial - p.NumDigit

	var letters, digits, specials charBytes
	var allChars bytes.Buffer

	letters = mixLetter.RandomGetChars(restLegnth)
	if _, err := allChars.Write(letters); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	digits = digit.RandomGetChars(p.NumDigit)
	if _, err := allChars.Write(digits); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	specials = specChar.RandomGetChars(p.NumSpecial)
	if _, err := allChars.Write(specials); err != nil {
		log.Println("Buffer Writing Error: %s\n", err)
	}

	pwd = string(utils.Shuffle(allChars.Bytes()))

	return
}

func (p *passwordGenerator) GenerateManyPasswords(q PwdQuery) {

	var pwd string

	for i := 0; i < q.MaxPwds; i++ {
		pwd = p.Generate(q)
		p.Passwords = append(p.Passwords, pwd)
	}

}
