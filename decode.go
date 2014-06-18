package bpickle

import (
	"fmt"
    "strconv"
)

func Unmarshall(value string) {
	fmt.Println("Not implemented yet")
	return
}

func decodeInt(value string) int64 {
    extracted := value[1:len(value) -1]
    result , _ := strconv.ParseInt(extracted, 10, 64)
    return result
}
