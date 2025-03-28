package pattern_interpreter


import ("slices")


type RationalInterpreter struct {}

func (ri RationalInterpreter) Interpret(ctx *Context) bool {
	integerInterpreter := IntegerInterpreter{}
	if (!integerInterpreter.Interpret(ctx)) {
		return false
	}
	
	if (ctx.IsEnd()) {return true}

	if (!slices.Contains(POINT, ctx.Read())) {
		return false
	}
	ctx.Next()
	digitsInterpreter := DigitsInterpreter{}
	return digitsInterpreter.Interpret(ctx)
}
