package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1RationalShortFirstInterpreter(t *testing.T) {
	rs := []string{".0", "-.2", "+.3", "+.0", "-.0", "-.45460048", "+.0123456789", ".45127", ".30082004"}
	rationlInterpreter := pi.NewRationalShortFirstInterpreter()
	for _, r := range rs {
		ctx := pi.NewContext(r)
		res := rationlInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", r)
	}
}


func Test2RationalShortFirstInterpreter(t *testing.T) {
	errors := []string{
		"+1.56", "..123", ".+.123", "--.123", "", "0", "+0.", " ", "  ", "авыф",
		"-.asd", "-.123.", ".123.", "+.123123.12321",
	}
	rationlInterpreter := pi.NewRationalShortFirstInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := rationlInterpreter.Interpret(ctx)
		if res == ctx.Len() {
			t.Errorf("Error contex: %s", error)
		}
	}
}
