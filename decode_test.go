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


