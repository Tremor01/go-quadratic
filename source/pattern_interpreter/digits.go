package pattern_interpreter


import ("slices")


type DigitInterpreter struct {}

func (di DigitInterpreter) Interpret(ctx *Context) bool {
	if ctx.IsEnd() {
		return false
	}
	return slices.Contains(NOT_ZERO, ctx.Read()) || slices.Contains(ZERO, ctx.Read()) 
}


type DigitsInterpreter struct {}

func (di DigitsInterpreter) Interpret(ctx *Context) bool {
	digitInterpreter := DigitInterpreter{}
	var is_digit bool = digitInterpreter.Interpret(ctx)
	if (!is_digit) {return false}
	
	for is_digit && !ctx.IsEnd() {
		ctx.Next()
		is_digit = digitInterpreter.Interpret(ctx)
	}
	return true
}

