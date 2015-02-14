package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
    "strconv"
    "math"
    "errors"
    "./functions"
)

func main() {
    if len(os.Args) > 1 {
        input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
        fmt.Printf("%s\n", strconv.FormatFloat(parens(input), 'G', -1, 64))
        return
    }

    fmt.Print("> ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text := strings.Replace(scanner.Text(), " ", "", -1)

        if text == "exit" || text == "quit" {
            return
        }

        fmt.Printf("%s\n", strconv.FormatFloat(parens(text), 'G', -1, 64))
        fmt.Print("> ")
    }
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("3*(3-(5+6)^12)*23^3-5^23"), 'G', -1, 64))
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("2^2^3^2"), 'G', -1, 64))
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("-3^2"), 'G', -1, 64))
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("2(1+1)4"), 'G', -1, 64))
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("1+2sin(-1024)tan(acos(1))^2"), 'G', -1, 64))
    // fmt.Printf("%s\n", strconv.FormatFloat(parens("lg(log(1e16))"), 'G', -1, 64))
}

func parens(exp string) float64 {
    for i := 0; i < len(exp); i++ {
        if exp[i] == '(' {
            var inner string
            var j int
            open := 1
            for j = i+1; open > 0; j++ {
                if exp[j] == ')' {
                    open--
                } else if exp[j] == '(' {
                    open++
                }
                inner += string(exp[j])
            }
            res := parens(inner[:len(inner)-1])
            insert := strconv.FormatFloat(res, 'f', -1, 64)

            if i > 0 && unicode.IsNumber(rune(exp[i-1])) {
                insert = "*" + insert
            }
            if j < len(exp) && unicode.IsNumber(rune(exp[j])) {
                insert = insert + "*"
            }

            exp = exp[0:i] + insert + exp[j:]
        }
    }

    exp = functions.Process(exp)

    res, err := compute(exp, 0)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return res
}

func compute(exp string, level int) (float64, error) {
    if level == 0 && strings.ContainsAny(exp, "+-") {
        nums, ops := args(exp, []rune{'+', '-'})

        if len(ops) == len(nums) && strings.HasPrefix(exp, "-") {
            nums = append([]string{"0"}, nums...)
        }

        for i, num := range nums {
            if !unicode.IsNumber(rune(num[len(num) - 1])) {
                newNum := nums[i] + ops[i] + nums[i+1]
                nums = append(append(nums[:i], newNum), nums[i+2:]...)
                ops = append(ops[:i], ops[i+1:]...)
            }
        }

        if len(ops) >= len(nums) {
            return 0, errors.New("Not evaluatable: " + exp)
        }

        if len(nums) == 1 {
            res, err := compute(nums[0], 1)
            if err != nil {
                return 0, err
            }
            nums[0] = strconv.FormatFloat(res, 'f', -1, 64)
        }

        res, err := subcompute(nums, ops, 1, false)
        if err != nil {
            return 0, err
        }
        return res, nil
    } else if level <= 1 && strings.ContainsAny(exp, "*/") {
        nums, ops := args(exp, []rune{'*', '/'})

        if len(ops) >= len(nums) {
            return 0, errors.New("Not evaluatable: " + exp)
        }

        res, err := subcompute(nums, ops, 2, false)
        if err != nil {
            return 0, err
        }
        return res, nil
    } else if level <= 2 && strings.Contains(exp, "^") {
        nums, ops := args(exp, []rune{'^'})

        if len(ops) >= len(nums) {
            return 0, errors.New("Not evaluatable: " + exp)
        }

        res, err := subcompute(nums, ops, 3, true)
        if err != nil {
            return 0, err
        }
        return res, nil
    } else {
        res, err := strconv.ParseFloat(exp, 64)
        return res, err
    }
}

func args(exp string, symbols []rune) ([]string, []string) {
    nums := strings.FieldsFunc(exp, func (r rune) bool {
        for _, symbol := range symbols {
            if r == symbol {
                return true
            }
        }
        return false
    })
    ops := strings.FieldsFunc(exp, func (r rune) bool {
        for _, symbol := range symbols {
            if r == symbol {
                return false
            }
        }
        return true
    })
    return nums, ops
}

func subcompute(nums []string, ops []string, level int, reverse bool) (float64, error) {
    for len(nums) > 1 {
        var op string
        var n1, n2 float64
        var err1, err2 error
        if !reverse {
            op = ops[0]
            n1, err1 = compute(nums[0], level)
            n2, err2 = compute(nums[1], level)
        } else {
            op = ops[len(ops)-1]
            n1, err1 = compute(nums[len(nums)-2], level)
            n2, err2 = compute(nums[len(nums)-1], level)
        }

        if err1 != nil {
            return 0, errors.New(err1.Error())
        }
        if err2 != nil {
            return 0, errors.New(err2.Error())
        }

        var res float64
        switch op {
        case "^":
            res = math.Pow(n1, n2)
        case "*":
                res = n1 * n2
        case "/":
            res = n1 / n2
        case "+":
                res = n1 + n2
        case "-":
            res = n1 - n2
        }

        if reverse {
            nums[len(nums)-2] = strconv.FormatFloat(res, 'f', -1, 64)
            nums = nums[0:len(nums)-1]
            ops = ops[0:len(ops)-1]
        } else {
            nums[1] = strconv.FormatFloat(res, 'f', -1, 64)
            nums = append(nums[:0], nums[1:]...)
            ops = append(ops[:0], ops[1:]...)
        }
    }
    res, _ := strconv.ParseFloat(nums[0], 64)
    return res, nil
}

func Print(a interface{}) {
    fmt.Printf("%q\n", a)
}
