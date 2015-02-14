package functions

import "math"

var (
    sin = &Function{
        Name: "sin",
        Call: func (val float64) float64 {
            return math.Sin(val)
        },
    }
    cos = &Function{
        Name: "cos",
        Call: func (val float64) float64 {
            return math.Cos(val)
        },
    }
    tan = &Function{
        Name: "tan",
        Call: func (val float64) float64 {
            return math.Tan(val)
        },
    }
    asin = &Function{
        Name: "asin",
        Call: func (val float64) float64 {
            return math.Asin(val)
        },
    }
    acos = &Function{
        Name: "acos",
        Call: func (val float64) float64 {
            return math.Acos(val)
        },
    }
    atan = &Function{
        Name: "atan",
        Call: func (val float64) float64 {
            return math.Atan(val)
        },
    }
)

func init() {
    // We need to register the longer names first to avoid matching the shorter ones
    // Order is maintained in a separate function name list
    Register(asin)
    Register(acos)
    Register(atan)
    Register(sin)
    Register(cos)
    Register(tan)
}
