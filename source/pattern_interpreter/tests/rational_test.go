package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1RationalInterpreter(t *testing.T) {
	numbers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "-4645641", 
		"+54102", "3487", "4999", "123456789", "-68900",
	}
	rationalInterpreter := pi.RationalInterpreter{}
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		rationalInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {continue}
		t.Error()
	}
}


func Test2RationalInterpreter(t *testing.T) {
	numbers := []string{
		"1.123", "+1.23", "-123.21", "0.01654", "0.000978454", "-1.1", 
		"-2.2", "+3.3", "-45456.4", "99.99", "0.123", "4754.57", "-5",
	}
	rationalInterpreter := pi.RationalInterpreter{}
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		rationalInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {continue}
		t.Errorf("Error contex: %s", number)
	}
}


func Test3RationalInterpreter(t *testing.T) {
	errors := []string{
		"00.7987", "0.", "fdg", "e.1", "123.", "1 1", "  ", "89*", "авыф",
		".0.", "00.123", ".213", "0.", "1231.24e",
	}
	rationalInterpreter := pi.RationalInterpreter{}
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := rationalInterpreter.Interpret(ctx)
		if (ctx.IsEnd() && res) {
			t.Errorf("Error contex: %s", error)
		}
	}
}
