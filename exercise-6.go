package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var RegexParseString = regexp.MustCompile(`^(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func ex6NewImage(lights [1000][1000]bool, palette []color.Color) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, 1000, 1000), palette)
	colorMap := map[bool]color.Color{
		false: palette[0],
		true:  palette[1],
	}
	for x, columns := range lights {
		for y, illuminated := range columns {
			img.Set(x, y, colorMap[illuminated])
		}
	}
	return img
}

func ex6NewOpacityImage(lights [1000][1000]int, palette []color.Color) *image.Paletted {
	img := image.NewPaletted(image.Rect(0,0,1000,1000), palette)
	for x, columns := range lights {
		for y, value := range columns {
			img.Set(x,y, color.RGBA{uint8(value), uint8(value), uint8(value), 0xff})
		}
	}
	return img
}

func ex6CompileImages(filePath string, prefix string, images []*image.Paletted) string {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.Mkdir(filePath, 0700)
		if err != nil {
			panic(err)
		}
	}

	delays := make([]int, len(images))
	timestamp := time.Now().Format("20060102-150405")
	fileName := fmt.Sprintf("%s/ex6-%s-%s.gif", filePath, prefix, timestamp)
	f, err := os.Create(fileName)
	check(err)

	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	check(err)

	return fileName
}

func ex6NumOnLights(lights [1000][1000]bool) int {
	var numOn int
	for _, column := range lights {
		for _, illuminated := range column {
			if illuminated {
				numOn++
			}
		}
	}
	return numOn
}

func ex6Brightness(lights [1000][1000]int) int {
	var brightness int
	for _, column := range lights {
		for _, value := range column {
			brightness += value
		}
	}
	return brightness
}

func ex6GetMatches(input string) (int, int, int, int, string) {
	matches := RegexParseString.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return 0, 0, 0, 0, ""
	}

	x1, err := strconv.Atoi(matches[0][2])
	check(err)
	x2, err := strconv.Atoi(matches[0][4])
	check(err)
	y1, err := strconv.Atoi(matches[0][3])
	check(err)
	y2, err := strconv.Atoi(matches[0][5])
	return x1, x2, y1, y2, matches[0][1]
}

func ex6AdjustImage(lights *[1000][1000]bool, input string) {
	x1, x2, y1, y2, instruction := ex6GetMatches(input)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch instruction {
			case "turn on":
				lights[x][y] = true
			case "turn off":
				lights[x][y] = false
			case "toggle":
				lights[x][y] = !lights[x][y]
			}
		}
	}
}

func ex6AdjustOpacity(lights *[1000][1000]int, input string) {
	x1, x2, y1, y2, instruction := ex6GetMatches(input)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch instruction {
			case "turn on":
				lights[x][y]++
			case "turn off":
				lights[x][y]--
				if lights[x][y] < 0 {
					lights[x][y] = 0
				}
			case "toggle":
				lights[x][y] += 2
			}
		}
	}
}

func ex6Action(c *cli.Context) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", c.String("dir"), "exercise-6.txt"))
	check(err)
	defer file.Close()

	var lights1 [1000][1000]bool
	var lights2 [1000][1000]int
	var images1 []*image.Paletted
	var images2 []*image.Paletted

	palette1 := []color.Color{
		color.RGBA{A: 0xff},
		color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	}
	var palette2 []color.Color
	for i := 0; i <= 255; i++ {
		palette2 = append(palette2, color.RGBA{R: uint8(i), G: uint8(i), B: uint8(i)})
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil && err != io.EOF {
			panic(err)
		}

		ex6AdjustImage(&lights1, line)
		ex6AdjustOpacity(&lights2, line)

		if c.Bool("with-gif") {
			fmt.Println("Making new image for " + line)
			images1 = append(images1, ex6NewImage(lights1, palette1))
			images2 = append(images2, ex6NewOpacityImage(lights2, palette2))
		}

		if err == io.EOF {
			break
		}

	}

	if c.Bool("with-gif") {
		filePath := ex6CompileImages(c.String("gif-dir"), "part1", images1)
		fmt.Printf("GifTransitionsFile(%s)\n", filePath)
		filePath = ex6CompileImages(c.String("gif-dir"), "part2", images2)
		fmt.Printf("GifTransisiotsFile2(%s)\n", filePath)
	}

	fmt.Printf("LightsOn(%d)\n", ex6NumOnLights(lights1))
	fmt.Printf("TotalBrightness(%d)\n", ex6Brightness(lights2))
	return nil
}
