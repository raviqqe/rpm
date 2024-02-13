package main

import (
	"strconv"

	"github.com/docopt/docopt-go"
)

func getArguments() (map[string]interface{}, error) {
	const usage = `rpm, the RPM reporter

Usage:
	rpm [-n <number>] [-c <number>] [-h] <url>

Options:
	-n <number>  Number of total requests. [default: 10000]
	-c <number>  Number of concurrent requests. [default: 1000]
	-h           Show this help.`

	args, err := docopt.ParseArgs(usage, nil, true, "", false)

	if err != nil {
		return nil, err
	}

	for _, s := range []string{"-n", "-c"} {
		args[s], err = parseInt(args[s].(string))

		if err != nil {
			return nil, err
		}
	}

	return args, nil
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(i), nil
}
