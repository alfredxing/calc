package compute

import (
	"errors"
	"go/scanner"
	"go/token"
	"strconv"
	"strings"
)

import (
	"github.com/alfredxing/calc/constants"
	"github.com/alfredxing/calc/operators"
	"github.com/alfredxing/calc/operators/functions"
)

var resHistory = []float64{}

func Evaluate(in string) (float64, error) {
	floats := NewFloatStack()
	ops := NewStringStack()
	s := initScanner(in)

	var prev token.Token = token.ILLEGAL
	var back int = -1

ScanLoop:
	for {
		_, tok, lit := s.Scan()

		if lit != "@" && back > -1 && len(resHistory) > 0 {
			floats.Push(getHistory(back))
			if prev == token.RPAREN || constants.IsConstant(prev.String()) {
				evalUnprecedenced("*", ops, floats)
			}
			back = -1
		}

		switch {
		case tok == token.EOF:
			break ScanLoop
		case lit == "@":
			back += 1
		case constants.IsConstant(lit):
			floats.Push(constants.GetValue(lit))
			if prev == token.RPAREN || isOperand(prev) {
				evalUnprecedenced("*", ops, floats)
			}
		case isOperand(tok):
			val, err := parseFloat(lit)
			if err != nil {
				return 0, err
			}
			floats.Push(val)
			if prev == token.RPAREN || constants.IsConstant(prev.String()) {
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
			for ops.Pos >= 0 && ops.SafeTop() != "(" {
				err := evalOp(ops.SafePop(), floats)
				if err != nil {
					return 0, err
				}
			}
			_, err := ops.Pop()
			if err != nil {
				return 0, errors.New("Can't find matching parenthesis!")
			}
			if ops.Pos >= 0 {
				if functions.IsFunction(ops.SafeTop()) {
					err := evalOp(ops.SafePop(), floats)
					if err != nil {
						return 0, err
					}
				}
			}
		case tok == token.SEMICOLON:
		default:
			inspect := tok.String()
			if strings.TrimSpace(lit) != "" {
				inspect += " (`" + lit + "`)"
			}
			return 0, errors.New("Unrecognized token " + inspect + " in expression")
		}
		prev = tok
	}

	for ops.Pos >= 0 {
		op, _ := ops.Pop()
		err := evalOp(op, floats)
		if err != nil {
			return 0, err
		}
	}

	res, err := floats.Top()
	if err != nil {
		return 0, errors.New("Expression could not be parsed!")
	}
	pushHistory(res)
	return res, nil
}

func evalUnprecedenced(op string, ops *StringStack, floats *FloatStack) {
	for ops.Pos >= 0 && shouldPopNext(op, ops.SafeTop()) {
		evalOp(ops.SafePop(), floats)
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

func evalOp(opName string, floats *FloatStack) error {
	op := operators.FindOperatorFromString(opName)
	if op == nil {
		return errors.New("Either unmatched paren or unrecognized operator")
	}

	var args = make([]float64, op.Args)
	for i := op.Args - 1; i >= 0; i-- {
		arg, err := floats.Pop()
		if err != nil {
			return errors.New("Not enough arguments to operator!")
		}
		args[i] = arg
	}

	floats.Push(op.Operation(args))

	return nil
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

func parseFloat(lit string) (float64, error) {
	f, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		return 0, errors.New("Cannot parse recognized float: " + lit)
	}
	return f, nil
}

func parseOperator(lit string) *operators.Operator {
	return operators.FindOperatorFromString(lit)
}

func getHistory(back int) float64 {
	return resHistory[back]
}

func pushHistory(res float64) {
	resHistory = append([]float64{res}, resHistory...)
}

func initScanner(in string) scanner.Scanner {
	var s scanner.Scanner
	src := []byte(in)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	return s
}
