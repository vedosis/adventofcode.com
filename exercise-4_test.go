package main

import "testing"

var HashData = map[string]int{
	"abcdef": 609043,
	"pqrstuv": 1048970,
}

func Test_ex4FindFirstHashStartsWith(t *testing.T){
	for key, value := range HashData {
		result := ex4FindHashStartsWith(key, '0', 5)
		if result != value {
			t.Error(key, value, result)
		}
	}
}