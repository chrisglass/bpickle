package bpickle

import (
	"fmt"
	"strconv"
	"strings"
)

// The main entry point for decoding bpickle.
// Simly pass it the string to decode and it will return a Go object/structure
// representation.
func Unmarshall(s string) (result interface{}) {
	result, _ = unmarshallInner(s, 0)
	return
}

// This is only there to allow for containers to call it "recursively" (not actually
// recursive, but close).
func unmarshallInner(s string, pos int) (result interface{}, nextPos int) {

	fmt.Println(fmt.Sprintf("Calling unmarshallInner: '%s', pos: '%d'", s, pos))
	letter := string([]rune(s)[pos])
	switch letter {
	case "i":
		return decodeInt(s, pos)
	case "f":
		return decodeFloat(s, pos)
	case "b":
		return decodeBool(s, pos)
	case "u":
		return decodeString(s, pos)
	case "l":
		return decodeSlice(s, pos)
	default:
		panic(fmt.Sprintf("Unknown format letter: '%s'", letter))
	}
}

func decodeInt(s string, pos int) (result int64, nextPos int) {
	pos += 1
	var remaining string = s[pos:]
	var endPos int = strings.Index(remaining, ";")
	var toParse = remaining[:endPos]
	result, _ = strconv.ParseInt(toParse, 10, 64)
	nextPos = endPos + 1
	return
}

func decodeFloat(s string, pos int) (result float64, nextPos int) {
	pos += 1
	var remaining string = s[pos:]
	var endPos int = strings.Index(remaining, ";")
	var toParse = remaining[:endPos]
	result, _ = strconv.ParseFloat(toParse, 64)
	nextPos = endPos + 1
	return
}

func decodeBool(s string, pos int) (result bool, nextPos int) {
	pos += 1
	var err error
	result, err = strconv.ParseBool(s[pos : pos+1])
	if err != nil {
		panic(fmt.Sprintf("Invalid value for boolean '%s'", s[pos:pos+1]))
	}
	nextPos = pos + 1
	return
}

func decodeString(s string, pos int) (result string, nextPos int) {
	pos += 2                            // Skip "u:"
	var remaining string = s[pos:]      // The remaining string, without "u:"
	pos = strings.Index(remaining, ":") // Now the position of the ":" after the int
	var toParse = remaining[:pos]       // the string beween "u:" and the next ":"
	var lenght int64
	lenght, _ = strconv.ParseInt(toParse, 10, 0)
	pos += 1                                        // Skip ":"
	var runes []rune = []rune(remaining)            // We need to count in runes, not in chars/bytes.
	result = string(runes[pos : int64(pos)+lenght]) // The string, from after ":" to the specified rune lenght.
	nextPos = pos + 1                               // Put the pos at the next position and return
	return
}

func decodeSlice(s string, pos int) (result []interface{}, nextPos int) {
	pos += 1 // Skip the "l"
	fmt.Println(fmt.Sprintf("Before loop: string: '%s' pos: '%d'", s, pos))
	for string([]rune(s)[pos]) != ";" {
		var object interface{}
		object, pos = unmarshallInner(s, pos)
		fmt.Println(fmt.Sprintf("Got: '%s' pos: '%d'", object, pos))
		result = append(result, object)
	}
	pos += 1 // Skip the ";"
	nextPos = pos
	return
}
