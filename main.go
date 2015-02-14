package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

import (
    "github.com/wylst/calc.go/compute"
)

import (
    "golang.org/x/crypto/ssh/terminal"
)

func main() {
    if len(os.Args) > 1 {
        input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
        fmt.Printf("%s\n", strconv.FormatFloat(compute.Evaluate(input), 'G', -1, 64))
        return
    }

    oldState, err := terminal.MakeRaw(0)
    if err != nil {
        panic(err)
    }
    defer terminal.Restore(0, oldState)

    term := terminal.NewTerminal(os.Stdin, "> ")
    term.AutoCompleteCallback = handleKey
    for {
        text, err := term.ReadLine()
        if err != nil {
            panic(err)
        }

        text = strings.Replace(text, " ", "", -1)
        if text == "exit" || text == "quit" {
            break
        }
        fmt.Printf("%s\n", strconv.FormatFloat(compute.Evaluate(text), 'G', -1, 64))
    }
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
    if key == '\x03' {
        fmt.Println()
        os.Exit(0)
    }
    return "", 0, false
}
