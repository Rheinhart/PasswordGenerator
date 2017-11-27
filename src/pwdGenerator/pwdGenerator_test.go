package pwdGenerator

import (
	"reflect"
	"testing"
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

	p := NewPwdGenerator()
	q := &passwordGenerator{}
	v := reflect.TypeOf(p)
	if v == reflect.TypeOf(q) {
		t.Log("Return right type")
	} else {
		t.Log("Return wrong type")
		t.Fail()
	}
}

func TestPasswordGenerator_Generate(t *testing.T) {

}

func TestPasswordGenerator_Map(t *testing.T) {

}

func TestPasswordGenerator_CreateVowelDigitRandomMapping(t *testing.T) {

}
