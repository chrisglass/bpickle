package bpickle

import (
	"fmt"
	"testing"
)

// Booleans
func Test_marshall_boolean_true(t *testing.T) {
	var result string = Marshal(true)
	if result != "b1" {
		t.Fail()
	}
}

func Test_marshall_boolean_false(t *testing.T) {
	var result string = Marshal(false)
	if result != "b0" {
		t.Fail()
	}
}

//Integers
func Test_marshall_integer_positive(t *testing.T) {
	var result string = Marshal(100)
	if result != "i100;" {
		t.Fail()
	}
}

func Test_marshall_integer_negative(t *testing.T) {
	var result string = Marshal(-100)
	if result != "i-100;" {
		t.Fail()
	}
}

//Floats
func Test_marshall_float32_positive(t *testing.T) {
	var input float32 = 123.45
	var result string = Marshal(input)
	if result != "f123.45;" {
		t.Error(result)
	}
}

func Test_marshall_float32_negative(t *testing.T) {
	var input float32 = -123.45
	var result string = Marshal(input)
	if result != "f-123.45;" {
		t.Error(result)
	}
}

//Strings
func Test_marshall_string(t *testing.T) {
	var result string = Marshal("some string")
	if result != "u:11:some string" {
		t.Error(result)
	}
}

func Test_marshall_string_utf8(t *testing.T) {
	var result string = Marshal("une chaîne de caractrères")
	if result != "u:27:une chaîne de caractrères" {
		t.Error(result)
	}
}

//Slices
func Test_marshall_slices(t *testing.T) {
	somearray := [3]int{1, 2, 3}
	someslice := somearray[:]
	var result = Marshal(someslice)
	if result != "li1;i2;i3;;" {
		t.Error(result)
	}
}

func Test_marshall_slices_string(t *testing.T) {
	somearray := [3]string{"test", "test2", "test3"}
	someslice := somearray[:]
	var result = Marshal(someslice)
	if result != "lu:4:testu:5:test2u:5:test3;" {
		t.Error(result)
	}
}

//Maps
func Test_marshall_maps_string_string(t *testing.T) {
	somemap := map[string]string{"test": "blah"}
	var result = Marshal(somemap)
	if result != "du:4:testu:4:blah;" {
		t.Error(result)
	}
}

func Test_marshall_maps_int_string(t *testing.T) {
	somemap := map[int]string{1: "blah"}
	var result = Marshal(somemap)
	if result != "di1;u:4:blah;" {
		t.Error(result)
	}
}

func Test_marshall_maps_int_int(t *testing.T) {
	somemap := map[int]int{1: 23}
	var result = Marshal(somemap)
	if result != "di1;i23;;" {
		t.Error(result)
	}
}

func Test_marshall_maps_string_anything(t *testing.T) {
	submap := map[string]int{"value1": 1, "value2": 2}
	somemap := map[string]interface{}{"type": "some-type", "values": submap}
	var result = Marshal(somemap)
	var expected = "du:4:typeu:9:some-typeu:6:valuesdu:6:value1i1;u:6:value2i2;;;"
	if result != expected {
		t.Error(fmt.Sprintf("%s is not equal to expected %s", result, expected))
	}
}
