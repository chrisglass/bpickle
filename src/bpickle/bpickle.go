package bpickle

import "fmt"


type Anything interface{}

// The main function to create Bpickle strings. You can pass "anything" to it
// and it will take care of converting it to whatever is relevant.
func Dumps(anything Anything) string {
    var result string
    switch anything.(type){
        case bool:
            result = dumpBool(anything.(bool))
        case float32:
            result = dumpFloat32(anything.(float32))
        case float64:
            result = dumpFloat64(anything.(float64))
        case int:
            result = dumpInt(anything.(int))
        case string:
            result = dumpString(anything.(string))
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

func dumpInt(object int) string {
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
