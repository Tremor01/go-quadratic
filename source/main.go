package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("1.123")
	inter := pi.NewRationalInterpreter()
	fmt.Println(inter.Interpret(ctx), ctx.IsEnd())
}

