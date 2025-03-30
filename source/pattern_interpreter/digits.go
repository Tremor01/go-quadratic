package pattern_interpreter


import ("slices")


type DigitInterpreter struct {}

func NewDigitInterpreter() DigitInterpreter {
	return DigitInterpreter{}
}

func (di DigitInterpreter) Interpret(ctx *Context) bool {
	if ctx.IsEnd() {
		return false
	}
	return slices.Contains(NOT_ZERO, ctx.Read()) || slices.Contains(ZERO, ctx.Read()) 
}


type DigitsInterpreter struct {
	digitInterpreter DigitInterpreter
}

func NewDigitsInterpreter() DigitsInterpreter {
	return DigitsInterpreter{digitInterpreter: NewDigitInterpreter()}
}

func (di DigitsInterpreter) Interpret(ctx *Context) bool {
	var is_digit bool = di.digitInterpreter.Interpret(ctx)
	if (!is_digit) {return false}
	
	for is_digit && !ctx.IsEnd() {
		ctx.Next()
		is_digit = di.digitInterpreter.Interpret(ctx)
	}
	return true
}
