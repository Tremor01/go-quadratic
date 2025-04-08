package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1RationalShortSecondInterpreter(t *testing.T) {
	rs := []string{"-123.", "+423.", "213120.", "12.", "0.", "+0.", "-0.", "1234567890."}
	rationlInterpreter := pi.NewRationalShortSecondInterpreter()
	for _, r := range rs {
		ctx := pi.NewContext(r)
		res := rationlInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", r)
	}
}


func Test2RationalShortSecondInterpreter(t *testing.T) {
	errors := []string{
		"-0123.", "..123", ".+.123", "--.123", "", "0", " ", "  ", "авыф",
		"-.asd", "-.123.", ".123.", "+.123123.12321",
	}
	rationlInterpreter := pi.NewRationalShortSecondInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := rationlInterpreter.Interpret(ctx)
		if res == ctx.Len() {
			t.Errorf("Error contex: %s", error)
		}
	}

}
