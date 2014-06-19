package bpickle

import (
	//"fmt"
	"testing"
)

//Integers
func Test_unmarshall_integer_positive(t *testing.T) {
	var result int64 = Unmarshall("i100;").(int64)
	if result != 100 {
		t.Error(result)
	}
}

func Test_unmarshall_integer_negative(t *testing.T) {
	var result int64 = Unmarshall("i-100;").(int64)
	if result != -100 {
		t.Error(result)
	}
}

//Floats
func Test_unmarshall_float_positive(t *testing.T) {
	var input string = "f123.45;"
	var result float64 = Unmarshall(input).(float64)
	if result != 123.45 {
		t.Error(result)
	}
}

func Test_unmarshall_float_negative(t *testing.T) {
	var input string = "f-123.45;"
	var result float64 = Unmarshall(input).(float64)
	if result != -123.45 {
		t.Error(result)
	}
}

// Booleans
func Test_unmarshall_boolean_true(t *testing.T) {
	var result bool = Unmarshall("b1").(bool)
	if result != true {
		t.Fail()
	}
}

func Test_unmarshall_boolean_false(t *testing.T) {
	var result bool = Unmarshall("b0").(bool)
	if result != false {
		t.Fail()
	}
}

//Strings
func Test_unmarshall_string(t *testing.T) {
	var result string = Unmarshall("u:6:string").(string)
	if result != "string" {
		t.Error(result)
	}
}

func Test_unmarshall_string_utf8(t *testing.T) {
	var result string = Unmarshall("u:5:santé").(string)
	if result != "santé" {
		t.Error(result)
	}
}
