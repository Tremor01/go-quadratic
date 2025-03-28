package pattern_interpreter


import ("slices")


type ExpInterpreter struct {}

func (ri ExpInterpreter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd()) {return false}

	if (!slices.Contains(EXP, ctx.Read())) {
		return false
	}
	ctx.Next()
	if (ctx.IsEnd()) {return false}
	
	if (slices.Contains(SIGNS, ctx.Read())) {
		ctx.Next()
	}
	digitsInterpeter := DigitsInterpreter{}
	return digitsInterpeter.Interpret(ctx)
}
