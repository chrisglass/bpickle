package bpickle

import (
    "testing"
)

func Test_boolean_true(t *testing.T) {
    var result string = dumpBool(true)
    if result != "b1" {
        t.Fail()
    }
}

func Test_boolean_false(t *testing.T) {
    var result string = dumpBool(false)
    if result != "b0" {
        t.Fail()
    }
}

func Test_integer_positive(t *testing.T) {
    var result string = dumpInt(100)
    if result != "i100;" {
        t.Fail()
    }
}

func Test_integer_negative(t *testing.T) {
    var result string = dumpInt(-100)
    if result != "i-100;" {
        t.Fail()
    }
}
