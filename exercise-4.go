package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

func ex4FindHashStartsWith(salt string, char rune, num int) int {
	searchFor := strings.Repeat(string(char), num)
	var batch int
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(salt + strconv.Itoa(i)))
		hashString := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashString, searchFor){
			fmt.Print("\n")
			return i
		}

		batch++
		if batch >= 64 {
			fmt.Print(".")
			batch = 0
		}
	}
	fmt.Print("\n")
	return 0
}

func ex4Action(c *cli.Context) error {
	input, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-4.txt"))
	check(err)
	s := string(input)
	first := ex4FindHashStartsWith(s, '0', 5)
	fmt.Printf("FirstIncidenceOfFiveZeros(%d)", first)
	first = ex4FindHashStartsWith(s, '0', 6)
	fmt.Printf("FirstIncidenceOfSixZeros(%d)", first)

	return nil
}
