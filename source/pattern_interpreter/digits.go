package pattern_interpreter


import ("slices")


type DigitsInterpreter struct {}

func NewDigitsInterpreter() DigitsInterpreter {
	return DigitsInterpreter{}
}

func (di DigitsInterpreter) Interpret(ctx *Context) bool {
    if ctx.IsEnd() || !isDigit(ctx.Read()) {  
        return false
    }
    
    for !ctx.IsEnd() {
        ctx.Next()
        if ctx.IsEnd() || !isDigit(ctx.Read()) {
            break
        }
    }
    return true
}


func isDigit(r rune) bool {
    return slices.Contains(NOT_ZERO, r) || slices.Contains(ZERO, r)
}