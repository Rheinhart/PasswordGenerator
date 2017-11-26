package pwdGenerator

import (
	"reflect"
	"testing"
)

func TestCharBytes_RandomGetChars(t *testing.T) {

	many := 10
	out := mixedVowel.RandomGetChars(many)
	if len(out) != many {
		t.Log("get number wrong!")
		t.Fail()
	}
}

func TestNewPwdGenerator(t *testing.T) {

	p := NewPwdGenerator()
	q := &passwordGenerator{}
	v := reflect.TypeOf(p)
	if v == reflect.TypeOf(q) {
		t.Log("Right right type")
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
