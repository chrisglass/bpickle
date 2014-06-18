package bpickle

import (
	"testing"
)

// Booleans
func Test_boolean_true(t *testing.T) {
	var result string = Marshall(true)
	if result != "b1" {
		t.Fail()
	}
}

func Test_boolean_false(t *testing.T) {
	var result string = Marshall(false)
	if result != "b0" {
		t.Fail()
	}
}

//Integers
func Test_integer_positive(t *testing.T) {
	var result string = Marshall(100)
	if result != "i100;" {
		t.Fail()
	}
}

func Test_integer_negative(t *testing.T) {
	var result string = Marshall(-100)
	if result != "i-100;" {
		t.Fail()
	}
}

//Floats
func Test_float32_positive(t *testing.T) {
	var input float32 = 123.45
	var result string = Marshall(input)
	if result != "f123.45;" {
		t.Error(result)
	}
}

func Test_float32_negative(t *testing.T) {
	var input float32 = -123.45
	var result string = Marshall(input)
	if result != "f-123.45;" {
		t.Error(result)
	}
}

//Strings
func Test_string(t *testing.T) {
	var result string = Marshall("some string")
	if result != "u:11:some string" {
		t.Error(result)
	}
}

func Test_string_utf8(t *testing.T) {
	var result string = Marshall("une chaîne de caractrères")
	if result != "u:27:une chaîne de caractrères" {
		t.Error(result)
	}
}

//Slices
func Test_slices(t *testing.T) {
	somearray := [3]int{1, 2, 3}
	someslice := somearray[:]
	var result = Marshall(someslice)
	if result != "li1;i2;i3;;" {
		t.Error(result)
	}
}

func Test_slices_string(t *testing.T) {
	somearray := [3]string{"test", "test2", "test3"}
	someslice := somearray[:]
	var result = Marshall(someslice)
	if result != "lu:4:testu:5:test2u:5:test3;" {
		t.Error(result)
	}
}

//Maps
func Test_maps_string_string(t *testing.T) {
	somemap := map[string]string{"test": "blah"}
	var result = Marshall(somemap)
	if result != "du:4:testu:4:blah;" {
		t.Error(result)
	}
}

func Test_maps_int_string(t *testing.T) {
	somemap := map[int]string{1: "blah"}
	var result = Marshall(somemap)
	if result != "di1;u:4:blah;" {
		t.Error(result)
	}
}

func Test_maps_int_int(t *testing.T) {
	somemap := map[int]int{1: 23}
	var result = Marshall(somemap)
	if result != "di1;i23;;" {
		t.Error(result)
	}
}

// THESE DONT ACTUALLY RUN YET TODO: FIX ETC...
//func Test_maps_string_anything(t *testing.T){
func maps_string_anything(t *testing.T) {
	submap := map[string]int{"value1": 1, "value2": 2}
	somemap := map[string]interface{}{"type": "some-type", "values": submap}
	var result = Marshall(somemap)
	var expected = "du:4:typeu:9:some-typeu:5:valuesdu:5:value1i1;u:5:value2i2;;"
	if result != expected {
		t.Error(result)
	}
}


