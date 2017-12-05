package main

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func decrementX(input map[int]map[int]int, x int, y int) (int, int) {
	if input[x-1] == nil {
		input[x-1] = map[int]int{y: 1}
	} else if input[x-1][y] == 0 {
		input[x-1][y] = 1
	} else {
		input[x-1][y] = input[x-1][y] + 1
	}
	return x - 1, y
}

func incrementX(input map[int]map[int]int, x int, y int) (int, int) {
	if input[x+1] == nil {
		input[x+1] = map[int]int{y: 1}
	} else if input[x+1][y] == 0 {
		input[x+1][y] = 1
	} else {
		input[x+1][y] = input[x+1][y] + 1
	}
	return x + 1, y
}

func decrementY(input map[int]map[int]int, x int, y int) (int, int) {
	if input[x][y - 1] == 0 {
		input[x][y - 1] = 1
	} else {
		input[x][y - 1] = input[x][y - 1] + 1
	}
	return x, y - 1
}

func incrementY(input map[int]map[int]int, x int, y int) (int, int) {
	if input[x][y + 1] == 0 {
		input[x][y + 1] = 1
	} else {
		input[x][y + 1] = input[x][y + 1] + 1
	}
	return x, y + 1
}

func ex3ParseMatrixByNumActors(input string, numActors int) map[int]map[int]int {
	path := map[int]map[int]int{
		0: {0: numActors},
	}

	x := make([]int, numActors)
	y := make([]int, numActors)
	actor := 0
	for _, chr := range input {
		switch chr {
		case '<':
			x[actor], y[actor] = decrementX(path, x[actor], y[actor])
		case '>':
			x[actor], y[actor] = incrementX(path, x[actor], y[actor])
		case '^':
			x[actor], y[actor] = incrementY(path, x[actor], y[actor])
		case 'v':
			x[actor], y[actor] = decrementY(path, x[actor], y[actor])
		}

		actor++
		if actor >= numActors {
			actor = 0
		}
	}
	return path
}

func ex3ParseMatrix(input string) map[int]map[int]int {
	path := map[int]map[int]int{
		0: {0: 1},
	}
	var x, y int

	for _, chr := range input {
		switch chr {
		case '<':
			x, y = decrementX(path, x, y)
		case '>':
			x, y = incrementX(path, x, y)
		case '^':
			x, y = incrementY(path, x, y)
		case 'v':
			x, y = decrementY(path, x, y)
		default:
			panic(fmt.Sprintf("(%s) is not a known direction", chr))
		}
	}

	return path
}

func ex3CountHouses(input map[int]map[int]int) int {
	var totalHouses int

	for _, submap := range input {
		totalHouses = totalHouses + len(submap)
	}

	return totalHouses
}

func ex3Action(c *cli.Context) error {
	input, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-3.txt"))
	check(err)
	s := string(input)

	houseMatrix := ex3ParseMatrixByNumActors(s, 1)
	fmt.Printf("NumSantaHouses(%d)\n", ex3CountHouses(houseMatrix))

	houseMatrix = ex3ParseMatrixByNumActors(s, 2)
	fmt.Printf("NumSantaAndRobotHouses(%d)\n", ex3CountHouses(houseMatrix))

	return nil
}
