package pattern_interpreter	


import ("slices")


type IntegerInterpreter struct {}

func (ii IntegerInterpreter) Interpret(ctx *Context) bool {
	if (ctx.IsEnd()) {return false}

	if (slices.Contains(SIGNS, ctx.Read())) {
		ctx.Next()
	}

	if (!ctx.IsEnd() && slices.Contains(ZERO, ctx.Read())) {
		ctx.Next()
		return true
	}

	natruralInterpreter := NaturalInterpeter{}
	return natruralInterpreter.Interpret(ctx)
} 

