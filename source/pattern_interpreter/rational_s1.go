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

func (interpreter RationalShortFirstInterpreter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd()) {return false}
	
	if (slices.Contains(SIGNS, ctx.Read())) {ctx.Next()}

	if (ctx.IsEnd()) {return false}
	if (!slices.Contains(POINT, ctx.Read())) {
		return false
	}
	ctx.Next()
	return interpreter.digitsInterpreter.Interpret(ctx)
}