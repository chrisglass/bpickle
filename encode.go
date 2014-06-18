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
		result = dumpBool(v.Bool())
	case reflect.Float32:
		result = dumpFloat32(v)
	case reflect.Float64:
		result = dumpFloat64(v)
	case reflect.Int:
		result = dumpInt(v.Int())
	case reflect.String:
		result = dumpString(v.String())
	case reflect.Slice:
		result = dumpSlice(v)
	case reflect.Map:
		result = dumpMap(v)
	case reflect.Interface:
		result = Marshall(v)
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

func dumpFloat(v reflect.Value, bits int) string {
    f := v.Float()
    var result string = "f"
    result += strconv.FormatFloat(f, 'g', -1, bits)
    result += ";"
    return result
}

func dumpFloat32(v reflect.Value) string {
	return dumpFloat(v, 32)
}

func dumpFloat64(v reflect.Value) string {
	return dumpFloat(v, 64)
}

func dumpString(object string) string {
	var result string = fmt.Sprintf("u:%d:%s", len(object), object)
	return result
}

func dumpSlice(v reflect.Value) string {
	var result string = "l"

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i)
		nested_result := MarshallValue(element)
		result += nested_result
	}

	result += ";"
	return result
}

func dumpMap(v reflect.Value) string {
	var result string = "d"
	var keys []reflect.Value = v.MapKeys()
	for i := range keys {
		result += MarshallValue(keys[i])
		result += MarshallValue(v.MapIndex(keys[i]))
	}
	result += ";"
	return result
}
