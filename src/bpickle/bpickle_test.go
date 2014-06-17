package bpickle

import (
	"testing"
)

func Test_boolean_true(t *testing.T) {
	var result string = Dumps(true)
	if result != "b1" {
		t.Fail()
	}
}

func Test_boolean_false(t *testing.T) {
	var result string = Dumps(false)
	if result != "b0" {
		t.Fail()
	}
}

func Test_integer_positive(t *testing.T) {
	var result string = Dumps(100)
	if result != "i100;" {
		t.Fail()
	}
}

func Test_integer_negative(t *testing.T) {
	var result string = Dumps(-100)
	if result != "i-100;" {
		t.Fail()
	}
}

func Test_string(t *testing.T) {
	var result string = Dumps("some string")
	if result != "u:11:some string" {
		t.Error(result)
	}
}

func Test_string_utf8(t *testing.T) {
	var result string = Dumps("une chaîne de caractrères")
	if result != "u:27:une chaîne de caractrères" {
		t.Error(result)
	}
}

//func Test_float32_positive(t *testing.T) {
func float32_positive(t *testing.T) {
	var input float32 = 123.45
	var result string = dumpFloat32(input)
	if result != "f123.45;" {
		t.Error(result)
	}
}

//func Test_float32_negative(t *testing.T) {
func float32_negative(t *testing.T) {
	var input float32 = -123.45
	var result string = dumpFloat32(input)
	if result != "f-123.45;" {
		t.Error(result)
	}
}
