package pattern_interpreter


type NumberInterpreter struct {
	expInterpreter ExpInterpreter
	rationalInterpreter            RationalInterpreter
	rationalShortFirstInterpreter  RationalShortFirstInterpreter
	rationalShortSecondInterpreter RationalShortSecondInterpreter
}

func NewNumberInterpreter() NumberInterpreter {
	return NumberInterpreter{
		expInterpreter: NewExpInterpreter(),
		rationalInterpreter:            NewRationalInterpreter(),
		rationalShortFirstInterpreter:  NewRationalShortFirstInterpreter(),
		rationalShortSecondInterpreter: NewRationalShortSecondInterpreter(),
	}
}

func (fi NumberInterpreter) Interpret(ctx *Context) bool {
	if fi.rationalInterpreter.Interpret(ctx) {
		return fi.is_exp(ctx)
	}
	ctx.ResetPos()
	if fi.rationalShortFirstInterpreter.Interpret(ctx) {
		return fi.is_exp(ctx)
	}
	ctx.ResetPos()
	if fi.rationalShortSecondInterpreter.Interpret(ctx) {
		return fi.is_exp(ctx)
	}
	return false
}

func (fi NumberInterpreter) is_exp(ctx *Context) bool {
	if ctx.IsEnd() {
		return true
	}
	return fi.expInterpreter.Interpret(ctx)
}
 