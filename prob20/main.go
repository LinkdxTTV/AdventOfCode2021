package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	parsingImage := false
	image := [][]string{}
	var enhancementAlgorithm string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingImage = true
			continue
		}
		if !parsingImage {
			enhancementAlgorithm = line
		}
		if parsingImage {
			row := []string{}
			for _, char := range line {
				row = append(row, string(char))
			}
			image = append(image, row)
		}
	}

	fmt.Println(enhancementAlgorithm, len(enhancementAlgorithm))
	// Haha infinite
	for i := 0; i <= 55; i++ {
		image = expandImageByOne(image)
	}

	printImage(image)

	part1 := 0

	for iter := 0; iter < 50; iter++ {
		newImage := [][]string{}
		for y, row := range image {
			newRow := []string{}
			for x := range row {
				newRow = append(newRow, enhancePoint(x, y, iter, image, enhancementAlgorithm))
			}
			newImage = append(newImage, newRow)
		}
		image = newImage
		printImage(image)

		// Part 1
		if iter == 1 {
			part1 = countLights(image)
		}
	}

	part2 := countLights(image)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func countLights(image [][]string) int {
	count := 0
	for _, row := range image {
		for _, char := range row {
			if string(char) == "#" {
				count++
			}
		}
	}
	return count
}

func printImage(image [][]string) {
	for _, row := range image {
		fmt.Println(row)
	}
	fmt.Println()
}

func expandImageByOne(image [][]string) [][]string {
	maxX := len(image[0]) - 1
	output := [][]string{}
	firstRow := []string{}
	for i := 0; i <= maxX+2; i++ {
		firstRow = append(firstRow, ".")
	}
	output = append(output, firstRow)

	for _, row := range image {
		newRow := []string{"."}
		newRow = append(newRow, row...)
		newRow = append(newRow, ".")
		output = append(output, newRow)
	}

	lastRow := []string{}
	for i := 0; i <= maxX+2; i++ {
		lastRow = append(lastRow, ".")
	}
	output = append(output, lastRow)
	return output

}

// The infinite grid behavior of this function ONLY works if the periodicity of the flip is 2. I have not tested it on any advanced patterns
// Namely it matters that the first and last character in the enhancement algorithm string are opposites
// Technically we could approach some weird closed loop patterns here i.e. game of life
func enhancePoint(x, y int, iteration int, image [][]string, enhance string) string {
	binStr := ""
	yMax := len(image) - 1
	xMax := len(image[0]) - 1
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			xNew := x + i
			yNew := y + j
			if xNew < 0 || yNew < 0 || xNew > xMax || yNew > yMax {
				if iteration%2 == 0 { // Flip flop has period 2
					binStr += "."
				} else {
					binStr += enhance[0:1] // So this will only work on flip flop behavior
				}
				continue
			}
			binStr += image[yNew][xNew]
		}
	}
	// Convert binStr
	index := convertToDecimal(binStr)
	return enhance[index : index+1]
}

func convertToDecimal(input string) int {
	binary := ""
	for _, char := range input {
		if string(char) == "." {
			binary += "0"
		}
		if string(char) == "#" {
			binary += "1"
		}
	}

	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err, binary)
	}
	return int(num)
}
