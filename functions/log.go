package functions

import "math"

var (
    log = &Function{
        Name: "log",
        Call: func (val float64) float64 {
            return math.Log10(val)
        },
    }
    ln = &Function{
        Name: "ln",
        Call: func (val float64) float64 {
            return math.Log(val)
        },
    }
    lg = &Function{
        Name: "lg",
        Call: func (val float64) float64 {
            return math.Log2(val)
        },
    }
)

func init() {
    Register(log)
    Register(ln)
    Register(lg)
}
