package pattern_interpreter


import ("slices")


type RationalShortFirstInterpreter struct{
	digitsInterpreter DigitsInterpreter
}

func NewRationalShortFirstInterpreter() RationalShortFirstInterpreter {
	return RationalShortFirstInterpreter{
		digitsInterpreter: NewDigitsInterpreter(),
	}
}

func (interpreter RationalShortFirstInterpreter) Interpret(ctx Context) int {
	if (ctx.IsEnd()) {return -1}
	
	if (slices.Contains(SIGNS, ctx.Read())) {ctx.Next()}

	if (ctx.IsEnd()) {return -1}
	if (!slices.Contains(POINT, ctx.Read())) {
		return -1
	}
	ctx.Next()
	return interpreter.digitsInterpreter.Interpret(ctx)
}