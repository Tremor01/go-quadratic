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

func (interpreter RationalShortSecondInterpreter) Interpret(ctx *Context) bool {
	var is_int bool = interpreter.integerInterpreter.Interpret(ctx)
	if(ctx.IsEnd()) {return false}
	
	if (is_int && slices.Contains(POINT, ctx.Read())) {
		ctx.Next()
		return true
	}
	return false
}