package pattern_interpreter


import ("slices")


type RationalShortSecondInterpreter struct{
	integerInterpreter IntegerInterpreter
}

func NewRationalShortSecondInterpreter() RationalShortSecondInterpreter {
	return RationalShortSecondInterpreter{
		integerInterpreter: NewIntegerInterpreter(),
	}
}

func (interpreter RationalShortSecondInterpreter) Interpret(ctx Context) int {
	var pos int = interpreter.integerInterpreter.Interpret(ctx)
	if pos == ctx.Len() {return -1}
	
	if pos > -1 && slices.Contains(POINT, ctx.ReadAt(pos)) {
		return pos + 1
	}
	return -1
}