package main

import "bpickle"
import "fmt"

// This is just for mucking about while testing/developping, it's not meant for
// anything serious.
func main() {
	var result = bpickle.Dumps(666)
	fmt.Println(result)
}
