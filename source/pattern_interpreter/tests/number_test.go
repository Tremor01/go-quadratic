package tests

import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1NumberInterpreter(t *testing.T) {
	numbers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "-4645641", 
		"+54102", "3487", "4999", "123456789", "-68900",
	}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		res := numberInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Error()
	}
}


func Test2NumberInterpreter(t *testing.T) {
	numbers := []string{
		"1.123", "+1.23", "-123.21", "0.01654", "0.000978454", "-1.1", 
		"-2.2", "+3.3", "-45456.4", "99.99", "0.123", "4754.57", "-5",
	}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		res := numberInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", number)
	}
}


func Test3NumberInterpreter(t *testing.T) {
	numbers := []string{
		"1.123e12", "+1.23e+12", "-123.21e-12", "0.01654e0", "0.000978454e+0", "-1.1e-0", 
		"-2.2e12345", "+3.3e+1234", "-45456.4e1", "99.99e-2", "0.123e12", "4754.57e0", "-5e1",
	}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, number := range numbers {
		ctx := pi.NewContext(number)
		res := numberInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", number)
	}
}


func Test4NumberInterpreter(t *testing.T) {
	rs := []string{".0e0", "-.2e+0", "+.3e-0", "+.0e+123", "-.0e-123", "-.45460048e34", "+.0123456789e12", ".45127e99", ".30082004e-2"}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, r := range rs {
		ctx := pi.NewContext(r)
		res := numberInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", r)
	}
}


func Test5NumberInterpreter(t *testing.T) {
	rs := []string{"-123.e0", "+423.e-0", "213120.e+0", "12.e124", "0.e-563", "+0.e3008", "-0.e-2004", "1234567890.e2007", "1.", "0.", ".0"}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, r := range rs {
		ctx := pi.NewContext(r)
		res := numberInterpreter.Interpret(ctx)
		if res == ctx.Len() {continue}
		t.Errorf("Error contex: %s", r)
	}
}


func Test6NumberInterpreter(t *testing.T) {
	rs := []string{"-123.ee0", "+423.e+-0", "213120.e+0.", "12.e124-", "0.e-563/", "+0.e3008e", "-0.ee-2004", "1234567e890.e2007"}
	numberInterpreter := pi.NewNumberInterpreter()
	for _, r := range rs {
		ctx := pi.NewContext(r)
		res := numberInterpreter.Interpret(ctx)
		if ctx.Len() == res {t.Errorf("Error contex: %s", r)}
	}
}
