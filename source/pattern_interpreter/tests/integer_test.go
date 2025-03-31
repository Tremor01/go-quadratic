package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1IntegerInterpreter(t *testing.T) {
	numbers := []string{"+0", "2", "3", "-4", "5", "+6", "7", "0", "-0"}
	integerInterpreter := pi.NewIntegerInterpreter()
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		integerInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			continue
		}
		t.Error()
	}
}


func Test2IntegerInterpreter(t *testing.T) {
	numbers := []string{"-4645641", "+54102", "3487", "4999", "123456789", "-68900"}
	integerInterpreter := pi.NewIntegerInterpreter()
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		integerInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			continue
		}
		t.Error()
	}
}


func Test3IntegerInterpreter(t *testing.T) {
	errors := []string{"--1", "++2", "0-3", "-04", "545/*", "6+2", "7-3", "+8.", "-9.", "00"}
	integerInterpreter := pi.NewIntegerInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		integerInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			t.Errorf("Error contex: %s", error)
		}
	}
}


func Test4IntegerInterpreter(t *testing.T) {
	valid := []string{
		"0", "1", "123456789", "+0", "-0", "+123", "-987", "1", "123",
		"+5", "-7", "999999999999999", "1", "9",
	}
	integerInterpreter := pi.NewIntegerInterpreter()
	for _, v := range valid {
		ctx := pi.NewContext(v)
		integerInterpreter.Interpret(ctx)
		if (ctx.IsEnd()) {
			continue
		}
		t.Errorf("Error contex: %s", v)
	}
}


func Test5IntegerInterpreter(t *testing.T) {
	errors := []string{
		"0123", "1.23", "e12", "1e", "-1e", "1e-", "+", "-e3", "1E+", "1a3", "123!", "1.2e3", "1 23", "12 3", "",
	}
	integerInterpreter := pi.NewIntegerInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := integerInterpreter.Interpret(ctx)
		if (res && ctx.IsEnd()) {
			t.Errorf("Error contex: %s", error)
		}
	}
}