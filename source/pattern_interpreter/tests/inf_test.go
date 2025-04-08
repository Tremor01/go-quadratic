package tests



import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1InfInterpreter(t *testing.T) {
	ctx := pi.NewContext("inf")
	interpret := pi.NewInfInterpreter()
	res := interpret.Interpret(ctx)
	if res == -1 {
		t.Error()
	}
}


func Test2InfInterpreter(t *testing.T) {
	errors := []string{" inf", "", "5", "inf ", "INF"}
	interpret := pi.NewInfInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := interpret.Interpret(ctx)
		if (ctx.Len() == res) {
			t.Errorf("Error contex: %s", error)
		}
	}
}