package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/alfredxing/calc/compute"
)

import (
	"golang.org/x/crypto/ssh/terminal"
)

// Stores the state of the terminal before making it raw
var regularState *terminal.State

func main() {
	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := compute.Evaluate(input)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
		return
	}

	var err error
	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, regularState)

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit()
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := compute.Evaluate(text)
		if err != nil {
			term.Write([]byte(fmt.Sprintln("Error: " + err.Error())))
			continue
		}
		term.Write([]byte(fmt.Sprintln(strconv.FormatFloat(res, 'G', -1, 64))))
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	terminal.Restore(0, regularState)
	fmt.Println()
	os.Exit(0)
}
