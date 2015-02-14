package functions

import (
    "strings"
    "strconv"
    "unicode"
    // "fmt"
)

var funcs = map[string]*Function{}
var names = []string{}

type Function struct {
    Name string
    Call func(val float64) float64
}

func Process(exp string) string {
    for _, name := range names {
        if strings.Contains(exp, name) {
            nums := strings.Split(exp, name)
            for i := 1; i < len(nums); i++ {
                val, rest := NextFloat(nums[i])
                replacement := strconv.FormatFloat(funcs[name].Call(val), 'f', -1, 64) + rest

                prev := nums[i-1]
                if len(prev) >= 1 {
                    prevChar := prev[len(prev)-1]
                    if unicode.IsNumber(rune(prevChar)) || prevChar == ')' {
                        replacement = "*" + replacement
                    }
                }

                nums[i] = replacement
            }
            exp = strings.Join(nums, "")
        }
    }
    return exp
}

func NextFloat(str string) (float64, string) {
    negative := strings.HasPrefix(str, "-")
    if negative {
        str = str[1:]
    }

    nextString := strings.FieldsFunc(str, func (r rune) bool {
        return !unicode.IsNumber(r) && r != '.'
        })[0]
    nextFloat, _ := strconv.ParseFloat(nextString, 64)

    if negative {
        nextFloat = 0 - nextFloat
    }

    return nextFloat, str[len(nextString):]
}

func Register(function *Function) {
    funcs[function.Name] = function
    names = append(names, function.Name)
}
