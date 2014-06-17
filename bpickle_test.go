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

func Test_slices(t *testing.T) {
	somearray := [3]int{1, 2, 3}
	someslice := somearray[:]
	var result = Dumps(someslice)
	if result != "li1;i2;i3;;" {
		t.Error(result)
	}
}

func Test_slices_string(t *testing.T) {
	somearray := [3]string{"test", "test2", "test3"}
	someslice := somearray[:]
	var result = Dumps(someslice)
	if result != "lu:4:testu:5:test2u:5:test3;" {
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
