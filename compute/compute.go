package compute

import (
    "fmt"
    "strconv"
    "go/scanner"
    "go/token"
)

import (
    "../operators"
    "../operators/functions"
)

func Evaluate(in string) float64 {
    floats := NewFloatStack()
    ops := NewStringStack()
    s := initScanner(in)

    var prev token.Token = token.ILLEGAL

ScanLoop:
    for {
        _, tok, lit := s.Scan()
        switch {
        case tok == token.EOF:
            break ScanLoop
        case isOperand(tok):
            floats.Push(parseFloat(lit))
            if prev == token.RPAREN {
                evalUnprecedenced("*", ops, floats)
            }
        case functions.IsFunction(lit):
            if isOperand(prev) || prev == token.RPAREN {
                evalUnprecedenced("*", ops, floats)
            }
            ops.Push(lit)
        case isOperator(tok.String()):
            op := tok.String()
            if isNegation(tok, prev) {
                op = "neg"
            }
            evalUnprecedenced(op, ops, floats)
        case tok == token.LPAREN:
            if isOperand(prev) {
                evalUnprecedenced("*", ops, floats)
            }
            ops.Push(tok.String())
        case tok == token.RPAREN:
            for ops.Pos >= 0 && ops.Top() != "(" {
                evalOp(ops.Pop(), floats)
            }
            ops.Pop()
            if functions.IsFunction(ops.Top()) {
                evalOp(ops.Pop(), floats)
            }
        }
        prev = tok
    }

    fmt.Println(floats)
    fmt.Println(ops)

    for ops.Pos >= 0 {
        evalOp(ops.Pop(), floats)
    }

    return floats.Top()
}

func evalUnprecedenced(op string, ops *StringStack, floats *FloatStack) {
    for ops.Pos >= 0 && shouldPopNext(op, ops.Top()) {
        evalOp(ops.Pop(), floats)
    }
    ops.Push(op)
}

func shouldPopNext(n1 string, n2 string) bool {
    if !isOperator(n2) {
        return false
    }
    if n1 == "neg" {
        return false
    }
    op1 := parseOperator(n1)
    op2 := parseOperator(n2)
    if op1.Associativity == operators.L {
        return op1.Precedence <= op2.Precedence
    }
    return op1.Precedence < op2.Precedence
}

func evalOp(opName string, floats *FloatStack) {
    op := operators.FindOperatorFromString(opName)

    var args = make([]float64, op.Args)
    for i := op.Args - 1; i >= 0; i-- {
        args[i] = floats.Pop()
    }

    fmt.Printf("Computing %s of %q\n", opName, args)
    floats.Push(op.Operation(args))
}

func isOperand(tok token.Token) bool {
    return tok == token.FLOAT || tok == token.INT
}

func isOperator(lit string) bool {
    return operators.IsOperator(lit)
}

func isNegation(tok token.Token, prev token.Token) bool {
    return tok == token.SUB &&
        (prev == token.ILLEGAL || isOperator(prev.String()) || prev == token.LPAREN)
}

func parseFloat(lit string) float64 {
    f, err := strconv.ParseFloat(lit, 64)
    if err != nil {
        panic("Cannot parse recognized float: " + lit)
    }
    return f
}

func parseOperator(lit string) *operators.Operator {
    return operators.FindOperatorFromString(lit)
}

func initScanner(in string) scanner.Scanner {
    var s scanner.Scanner
    src := []byte(in)
    fset := token.NewFileSet()
    file := fset.AddFile("", fset.Base(), len(src))
    s.Init(file, src, nil, 0)
    return s
}

