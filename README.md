## calc
A simple, fast, and intuitive command-line calculator written in Go.

### Introduction
A deviation from the org's usual web utilities, the idea for calc came up because I was frustrated with using the many alternatives; [each of them has their own shortcomings](#why-not).

### Install
Install calc as you would any other Go program:
```
go get github.com/wylst/calc
```

### Usage
You can use calc in two ways: command mode and shell mode.

#### Shell mode
This is probably the mode you'll want to use. It's like the `python` shell or `irb`. The shell mode uses the `terminal` package provided by [`golang.org/x/crypto/ssh/terminal`](https://godoc.org/golang.org/x/crypto/ssh/terminal), which means it supports many of the shell features you know and love (like history, pasting, and the `exit` command).
```shell
> 1+1
2
> 3(5/(3-4))
-15
> ln(-1)
NaN
```

### Why not ...?
