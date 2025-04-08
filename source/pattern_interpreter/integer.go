package pattern_interpreter	


import ("slices")


type IntegerInterpreter struct {
	natruralInterpreter NaturalInterpeter
}

func NewIntegerInterpreter() IntegerInterpreter {
	return IntegerInterpreter{natruralInterpreter: NewNaturalInterpreter()}
}

func (ii IntegerInterpreter) Interpret(ctx Context) int {
	if (ctx.IsEnd()) {return -1}

	if (slices.Contains(SIGNS, ctx.Read())) {
		ctx.Next()
	}

	if (!ctx.IsEnd() && slices.Contains(ZERO, ctx.Read())) {
		ctx.Next()
		return ctx.pos
	}
	return ii.natruralInterpreter.Interpret(ctx)
} 

