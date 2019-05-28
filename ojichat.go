package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/greymd/ojichat/generator"
)

var appVersion = "0.0.1"
var usage = `Usage:
  ojichat [options] [<name>]

Options:
  -h --help                Show this screen.
  --version                Show version.
  -e <num>, --emoji=<num>  Maximum number of continuous Emojis [default: 3].`

func main() {
	parser := &docopt.Parser{OptionsFirst: true}
	args, _ := parser.ParseArgs(usage, nil, appVersion)
	config := generator.Config{}
	args.Bind(&config)

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}