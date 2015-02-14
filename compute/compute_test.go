package compute

import (
    "testing"
    "math"
    "strconv"
)

var exps = map[string]float64{
    "3*(3-(5+6)^12)*23^3-5^23": -126476703133661843,
    "2^3^2": 512,
    "-3^2": -9,
    "2(1+1)4": 16,
    "1+2sin(-1024)tan(acos(1))^2": 1,
    "ln(3^15)": 16.47918433,
    "sqrt(10)": 3.16227766,
    "abs(-3/2)": 1.5,
}
const DELTA = 0.000001

func TestEvaluate(t *testing.T) {
    for expression, expected := range exps {
        res := Evaluate(expression)
        if math.Abs(res - expected) > DELTA {
            message := expression + " failed: actual value " +
                strconv.FormatFloat(res, 'G', -1, 64) +
                " differs from expected value " +
                strconv.FormatFloat(expected, 'G', -1, 64)
            t.Error(message)
        }
    }
}
