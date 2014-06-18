package bpickle

import (
	"fmt"
	"reflect"
)

// This is the main interface to this library.
// You can pass "anything" to it, and it will return a bpickle-string for the
// corresponding object.
func Dumps(anything interface{}) string {
	var v reflect.Value = reflect.ValueOf(anything)
	var result string = Marshall(v)
	return result
}

// You should pass reflect.Value instances to this.
// It calls itself recursively for Slices and Dicts.
func Marshall(v reflect.Value) string {
	var result string

	switch v.Kind() {
	case reflect.Bool:
		result = dumpBool(v.Bool())
	case reflect.Float32:
		result = dumpFloat64(v.Float())
	case reflect.Float64:
		result = dumpFloat64(v.Float())
	case reflect.Int:
		result = dumpInt(v.Int())
	case reflect.String:
		result = dumpString(v.String())
	case reflect.Slice:
		result = dumpSlice(v)
	case reflect.Map:
		result = dumpMap(v)
    case reflect.Interface:
        value := reflect.ValueOf(v)
        result = Marshall(value)
	default:
        panic(fmt.Sprintf("Unknown type %s", v.Kind()))
	}
	return result
}

func dumpBool(object bool) string {
	var representation int = 0
	if object {
		representation = 1
	}
	var result string = fmt.Sprintf("b%d", representation)
	return result
}

func dumpInt(object int64) string {
	var result string = fmt.Sprintf("i%d;", object)
	return result
}

func dumpFloat32(object float32) string {
	var result string = fmt.Sprintf("i%r;", object)
	return result
}

func dumpFloat64(object float64) string {
	var result string = fmt.Sprintf("i%r;", object)
	return result
}

func dumpString(object string) string {
	var result string = fmt.Sprintf("u:%d:%s", len(object), object)
	return result
}

func dumpSlice(v reflect.Value) string {
	var result string = "l"

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i)
		nested_result := Marshall(element)
		result += nested_result
	}

	result += ";"
	return result
}

// This is limited to maps of type map[string]Value for now.
func dumpMap(v reflect.Value) string {
	var result string = "d"
	var keys []reflect.Value = v.MapKeys()
	for i := range keys {
		result += Marshall(keys[i])
		result += Marshall(v.MapIndex(keys[i]))
	}
	result += ";"
	return result
}
