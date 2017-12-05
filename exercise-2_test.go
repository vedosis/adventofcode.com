package main

import "testing"

var ex2PackageTest = map[string]int{
	"2x3x4": 58,
	"1x1x10": 43,
}

var ex2RibbonTest = map[string]int{
	"2x3x4": 34,
	"1x1x10": 14,
}

func Test_ex2ParsePackageDimensions(t *testing.T){
	for key, value := range ex2PackageTest {
		box := ex2ParsePackageDimensions(key)
		result := box.PaperNeeded()
		if result != value {
			t.Error(key, value, result)
		}
	}
}

func Test_ex2ParseRibbonDimensions(t *testing.T){
	for key, value := range ex2RibbonTest {
		box := ex2ParsePackageDimensions(key)
		result := box.RibbonNeeded()
		if result != value {
			t.Error(key, value, result)
		}
	}
}