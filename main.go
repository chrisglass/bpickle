package main

import "bpickle"
import "fmt"

func main() {
	var result = bpickle.Dumps(666)
	fmt.Println(result)
}
