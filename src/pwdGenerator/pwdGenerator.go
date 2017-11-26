package pwdGenerator

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type passwordGenerator struct {
	name      int
	passwords []string
}

type PwdGenerable interface {
	Generate(min int) string
	Map(mapping map[byte]byte, old string) string
	CreateVowelDigitRandomMapping() map[byte]byte
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

	fmt.Printf("%c", mapping)
	return mapping
}

func (p *passwordGenerator) Map(mapping map[byte]byte, old string, on bool) string {

	if on {
		var buff bytes.Buffer
		for i := 0; i < len(old); i++ {
			if v, ok := mapping[byte(old[i])]; ok {
				buff.WriteByte(v)
			} else {
				buff.WriteByte(old[i])
			}
		}
		new := buff.String()
		return new
	} else {
		return old
	}
}

func (p *passwordGenerator) Generate(min int) (pwd string) {

	minLength := min
	distance := 5
	length := utils.Random(minLength, minLength+distance)
	numDigit := 2
	numSpec := 2
	numRest := length - numSpec - numDigit

	if numDigit > minLength || numSpec > minLength || minLength-numSpec-numDigit <= 0 {
		panic("error!")
	}

	var letters, digits, specials charBytes
	var allChars bytes.Buffer

	letters = mixLetter.RandomGetChars(numRest)
	if _, err := allChars.Write(letters); err != nil {
		fmt.Println("Buffer Writing Error: %s\n", err)
	}

	digits = digit.RandomGetChars(numDigit)
	if _, err := allChars.Write(digits); err != nil {
		fmt.Println("Buffer Writing Error: %s\n", err)
	}

	specials = specChar.RandomGetChars(numSpec)
	if _, err := allChars.Write(specials); err != nil {
		fmt.Println("Buffer Writing Error: %s\n", err)
	}

	pwd = string(utils.Shuffle(allChars.Bytes()))

	//fmt.Printf("%s\n", pwd)

	return
}

func GenerateManyPasswords(gen passwordGenerator, num int) {

	var pwd string
	mapping := gen.CreateVowelDigitRandomMapping()

	for i := 0; i < num; i++ {
		pwd = gen.Generate(8)
		fmt.Printf("%s\n", pwd)

		pwd = gen.Map(mapping, pwd, true)
		fmt.Printf("%s\n", pwd)

		gen.passwords = append(gen.passwords, pwd)
	}

	jsList, _ := json.MarshalIndent(gen.passwords, "", " ")

	//fmt.Printf("%s\n",string(jsList))

	var f interface{}
	json.Unmarshal(jsList, &f)

}
