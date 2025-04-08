package pattern_interpreter


type NanInterpreter struct {}

func NewNanInterpreter() NanInterpreter {
	return NanInterpreter{}
}

func (ni NanInterpreter) Interpret(ctx Context) int {
	if ctx.IsEnd() {return -1}
	
	var pos int = 0
	for pos < len(NAN) {
		if ctx.IsEnd() || NAN[pos] != ctx.Read() {
			return -1
		}
		ctx.Next()
		pos++
	}
	return ctx.GetPos()
}


 