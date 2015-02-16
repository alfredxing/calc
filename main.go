package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

import (
	"github.com/alfredxing/calc/compute"
)

import (
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := compute.Evaluate(input)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
		return
	}

	if runtime.GOOS != "windows" {
		oldState, err := terminal.MakeRaw(0)
		if err != nil {
			panic(err)
		}
		defer terminal.Restore(0, oldState)
	}

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	for {
		text, err := term.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := compute.Evaluate(text)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			continue
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		fmt.Println()
		os.Exit(0)
	}
	return "", 0, false
}
