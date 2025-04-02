package quadratic


type Coefficients struct {
	a string
	b string
	c string
}

func NewCoefficients(a string, b string, c string) Coefficients {
	return Coefficients{a: a, b: b, c: c}
}

func (coefficients Coefficients) GetA() string {
	return coefficients.a
}

func (coefficients Coefficients) GetB() string {
	return coefficients.b
}

func (coefficients Coefficients) GetC() string {
	return coefficients.c
}