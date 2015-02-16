package compute

import (
    "testing"
    "math"
    "strconv"
)

var exps = map[string]float64{
    "1+1": 2,
    "1+2^3^2": 513,
    "2^(3+4)": 128,
    "2^(3/(1+2))": 2,
    "2(1+3)": 8,
    "2^2(1+3)": 16,
    "-2": -2,
    "1+(-1)^2": 2,
    "3-5": -2,
    "-3*4": -12,
    "3*-4": -12,
    "3*(3-(5+6)^12)*23^3-5^23": -126476703133661843,
    "2^3^2": 512,
    "-3^2": -9,
    "2(1+1)4": 16,
    "3*abs(1-5)": 12,
    "ln(3^15)": 16.47918433,
    "sqrt(10)": 3.16227766,
    "abs(-3/2)": 1.5,
    "1+2sin(-1024)tan(acos(1))^2": 1,
    "tan(10)cos(20)": 0.2645844,
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
