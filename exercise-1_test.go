package adventofcode_com

import (
	"testing"
)

var ParseTest = map[string]int{
	"(())":    0,
	"()()":    0,
	"(((":     3,
	"(()(()(": 3,
	"))(((((": 3,
	"())":     -1,
	"))(":     -1,
	")))":     -3,
	")())())": -3,
}

var PosTest = map[string]int {
	")": 1,
	"()())": 5,
}

func Test_ex1ParseFloor(t *testing.T) {
	for input, output := range ParseTest{
		result := ex1ParseFloor(input)
		if result != output{
			t.Error(input, output, result)
		}
	}
}

func Test_ex1ParsePosition(t *testing.T) {
	for input, output := range PosTest {
		result := ex1ParsePosition(input)
		if result != output {
			t.Error(input, output, result)
		}
	}
}