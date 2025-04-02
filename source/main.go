package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("")
	inter := pi.NewNanInterpreter()
	fmt.Println(inter.Interpret(ctx), ctx.IsEnd())
}

