package operators

import (
	"math"
)

var (
	add = &Operator{
		Name:          "+",
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(args []float64) float64 {
			return args[0] + args[1]
		},
	}
	sub = &Operator{
		Name:          "-",
		Precedence:    1,
		Associativity: L,
		Args:          2,
		Operation: func(args []float64) float64 {
			return args[0] - args[1]
		},
	}
	neg = &Operator{
		Name:          "neg",
		Precedence:    2,
		Associativity: L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return 0 - args[0]
		},
	}
	mul = &Operator{
		Name:          "*",
		Precedence:    2,
		Associativity: L,
		Args:          2,
		Operation: func(args []float64) float64 {
			return args[0] * args[1]
		},
	}
	div = &Operator{
		Name:          "/",
		Precedence:    2,
		Associativity: L,
		Args:          2,
		Operation: func(args []float64) float64 {
			return args[0] / args[1]
		},
	}
	mod = &Operator{
		Name:          "%",
		Precedence:    2,
		Associativity: L,
		Args:          2,
		Operation: func(args []float64) float64 {
			return math.Mod(args[0], args[1])
		},
	}
	pow = &Operator{
		Name:          "^",
		Precedence:    3,
		Associativity: R,
		Args:          2,
		Operation: func(args []float64) float64 {
			return math.Pow(args[0], args[1])
		},
	}
)

func init() {
	Register(add)
	Register(sub)
	Register(neg)
	Register(pow)
	Register(mul)
	Register(mod)
	Register(div)
}
