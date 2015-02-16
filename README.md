## calc
A simple, fast, and intuitive command-line calculator written in Go.

### Introduction
A deviation from the org's usual web utilities, the idea for calc came up because I was frustrated with using [the alternatives](#why-not). Also because I wanted to learn Go!

### Install
Install calc as you would any other Go program:
```
go get github.com/wylst/calc
```

### Usage
You can use calc in two ways: shell mode and command.

#### Shell mode
This is probably the mode you'll want to use. It's like the `python` shell or `irb`. The shell mode uses the `terminal` package provided by [`golang.org/x/crypto/ssh/terminal`](https://godoc.org/golang.org/x/crypto/ssh/terminal), which means it supports many of the shell features you know and love (like history, pasting, and the `exit` command).
```shell
> 1+1
2
> 3(5/(3-4))
-15
> 3pi^2
29.608813203268074
> ln(-1)
NaN
```

#### Command
You can also use calc to evaluate an expression with just a single command (i.e. without opening the shell). To do this, just use `calc [expression]`:
```shell
bash$ calc 1+1
2
bash$
```

### Supported functions, operators, and constants
calc supports all the standard stuff, and I'm definitely adding more later (also feel free to fork and add your own!)

##### Operators
`+`, `-`, `*`, `/`, `^`

##### Functions
`sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `sqrt`, `log`, `lg`, `ln`, `abs`

##### Constants
`e`, `pi`, `π`

### Why not use ...?
- Google
  - Doesn't work without an internet connection
  - Slower
  - Doesn't show previous computations, so you end up with multiple tabs open at once.
- Spotlight (on OS X)
  - No history
  - Switching between Spotlight and other windows isn't too fun
- Python/IRB
  - Requires use of a separate math module for most functions and constants
  - A little bit slower to start up
- `bc`
  - Limited number of built-in functions; these have shortened (not too intuitive) names as well.

The alternatives above are all great, and have their own advantages over calc. I highly recommend looking into these if you don't like how calc works.
