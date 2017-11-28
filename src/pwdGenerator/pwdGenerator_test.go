package pwdGenerator

import (
	"reflect"
	"testing"
	"utils"
)

//test get random chars from charset
func TestCharBytes_RandomGetChars(t *testing.T) {
	many := 10

	out := mixedVowel.RandomGetChars(many)
	if len(out) != many {
		t.Log("get mixedVowel wrong numbers!")
		t.Fail()
	}

	out = specChar.RandomGetChars(many)
	if len(out) != many {
		t.Log("get special wrong numbers!")
		t.Fail()
	}

	out = digit.RandomGetChars(many)
	if len(out) != many {
		t.Log("get wrong digit numbers!!")
		t.Fail()
	}
}

//test create
func TestNewPwdGenerator(t *testing.T) {

	gen := NewPwdGenerator()
	q := &passwordGenerator{}
	v := reflect.TypeOf(gen)
	if v == reflect.TypeOf(q) {
		t.Log("Return right type")
	} else {
		t.Log("Return wrong type")
		t.Fail()
	}
}

func TestPwdQuery_CheckErrors(t *testing.T) {

	var query1 = PwdQuery{MinLength: 10, NumSpecial: 9, NumDigit: 3}
	var query2 = PwdQuery{MinLength: 10, NumSpecial: 2, NumDigit: 3}
	var query3 = PwdQuery{MinLength: -1, NumSpecial: 2, NumDigit: 3}

	var ok bool
	ok, _ = query1.CheckErrors()
	if ok == false {
		t.Log("query1 check pass")
	} else {
		t.Log("query1 check failed")
		t.Fail()
	}

	ok, _ = query2.CheckErrors()
	if ok == true {
		t.Log("query2 check pass")
	} else {
		t.Log("query2 check failed")
		t.Fail()
	}

	ok, _ = query3.CheckErrors()
	if ok == false {
		t.Log("query3 check pass")
	} else {
		t.Log("query3 check failed")
		t.Fail()
	}

}

func TestPasswordGenerator_Generate(t *testing.T) {

	var query = PwdQuery{MinLength: 10, NumSpecial: 2, NumDigit: 3}
	gen := NewPwdGenerator()
	if ok, err := query.CheckErrors(); ok {
		gen.Query = query
	} else {
		t.Log("query check failed!", err)
		t.Fail()
	}

	for i := 0; i < 1000; i++ {

		gen.Query.MinLength = utils.Random(13, 100)
		gen.Query.NumSpecial = utils.Random(1, 6)
		gen.Query.NumDigit = utils.Random(1, 6)

		password := gen.Generate()
		//test minimum length
		if len(password) < gen.Query.MinLength {
			t.Log("password length fail")
			t.Fail()
			panic(password)
		}

		//test number of digit and specials
		digitCount := 0
		specCount := 0

		for i := 0; i < len(password); i++ {
			for j := 0; j < len(specChar); j++ {
				if password[i] == specChar[j] {
					specCount++
				}
			}
			for k := 0; k < len(digit); k++ {
				if password[i] == digit[k] {
					digitCount++
				}
			}
		}

		if specCount != gen.Query.NumSpecial {
			t.Log("number of specials wrong: ", specCount, password)
			t.Fail()
			panic(password)
		}

		if digitCount != gen.Query.NumDigit {
			t.Log("number of digit wrong: ", digitCount, password)
			t.Fail()
			panic(password)
		}
	}

}

func TestPasswordGenerator_GenerateManyPasswords(t *testing.T) {

	var query = PwdQuery{MinLength: 12, NumSpecial: 2, NumDigit: 3, MaxPwds: 200}
	gen := NewPwdGenerator()
	gen.Query = query
	gen.GenerateManyPasswords()

	if len(gen.Passwords) != gen.Query.MaxPwds {
		t.Log("generate wrong number of passwords ")
		t.Fail()
	}
}

func TestPasswordGenerator_CreateVowelDigitRandomMapping(t *testing.T) {
	gen := NewPwdGenerator()
	mapping := gen.CreateVowelDigitRandomMapping()
	if len(mapping) != 10 {
		t.Fail()
	}
}

func TestPasswordGenerator_Map(t *testing.T) {

	gen := NewPwdGenerator()
	mapping := gen.CreateVowelDigitRandomMapping()

	for i := 0; i < 100; i++ {

		gen.Query.MinLength = utils.Random(13, 40)
		gen.Query.NumSpecial = utils.Random(1, 10)
		gen.Query.NumDigit = utils.Random(1, 12)

		oldPwd := gen.Generate()
		newPwd := gen.Map(mapping, oldPwd)

		//test length
		if len(newPwd) != len(oldPwd) {
			t.Log("wrong mapping length")
			t.Fail()
			panic("mapping length error")
		}

		//test vewol to number mapping
		digitCount := 0
		vewolCount := 0

		for i := 0; i < len(oldPwd); i++ {
			for j := 0; j < len(mixedVowel); j++ {
				if oldPwd[i] == mixedVowel[j] {
					vewolCount++
				}
			}
		}

		for i := 0; i < len(newPwd); i++ {
			for k := 0; k < len(digit); k++ {
				if newPwd[i] == digit[k] {
					digitCount++
				}
			}
		}
		// number of digit + mapping vowel in old passowrd == number of digti in new password
		if digitCount-vewolCount != gen.Query.NumDigit {

			t.Log("vewolCount; ", vewolCount, "mapping: ", oldPwd, "->", newPwd)
			t.Fail()
			panic("wrong mapping number")

		}

	}
}

func BenchmarkPasswordGenerator_GenerateManyPasswords(b *testing.B) {

	b.N = 99999
	var query = PwdQuery{MinLength: 15, NumSpecial: 6, NumDigit: 2, MaxPwds: b.N}
	gen := NewPwdGenerator()
	gen.Query = query
	gen.GenerateManyPasswords()
}
