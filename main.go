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
			Usage:   "Run the first exercise: http://adventofcode.com/2015/day/1",
			Action:  ex1Action,
			Flags: []cli.Flag{ file_path },
		},
		{
			Name:    "ex2",
			Aliases: []string{"2"},
			Usage:   "Run the second exercise: http://adventofcode.com/2015/day/2",
			Action:  ex2Action,
			Flags: []cli.Flag{ file_path },
		},
		{
			Name:    "ex3",
			Aliases: []string{"3"},
			Usage:   "Numba 3",
			Action:  ex3Action,
			Flags: []cli.Flag{ file_path },
		},
	}

	app.Run(os.Args)
}
