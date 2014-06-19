package bpickle

import (
	"fmt"
	"strconv"
	"strings"
)

func Unmarshall(s string) (result interface{}) {
	result, _ = unmarshallInner(s, 0)
	return
}

func unmarshallInner(s string, pos int) (result interface{}, nextPos int) {
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
	default:
		fmt.Println(fmt.Sprintf("Unknown format letter: '%s'", letter))
		return
	}
}

func decodeInt(s string, pos int) (result int64, nextPos int) {
	pos += 1
	var remaining string = s[pos:]
	var endPos int = strings.Index(remaining, ";")
	var toParse = remaining[:endPos]
	result, _ = strconv.ParseInt(toParse, 10, 64)
	return result, endPos + 1
}

func decodeFloat(s string, pos int) (result float64, nextPos int) {
	pos += 1
	var remaining string = s[pos:]
	var endPos int = strings.Index(remaining, ";")
	var toParse = remaining[:endPos]
	result, _ = strconv.ParseFloat(toParse, 64)
	return result, endPos + 1
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
	pos += 1                                    // Skip ":"
	result = remaining[pos : int64(pos)+lenght] // The string, from after ":" to the specified lenght
	nextPos = pos + 1                           // Put the pos at the next position and return
	return
}
