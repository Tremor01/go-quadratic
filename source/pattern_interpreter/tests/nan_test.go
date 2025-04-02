package tests



import (
	"testing"
	pi "go-quadratic/source/pattern_interpreter"
)


func Test1NanInterpreter(t *testing.T) {
	ctx := pi.NewContext("nan")
	interpret := pi.NewNanInterpreter()
	res := interpret.Interpret(ctx)
	if !res {
		t.Error()
	}
}


func Test2NanInterpreter(t *testing.T) {
	errors := []string{" nan", "", "5", "nan ", "nAn"}
	interpret := pi.NewNanInterpreter()
	for _, error := range errors {
		ctx := pi.NewContext(error)
		res := interpret.Interpret(ctx)
		if (ctx.IsEnd() && res) {
			t.Errorf("Error contex: %s", error)
		}
	}
}