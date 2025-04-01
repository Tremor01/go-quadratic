package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("-5e1")
	inter := pi.NewNumberInterpreter()
	fmt.Println(inter.Interpret(ctx), ctx.IsEnd())
}

