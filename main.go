package adventofcode_com

import (
	"os"

	"github.com/urfave/cli"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	app := cli.NewApp()
	app.Name = "Advent of Code 2015 ed"
	app.Usage = "Runs the exercises from adventofcode.com"
	app.Commands = []cli.Command{
		{
			Name: "ex1",
			Aliases: []string{"1"},
			Usage: "Run the first exercise: http://adventofcode.com/2015/day/1",
			Action: ex1Action,
		},
		{
			Name: "ex2",
			Aliases: []string{"2"},
			Usage: "Run the second exercise: http://adventofcode.com/2015/day/2",
			Action: ex2Action,
		},
	}

	app.Run(os.Args)
}
