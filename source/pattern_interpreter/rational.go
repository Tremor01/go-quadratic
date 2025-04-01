package pattern_interpreter


import ("slices")


type RationalInterpreter struct {
	integerInterpreter IntegerInterpreter
	digitsInterpreter DigitsInterpreter
}

func NewRationalInterpreter() RationalInterpreter {
	return RationalInterpreter{
		integerInterpreter: NewIntegerInterpreter(),
		digitsInterpreter: NewDigitsInterpreter(),
	}
}

func (ri RationalInterpreter) Interpret(ctx *Context) bool {
	if (!ri.integerInterpreter.Interpret(ctx)) {
		return false
	}

	if (ctx.IsEnd() || !slices.Contains(POINT, ctx.Read())) {
		return true
	}

	ctx.Next()
	return ri.digitsInterpreter.Interpret(ctx)
}
