package pattern_interpreter


import ("slices")


type NaturalInterpeter struct {}

func (ni NaturalInterpeter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd() || slices.Contains(ZERO, ctx.Read())) {
		return false
	}
	digitsInterpreter := DigitsInterpreter{}
	return digitsInterpreter.Interpret(ctx)
}

