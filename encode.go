package bpickle

import (
	"fmt"
	"reflect"
    "strconv"
)

// This is the main interface to this library.
// You can pass "anything" to it, and it will return a bpickle-string for the
// corresponding object.
func Marshall(anything interface{}) string {
	var v reflect.Value = reflect.ValueOf(anything)
	var result string = MarshallValue(v)
	return result
}

// You should pass reflect.Value instances to this.
// It calls itself recursively for Slices and Dicts.
func MarshallValue(v reflect.Value) string {
	var result string

	switch v.Kind() {
	case reflect.Bool:
		result = encodeBool(v.Bool())
	case reflect.Float32:
		result = encodeFloat32(v)
	case reflect.Float64:
		result = encodeFloat64(v)
	case reflect.Int:
		result = encodeInt(v.Int())
	case reflect.String:
		result = encodeString(v.String())
	case reflect.Slice:
		result = encodeSlice(v)
	case reflect.Map:
		result = encodeMap(v)
	case reflect.Interface:
		result = Marshall(v)
	default:
		panic(fmt.Sprintf("Unknown type %s", v.Kind()))
	}
	return result
}

func encodeBool(object bool) string {
	var representation int = 0
	if object {
		representation = 1
	}
	var result string = fmt.Sprintf("b%d", representation)
	return result
}

func encodeInt(object int64) string {
	var result string = fmt.Sprintf("i%d;", object)
	return result
}

func encodeFloat(v reflect.Value, bits int) string {
    f := v.Float()
    var result string = "f"
    result += strconv.FormatFloat(f, 'g', -1, bits)
    result += ";"
    return result
}

func encodeFloat32(v reflect.Value) string {
	return encodeFloat(v, 32)
}

func encodeFloat64(v reflect.Value) string {
	return encodeFloat(v, 64)
}

func encodeString(object string) string {
	var result string = fmt.Sprintf("u:%d:%s", len(object), object)
	return result
}

func encodeSlice(v reflect.Value) string {
	var result string = "l"

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i)
		nested_result := MarshallValue(element)
		result += nested_result
	}

	result += ";"
	return result
}

func encodeMap(v reflect.Value) string {
	var result string = "d"
	var keys []reflect.Value = v.MapKeys()
	for i := range keys {
		result += MarshallValue(keys[i])
		result += MarshallValue(v.MapIndex(keys[i]))
	}
	result += ";"
	return result
}
