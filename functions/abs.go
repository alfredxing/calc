package functions

import "math"

var (
    abs = &Function{
        Name: "abs",
        Call: func (val float64) float64 {
            return math.Abs(val)
        },
    }
)

func init() {
    Register(abs)
}
