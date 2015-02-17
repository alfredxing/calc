package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/alfredxing/calc/compute"

	"golang.org/x/crypto/ssh/terminal"
)

// Stores the state of the terminal before making it raw
var regularState *terminal.State
var fd = int(os.Stdin.Fd())

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

	var err error
	// Raw mode behaves different on windows in this case.
	if runtime.GOOS != "windows" {
		regularState, err = terminal.MakeRaw(fd)
		if err != nil {
			panic(err)
		}
		defer terminal.Restore(fd, regularState)
	}

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				break
			}
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
		if regularState != nil {
			terminal.Restore(fd, regularState)
		}
		os.Exit(0)
	}
	return "", 0, false
}
