package bpickle

import (
	//"fmt"
	"testing"
)

//Integers
func Test_unmarshall_integer_positive(t *testing.T) {
	var result int64 = decodeInt("i100;")
	if result != 100 {
		t.Error(result)
	}
}

func Test_unmarshall_integer_negative(t *testing.T) {
	var result int64 = decodeInt("i-100;")
	if result != -100 {
		t.Error(result)
	}
}

//Floats
func Test_unmarshall_float_positive(t *testing.T) {
	var input string = "f123.45;"
	var result float64 = decodeFloat(input)
	if result != 123.45 {
		t.Error(result)
	}
}

func Test_unmarshall_float_negative(t *testing.T) {
	var input string = "f-123.45;"
	var result float64 = decodeFloat(input)
	if result != -123.45 {
		t.Error(result)
	}
}
