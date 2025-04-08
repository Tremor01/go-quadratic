package pattern_interpreter


import ("slices")


type DigitsInterpreter struct {}

func NewDigitsInterpreter() DigitsInterpreter {
	return DigitsInterpreter{}
}

func (di DigitsInterpreter) Interpret(ctx Context) int {
    if ctx.IsEnd() || !isDigit(ctx.Read()) {  
        return -1
    }
    
    for !ctx.IsEnd() {
        ctx.Next()
        if ctx.IsEnd() || !isDigit(ctx.Read()) {
            break
        }
    }
    return ctx.GetPos()
}


func isDigit(r rune) bool {
    return slices.Contains(NOT_ZERO, r) || slices.Contains(ZERO, r)
}