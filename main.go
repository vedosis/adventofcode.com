package main

import (
	"os"

	"github.com/urfave/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Advent of Code 2015 ed"
	app.Usage = "Runs the exercises from adventofcode.com"

	file_path := cli.StringFlag{
		Name:  "dir",
		Value: "src/github.com/vedosis/adventofcode.com-2015/input",
		Usage: "directory for exercise input data",
	}

	app.Commands = []cli.Command{
		{
			Name:    "ex1",
			Aliases: []string{"1"},
			Usage:   "It's an elevator Santa... IT'S NOT THAT HARD",
			Action:  ex1Action,
			Flags:   []cli.Flag{file_path},
		},
		{
			Name:    "ex2",
			Aliases: []string{"2"},
			Usage:   "White boys wrapping",
			Action:  ex2Action,
			Flags:   []cli.Flag{file_path},
		},
		{
			Name:    "ex3",
			Aliases: []string{"3"},
			Usage:   "Santa and a Robot",
			Action:  ex3Action,
			Flags:   []cli.Flag{file_path},
		},
		{
			Name:    "ex4",
			Aliases: []string{"4"},
			Usage:   "AdventCoin Mining",
			Action:  ex4Action,
			Flags:   []cli.Flag{file_path},
		},
		{
			Name:    "ex5",
			Aliases: []string{"5"},
			Usage:   "Naughty or Nice Names",
			Action:  ex5Action,
			Flags:   []cli.Flag{file_path},
		},
	}

	app.Run(os.Args)
}
