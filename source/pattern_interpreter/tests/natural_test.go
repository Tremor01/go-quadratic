package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1NaturalInterpreter(t *testing.T) {
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	naturalInterpreter := pi.NaturalInterpeter{}
	for _, digit := range digits {
		ctx := pi.NewContext(digit)
		res := naturalInterpreter.Interpret(ctx)
		if (res) {
			continue
		}
		t.Error()
	}
}


func Test2NaturalInterpreter(t *testing.T) {
	digits := []string{"1021943", "344234219", "3241324", "10012", "1234567890"}
	naturalInterpreter := pi.NaturalInterpeter{}
	for _, digit := range digits {
		ctx := pi.NewContext(digit)
		res := naturalInterpreter.Interpret(ctx)
		if (res) {
			continue
		}
		t.Error()
	}
}


func Test3NaturalInterpreter(t *testing.T) {
	errors := []string{"0123", "0.", "-5", "fdg", "e.1", "123.", "1 1", "  ", "89*", "авыф"}
	naturalInterpreter := pi.NaturalInterpeter{}
	for _, error := range errors {
		ctx := pi.NewContext(error)
		naturalInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			t.Errorf("Error contex: %s", error)
		}
	}
}
