package main

import (
	"fmt"
	"testing"
)

func Test_ex6NumOnLights(t *testing.T) {
	var boolLights [1000][1000]bool
	var intLights [1000][1000]int

	boolLights[100][999] = true
	count := ex6NumOnLights(boolLights)
	if count != 1 {
		t.Error(1, count, "boolLights")
	}

	intLights[100][999] = 10
	intLights[999][100] = 2
	brightness := ex6Brightness(intLights)
	if brightness != 12 {
		t.Error(12, brightness, "intLights")
	}
}

type matchesStruct struct {
	X1          int
	X2          int
	Y1          int
	Y2          int
	Instruction string
}

var GetMatchesData = map[string]matchesStruct{
	"turn off 446,432 through 458,648": {446, 458, 432, 648, "turn off"},
	"turn on 715,871 through 722,890":  {715, 722, 871, 890, "turn on"},
	"toggle 424,675 through 740,862":   {424, 740, 675, 862, "toggle"},
}

func Test_ex6GetMatches(t *testing.T) {
	for input, value := range GetMatchesData {
		x1, x2, y1, y2, instruction := ex6GetMatches(input)
		if value.X1 != x1 || value.X2 != x2 || value.Y1 != y1 || value.Y2 != y2 || value.Instruction != instruction {
			t.Error(
				fmt.Sprintf(
					"Expected: (%d,%d)1, (%d,%d)2 and %s.\nReceived: (%d,%d)1, (%d,%d)2 and %s.\n",
					value.X1, value.Y1, value.X2, value.Y2, value.Instruction,
					x1, y1, x2, y2, instruction))
		}
	}

}
