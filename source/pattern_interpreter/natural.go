package pattern_interpreter


import ("slices")


type NaturalInterpeter struct {
	digitsInterpreter DigitsInterpreter
}

func NewNaturalInterpreter() NaturalInterpeter {
	return NaturalInterpeter{digitsInterpreter: NewDigitsInterpreter()}
}

func (ni NaturalInterpeter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd() || slices.Contains(ZERO, ctx.Read())) {
		return false
	}
	return ni.digitsInterpreter.Interpret(ctx)
}

