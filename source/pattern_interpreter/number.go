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

func (fi NumberInterpreter) Interpret(ctx Context) int {
	pos := fi.rationalInterpreter.Interpret(ctx)
	if pos > -1 {
		ctx = ctx.SetPos(pos)
		return fi.is_exp(ctx)
	}
	pos = fi.rationalShortFirstInterpreter.Interpret(ctx)
	if pos > -1 {
		ctx = ctx.SetPos(pos)
		return fi.is_exp(ctx)
	}
	pos = fi.rationalShortSecondInterpreter.Interpret(ctx)
	if pos > -1 {
		ctx = ctx.SetPos(pos)
		return fi.is_exp(ctx)
	}
	return -1
}

func (fi NumberInterpreter) is_exp(ctx Context) int {
	if ctx.IsEnd() {
		return ctx.pos
	}
	return fi.expInterpreter.Interpret(ctx)
}
 