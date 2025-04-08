package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("213120.e+0")
	inter := pi.NewNumberInterpreter()
	fmt.Println(inter.Interpret(ctx), ctx.Len())
}

