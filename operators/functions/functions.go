package functions

import (
    "../"
)

var Names = []string{}

func Register(op *operators.Operator) {
    operators.Register(op)
    Names = append(Names, op.Name)
}

func IsFunction(str string) bool {
    for _, name := range Names {
        if str == name {
            return true
        }
    }
    return false
}
