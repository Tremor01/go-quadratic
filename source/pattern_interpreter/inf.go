package pattern_interpreter


type InfInterpreter struct {}

func NewInfInterpreter() InfInterpreter {
	return InfInterpreter{}
}

func (ni InfInterpreter) Interpret(ctx Context) int {
	if ctx.IsEnd() {return -1}
	
	var pos int = 0
	for pos < len(INF) {
		if ctx.IsEnd() || INF[pos] != ctx.Read() {
			return -1
		}
		ctx.Next()
		pos++
	}
	return ctx.GetPos()
}


 