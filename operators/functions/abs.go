package functions

import (
	"math"
)

import (
	"github.com/alfredxing/calc/operators"
)

var (
	abs = &operators.Operator{
		Name:          "abs",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Abs(args[0])
		},
	}
)

func init() {
	Register(abs)
}
