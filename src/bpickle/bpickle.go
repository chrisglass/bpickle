package bpickle

import "fmt"

func Test() {
    fmt.Println(dumpBool(false))
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

func dumpFloat(object float32) string {
    var result string = fmt.Sprintf("i%f", object)
    return result
}
