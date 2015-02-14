package functions

import "math"

var (
    sqrt = &Function{
        Name: "sqrt",
        Call: func (val float64) float64 {
            return math.Sqrt(val)
        },
    }
)

func init() {
    Register(sqrt)
}
