package operators

var Ops = map[string]*Operator{}
var Names = []string{}

const (
    L = 0
    R = 1
)

type Operator struct {
    Name string
    Precedence int
    Associativity int
    Args int
    Operation func(args []float64) float64
}

func Register(op *Operator) {
    Ops[op.Name] = op
    Names = append(Names, op.Name)
}

func IsOperator(str string) bool {
    for _, name := range Names {
        if str == name {
            return true
        }
    }
    return false
}

func FindOperatorFromString(str string) *Operator {
    if IsOperator(str) {
        return Ops[str]
    }
    return nil
}
