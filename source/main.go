package main

import (
	"fmt"
	pi "go-quadratic/source/pattern_interpreter"
)


func main() {
	ctx := pi.NewContext("0.123")
	inter := pi.RationalInterpreter{}
	fmt.Println(inter.Interpret(ctx))
}

