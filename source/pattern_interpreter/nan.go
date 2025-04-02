package pattern_interpreter


type NanInterpreter struct {}

func NewNanInterpreter() NanInterpreter {
	return NanInterpreter{}
}

func (ni NanInterpreter) Interpret(ctx *Context) bool {
	if ctx.IsEnd() {return false}
	
	var pos int = 0
	for pos < len(NAN) {
		if ctx.IsEnd() || NAN[pos] != ctx.Read() {
			return false
		}
		ctx.Next()
		pos++
	}
	return true
}


 