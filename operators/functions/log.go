package functions

import (
	"math"
)

import (
	"github.com/alfredxing/calc/operators"
)

var (
	log = &operators.Operator{
		Name:          "log",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Log10(args[0])
		},
	}
	ln = &operators.Operator{
		Name:          "ln",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Log(args[0])
		},
	}
	lg = &operators.Operator{
		Name:          "lg",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Log2(args[0])
		},
	}
)

func init() {
	Register(log)
	Register(ln)
	Register(lg)
}
