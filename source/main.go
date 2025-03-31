package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("+0.")
	inter := pi.NewRationalShortSecondInterpreter()
	fmt.Println(inter.Interpret(ctx), ctx.IsEnd())
}

