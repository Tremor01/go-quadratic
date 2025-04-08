package pattern_interpreter


import ("slices")


type ExpInterpreter struct {}

func NewExpInterpreter() ExpInterpreter {
	return ExpInterpreter{}
}

func (ei ExpInterpreter) Interpret(ctx Context) int {
	if (ctx.IsEnd()) {return -1}

	if (!slices.Contains(EXP, ctx.Read())) {
		return -1
	}
	ctx.Next()
	if (ctx.IsEnd()) {return -1}
	
	if (slices.Contains(SIGNS, ctx.Read())) {
		ctx.Next()
	}
	digitsInterpeter := DigitsInterpreter{}
	return digitsInterpeter.Interpret(ctx)
}
