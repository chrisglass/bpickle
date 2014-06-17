package main

import "fmt"
import "reflect"

type Anything interface{}

func Something (anything Anything) {
    switch reflect.TypeOf(anything).Kind() {
        default:
            fmt.Printf("unexpected type %T", anything)
        case reflect.String:
            fmt.Println("string!")
        case reflect.Int:
            fmt.Println("int!")
        case reflect.Slice:
            fmt.Println("a slice millord")
        case reflect.Array:
            fmt.Println("an array")
        case reflect.Map:
            fmt.Println("a map")
    }
}

// This is just for mucking about while testing/developping, it's not meant for
// anything serious.
func main() {
    var thearray [2]int
    thearray[0] = 1
    thearray[1] = 2

	Something(thearray)
}
