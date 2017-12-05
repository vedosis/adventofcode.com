package main

import "testing"

var NameData = map[string]bool {
	"ugknbfddgicrmopn": true,
	"aaa": true,
	"jchzalrnumimnmhp": false,
	"haegwjzuvuyypxyu": false,
	"dvszwmarrgswjxmb": false,
}

var NicerNameData = map[string]bool {
	"qjhvhtzxzqqjkmpb": true,
	"xxyxx": true,
	"uurcxstgmygtbstg": false,
	"ieodomkazucvgmuy": false,
}

func Test_ex5NiceStringCheck(t *testing.T){
	for key, value := range NameData {
		isNice := ex5NiceStringCheck(key)
		if isNice != value {
			t.Error(key, value, isNice)
		}
	}
}

func Test_ex5NicerStringCheck(t *testing.T){
	for key, value := range NicerNameData {
		isNice := ex5NicerStringCheck(key)
		if isNice != value {
			t.Error(key, value, isNice)
		}
	}
}