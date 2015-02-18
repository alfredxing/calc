package functions

import (
	"math"
)

import (
	"github.com/alfredxing/calc/operators"
)

var (
	sin = &operators.Operator{
		Name:          "sin",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Sin(args[0])
		},
	}
	cos = &operators.Operator{
		Name:          "cos",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Cos(args[0])
		},
	}
	tan = &operators.Operator{
		Name:          "tan",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Tan(args[0])
		},
	}
	cot = &operators.Operator{
		Name:          "cot",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return 1 / math.Tan(args[0])
		},
	}
	sec = &operators.Operator{
		Name:          "sec",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return 1 / math.Cos(args[0])
		},
	}
	csc = &operators.Operator{
		Name:          "csc",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return 1 / math.Sin(args[0])
		},
	}
	asin = &operators.Operator{
		Name:          "asin",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Asin(args[0])
		},
	}
	acos = &operators.Operator{
		Name:          "acos",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Acos(args[0])
		},
	}
	atan = &operators.Operator{
		Name:          "atan",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Atan(args[0])
		},
	}
	acot = &operators.Operator{
		Name:          "acot",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return (-90 * ((math.Pi * math.Atan(args[0]) / 90) - math.Pi)) / math.Pi
		},
	}
	asec = &operators.Operator{
		Name:          "asec",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Acos(1 / args[0])
		},
	}
	acsc = &operators.Operator{
		Name:          "acsc",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Asin(1 / args[0])
		},
	}
)

func init() {
	Register(sin)
	Register(cos)
	Register(tan)
	Register(cot)
	Register(sec)
	Register(csc)
	Register(asin)
	Register(acos)
	Register(atan)
	Register(acot)
	Register(asec)
	Register(acsc)
}
