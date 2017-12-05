package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli")

func ex6NewImage(lights map[string][1000]int, file_path string, file_number int) {

}

func ex6CompileImages(file_path string) {

}

func ex6NumOnLights(lights map[string][1000]int) int {
	return 0
}

func ex6Action(c *cli.Context) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-5.txt"))
	check(err)
	defer file.Close()

	lights := map[string][1000]int{
		"X": {},
		"Y": {},
	}

	reader := bufio.NewReader(file)
	fileNumber := 0
	for {
		fileNumber++
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil && err != io.EOF {
			panic(err)
		}

		if c.Bool("with-gif"){
			ex6NewImage(lights, c.String("gif-dir"), fileNumber)
		}

		if err == io.EOF {
			break
		}

	}

	if c.Bool("with-gif"){
		ex6CompileImages(c.String("gif-dir"))
	}

	fmt.Printf("LightsOn(%d)\n", ex6NumOnLights(lights))
	return nil
}