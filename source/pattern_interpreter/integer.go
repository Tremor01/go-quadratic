package pattern_interpreter	


import ("slices")


type IntegerInterpreter struct {
	natruralInterpreter NaturalInterpeter
}

func NewIntegerInterpreter() IntegerInterpreter {
	return IntegerInterpreter{natruralInterpreter: NewNaturalInterpreter()}
}

func (ii IntegerInterpreter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd()) {return false}

	if (slices.Contains(SIGNS, ctx.Read())) {
		ctx.Next()
	}

	if (!ctx.IsEnd() && slices.Contains(ZERO, ctx.Read())) {
		ctx.Next()
		return true
	}
	return ii.natruralInterpreter.Interpret(ctx)
} 

