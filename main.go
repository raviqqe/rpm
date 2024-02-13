package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	args, err := getArguments()

	if err != nil {
		fail(err.Error())
	}

	a := newAttacker()
	r := newRPMReporter()
	e := newErrorReporter()

	go e.Analyze()
	go r.Analyze()

	a.Attack(args["<url>"].(string), args["-n"].(int), args["-c"].(int), r.Successes(), e.Errors())

	fmt.Println(r.Report())

	if s := e.Report(); s != "" {
		fail(color.RedString(s))
	}
}

func fail(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}
