package bpickle

import (
	"fmt"
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

func Test_decode_int_direct(t *testing.T) {
	result, pos := decodeInt("i123;", 0)
	if result != 123 {
		t.Error(fmt.Sprintf("Wrong result: '%d', should be 123", result))
	}
	if pos != 5 {
		t.Error(fmt.Sprintf("Wrong pos: '%d', should be 5", pos))
	}
}

func Test_decode_int_offset(t *testing.T) {
	result, pos := decodeInt("i123;i456;", 5)
	if result != 456 {
		t.Error(fmt.Sprintf("Wrong result: '%d', should be 456", result))
	}
	if pos != 10 {
		t.Error(fmt.Sprintf("Wrong pos: '%d', should be 10", pos))
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

func Test_decode_float_direct(t *testing.T) {
	result, pos := decodeFloat("f123.456;", 0)
	if result != 123.456 {
		t.Error(fmt.Sprintf("Wrong result: '%d', should be 123.456", result))
	}
	if pos != 9 {
		t.Error(fmt.Sprintf("Wrong pos: '%d', should be 9", pos))
	}
}

func Test_decode_float_offset(t *testing.T) {
	result, pos := decodeFloat("i123;f123.456;", 5)
	if result != 123.456 {
		t.Error(fmt.Sprintf("Wrong result: '%d', should be 123.456", result))
	}
	if pos != 14 {
		t.Error(fmt.Sprintf("Wrong pos: '%d', should be 14", pos))
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

func Test_decode_bool_offset(t *testing.T) {
	result, pos := decodeBool("i123;b1;", 5)
	if result != true {
		t.Error()
	}
	if pos != 8 {
		t.Error(fmt.Sprintf("Wrong pos: '%d', should be 8", pos))
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

// DOES NOT WORK FOR NOW
func Test_decode_string_offset(t *testing.T) {
	result, pos := decodeString("i123;u:4:test", 5)

	if result != "test" {
		t.Error(result)
	}
	if pos != 14 {
		t.Error(fmt.Sprintf("Position is '%d', should be 14", pos))
	}
}

// Lists (slices)
func Test_unmarshall_slice(t *testing.T) {
	var result []interface{} = Unmarshall("li123;i456;;").([]interface{})
	for _, v := range result {
		fmt.Println(int(v.(int64)))
	}
}
