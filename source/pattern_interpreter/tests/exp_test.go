package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1ExpInterpreter(t *testing.T) {
	exps := []string{"E2", "e+2", "e-2", "E9", "e123", "e0", "e-0", "E+0", "e-12"}
	ExpInterpreter := pi.NewExpInterpreter()
	for _, exp := range exps {
		ctx := pi.NewContext(exp)
		res := ExpInterpreter.Interpret(ctx)
		if (res) {
			continue
		}
		t.Error()
	}
}


func Test2ExpInterpreter(t *testing.T) {
	errors := []string{"1e", "-e1", "-12e2", ".e12", "-0e1", "12-e2", "12e-", "+0E1", "12ee+2", "1-E2"}
	ExpInterpreter := pi.NewExpInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := ExpInterpreter.Interpret(ctx)
		if (res) {
			t.Errorf("Error contex: %s", error)
		}
	}
}

