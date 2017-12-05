package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

var ThreeVowels = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
var NoBadPairs = regexp.MustCompile(`(ab|cd|pq|xy)`)

func ex5NiceStringCheck(str string) bool {
	if !ThreeVowels.MatchString(str){
		return false
	}

	if !ex5DoubleCheck(str){
		return false
	}

	if NoBadPairs.MatchString(str){
		return false
	}

	return true
}

func ex5NicerStringCheck(str string) bool {
	if !ex5DoublePair(str, 2){
		return false
	}

	if !ex5HasNeighbor(str){
		return false
	}
	return true
}

func ex5HasNeighbor(str string) bool {
	strLength := len(str)
	for i := 0; i <= strLength-3; i++ {
		if str[i] == str[i+2]{
			return true
		}
	}
	return false
}

func ex5DoublePair(str string, length int) bool {
	strLength := len(str)
	for i := 0; i <= strLength-length; i++ {
		if strings.Contains(fmt.Sprintf("%s**%s", str[0:i], str[i+length:]), str[i:i+length]){
			return true
		}
	}
	return false
}

func ex5DoubleCheck(str string) bool {
	var current rune
	for _, char := range str {
		if current == char {
			return true
		}
		current = char
	}
	return false
}


func ex5Action(c *cli.Context) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-5.txt"))
	check(err)
	defer file.Close()

	var niceStrings int
	var nicerStrings int

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil && err != io.EOF {
			panic(err)
		}

		if ex5NiceStringCheck(line) {
			niceStrings++
		}

		if ex5NicerStringCheck(line) {
			nicerStrings++
		}

		if err == io.EOF {
			break
		}

	}

	fmt.Printf("NiceNames(%d)\n", niceStrings)
	fmt.Printf("NicerNames(%d)\n", nicerStrings)
	return nil
}
