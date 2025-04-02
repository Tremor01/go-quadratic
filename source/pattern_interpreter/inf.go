package pattern_interpreter


type InfInterpreter struct {}

func NewInfInterpreter() InfInterpreter {
	return InfInterpreter{}
}

func (ni InfInterpreter) Interpret(ctx *Context) bool {
	if ctx.IsEnd() {return false}
	
	var pos int = 0
	for pos < len(INF) {
		if ctx.IsEnd() || INF[pos] != ctx.Read() {
			return false
		}
		ctx.Next()
		pos++
	}
	return true
}


 