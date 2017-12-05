package main

import (
	"testing"
)

var ex3ParseMatrixData = map[string]int{
	">": 2,
	"^>v<": 4,
	"^v^v^v^v^v": 2,
}

func Test_ex3ParseMatrix(t *testing.T) {
	for key, value := range ex3ParseMatrixData{
		result := ex3ParseMatrix(key)
		houseCount := ex3CountHouses(result)
		if houseCount != value {
			t.Error(key, value, result)
		}
	}

}
