package adventofcode_com

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type Box struct {
	Length int
	Width  int
	Height int
}

func (b Box) LW() int {
	return b.Length * b.Width
}

func (b Box) WH() int {
	return b.Width * b.Height
}

func (b Box) HL() int {
	return b.Height * b.Length
}

func (b Box) Area() int {
	return (b.LW() + b.WH() + b.HL()) * 2
}

func (b Box) PaperNeeded() int {
	dimensions := []int{b.Length, b.Height, b.Width}
	sort.Ints(dimensions)
	extraPadding := dimensions[0] * dimensions[1]
	return extraPadding + b.Area()
}

func (b Box) RibbonNeeded() int {
	dimensions := []int{b.Length, b.Height, b.Width}
	sort.Ints(dimensions)
	length := 2 * (dimensions[0] + dimensions[1])
	length = length + (dimensions[0] * dimensions[1] * dimensions[2])
	return length
}

func ex2ParsePackageDimensions(input string) *Box {
	re := regexp.MustCompile(`^(\d+)[xX](\d+)[xX](\d+)$`)
	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) != 1 {
		panic(fmt.Sprintf("%s doesn't match regexp", input))
	}

	length, _ := strconv.Atoi(matches[0][1])
	width, _ := strconv.Atoi(matches[0][2])
	height, _ := strconv.Atoi(matches[0][3])

	b := Box{length, width, height}
	return &b
}

func ex2Action(c *cli.Context) error {
	file, err := os.Open("input/2015/exercise-2.txt")
	check(err)
	defer file.Close()

	requiredSqFt := 0
	requiredRibbon := 0

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil && err != io.EOF{
			panic(err)
		}

		box := ex2ParsePackageDimensions(line)
		requiredSqFt = requiredSqFt + box.PaperNeeded()
		requiredRibbon = requiredRibbon + box.RibbonNeeded()

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("MinimumPaperSqFt(%d)\n", requiredSqFt)
	fmt.Printf("MinimumRibbonLnFt(%d)\n", requiredRibbon)

	return nil
}
