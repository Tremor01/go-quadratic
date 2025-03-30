package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1DigitsInterpreter(t *testing.T) {
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	digitsInterpreter := pi.NewDigitsInterpreter()
	for _, digit := range digits {
		ctx := pi.NewContext(digit)
		res := digitsInterpreter.Interpret(ctx)
		if (res) {
			continue
		}
		t.Error()
	}
}


func Test2DigitsInterpreter(t *testing.T) {
	digits := []string{"1021943", "344234219", "3241324", "00012", "1234567890"}
	digitsInterpreter := pi.NewDigitsInterpreter()
	for _, digit := range digits {
		ctx := pi.NewContext(digit)
		res := digitsInterpreter.Interpret(ctx)
		if (res) {
			continue
		}
		t.Error()
	}
}


func Test3DigitsInterpreter(t *testing.T) {
	errors := []string{"0.", "-5", "fdg", "e.1", "123.", "1 1", "  ", "89*", "авыф"}
	digitsInterpreter := pi.NewDigitsInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		digitsInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			t.Error()
		}
	}
}

