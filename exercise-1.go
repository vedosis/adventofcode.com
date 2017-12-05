package main

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func ex1ParseFloor(input string) int {

	position := 0

	for _, paren := range input {
		if paren == '(' {
			position++
		} else if paren == ')' {
			position--
		} else {
			panic(fmt.Sprintf("Unknown char '%s'", paren))
		}
	}

	return position
}

func ex1ParsePosition(input string) int {
	position := 0

	for pos, paren := range input {
		if paren == '(' {
			position++
		} else if paren == ')' {
			position--
		} else {
			panic(fmt.Sprintf("Unknown char '%s'", paren))
		}

		if position < 0 {
			return pos + 1
		}
	}
	return 0
}

func ex1Action(c *cli.Context) {
	c.String("dir")
	input, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-1.txt"))
	check(err)
	s := string(input)

	floor := ex1ParseFloor(s)
	fmt.Printf("Floor(%d)\n", floor)
	position := ex1ParsePosition(s)
	fmt.Printf("Position(%d)\n", position)
}
