package main

import (
	"errors"
	"flag"
	"fmt"
)

type Args struct {
	Config string
}

func parseArgs() (Args, error) {
	args := Args{}
	config := flag.String("config", "./teo.json", "path to config file")

	flag.Parse()

	if len(flag.Args()) != 0 {
		return args, errors.New(
			fmt.Sprintf("extra positional argument", flag.Arg(0)),
		)
	}

	args.Config = *config

	return args, nil
}
