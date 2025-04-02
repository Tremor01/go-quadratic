package quadratic


type Output struct {
	x1 string
	x2 string
}

func (out Output) Print() {
	var x2 string
	if out.x2 == "" {
		x2 = out.x1
	} 
	print("x1 = %s\nx2 = %s\n", out.x1, x2)
}

