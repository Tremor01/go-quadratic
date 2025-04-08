package pattern_interpreter


import ("slices")


type NaturalInterpeter struct {
	digitsInterpreter DigitsInterpreter
}

func NewNaturalInterpreter() NaturalInterpeter {
	return NaturalInterpeter{digitsInterpreter: NewDigitsInterpreter()}
}

func (ni NaturalInterpeter) Interpret(ctx Context) int {
	if (ctx.IsEnd() || slices.Contains(ZERO, ctx.Read())) {
		return -1
	}
	return ni.digitsInterpreter.Interpret(ctx)
}

