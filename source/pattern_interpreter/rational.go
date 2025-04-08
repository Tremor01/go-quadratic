package pattern_interpreter


import ("slices")


type RationalInterpreter struct {
	integerInterpreter IntegerInterpreter
	digitsInterpreter DigitsInterpreter
}

func NewRationalInterpreter() RationalInterpreter {
	return RationalInterpreter{
		integerInterpreter: NewIntegerInterpreter(),
		digitsInterpreter: NewDigitsInterpreter(),
	}
}

func (ri RationalInterpreter) Interpret(ctx Context) int {
	pos := ri.integerInterpreter.Interpret(ctx)
	if (pos == -1) {
		return -1
	}
	if pos == ctx.Len() || !slices.Contains(POINT, ctx.ReadAt(pos)) {
		return pos
	}
	ctx = ctx.SetPos(pos + 1)
	return ri.digitsInterpreter.Interpret(ctx)
}
