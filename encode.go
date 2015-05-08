package bpickle

import (
	"fmt"
	"reflect"
	"strconv"
)

// This is the main interface to this library.
// You can pass "anything" to it, and it will return a bpickle-string for the
// corresponding object.
func Marshal(anything interface{}) (result string) {
	result = MarshalValue(reflect.ValueOf(anything))
	return
}

// You should pass reflect.Value instances to this.
// It calls itself recursively for Slices and Dicts.
func MarshalValue(v reflect.Value) (result string) {

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
		result = MarshalValue(v.Elem())
	case reflect.Struct:
		fmt.Println(fmt.Sprintf("Unknown struct: %T", v))
	default:
		panic(fmt.Sprintf("Unknown type %s", v.Kind()))
	}
	return
}

func encodeBool(object bool) (result string) {
	representation := 0
	if object {
		representation = 1
	}
	result = fmt.Sprintf("b%d", representation)
	return
}

func encodeInt(object int64) string {
	return fmt.Sprintf("i%d;", object)
}

func encodeFloat(v reflect.Value, bits int) (result string) {
	f := v.Float()
	result = "f"
	result += strconv.FormatFloat(f, 'g', -1, bits)
	result += ";"
	return
}

func encodeFloat32(v reflect.Value) string {
	return encodeFloat(v, 32)
}

func encodeFloat64(v reflect.Value) string {
	return encodeFloat(v, 64)
}

func encodeString(object string) string {
	return fmt.Sprintf("u:%d:%s", len(object), object)
}

func encodeSlice(v reflect.Value) (result string) {
	result = "l"

	for i := 0; i < v.Len(); i++ {
		element := v.Index(i)
		nested_result := MarshalValue(element)
		result += nested_result
	}

	result += ";"
	return
}

func encodeMap(v reflect.Value) (result string) {
	result = "d"
	var keys []reflect.Value = v.MapKeys()
	for _, k := range keys {
		result += MarshalValue(k)
		result += MarshalValue(v.MapIndex(k))
	}
	result += ";"
	return
}
