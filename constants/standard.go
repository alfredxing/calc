package constants

import (
	"math"
)

var (
	e = &Constant{
		Name:  "e",
		Value: math.E,
	}
	pi = &Constant{
		Name:  "pi",
		Value: math.Pi,
	}
	piSym = &Constant{
		Name:  "π",
		Value: math.Pi,
	}
)

func init() {
	Register(e)
	Register(pi)
	Register(piSym)
}
