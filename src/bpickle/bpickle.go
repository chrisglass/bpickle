package bpickle

import (
    "fmt"
    "reflect")

type Anything interface{}

// The main function to create Bpickle strings. You can pass "anything" to it
// and it will take care of converting it to whatever is relevant.
func Dumps(anything Anything) string {
	var result string

	switch reflect.TypeOf(anything).Kind() {
	case reflect.Bool:
		result = dumpBool(reflect.ValueOf(anything).Bool())
	case reflect.Float32:
		result = dumpFloat64(reflect.ValueOf(anything).Float())
	case reflect.Float64:
		result = dumpFloat64(reflect.ValueOf(anything).Float())
	case reflect.Int:
		result = dumpInt(reflect.ValueOf(anything).Int())
	case reflect.String:
		result = dumpString(reflect.ValueOf(anything).String())
    case reflect.Slice:
        var value = reflect.ValueOf(anything)
        result = dumpSlice(value)
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

func dumpSlice(object []Anything) {
    var result string = "l"
    for i := range object {
        result += Dumps(i)
    }
    result += ";"
    return result
}
