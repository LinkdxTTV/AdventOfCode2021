package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type fold struct {
	dimension string
	value     int
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	foldScan := false
	points := []point{}
	folds := []fold{}

	// For making the grid
	maxX := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			foldScan = true
			continue
		}
		if !foldScan {
			nums := strings.Split(line, ",")
			numX, err := strconv.ParseInt(nums[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numY, err := strconv.ParseInt(nums[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if int(numX) > maxX {
				maxX = int(numX)
			}
			if int(numY) > maxY {
				maxY = int(numY)
			}

			points = append(points, point{x: int(numX), y: int(numY)})
		} else {
			foldSplit := strings.Split(line, " ")
			dimSplit := strings.Split(foldSplit[2], "=")
			value, err := strconv.ParseInt(dimSplit[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, fold{dimension: dimSplit[0], value: int(value)})
		}
	}

	// Create grid
	grid := makeGrid(maxX, maxY)

	// Fill in points
	for _, point := range points {
		grid[point.y][point.x] = "#"
	}

	for i, fold := range folds {
		if i == 1 {
			// part 1 only wants 1 fold
			dots := countDots(grid)
			fmt.Println("Part 1:", dots)
		}
		grid = foldGrid(fold, grid)
	}

	// Part 2
	fmt.Println("Part 2:")
	printGrid(grid)

}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func foldGrid(fold fold, grid [][]string) [][]string {
	var maxX, maxY int
	var outGrid [][]string
	if fold.dimension == "y" {
		maxX = len(grid[0])
		maxY = fold.value
	} else {
		maxX = fold.value
		maxY = len(grid)
	}
	outGrid = makeGrid(maxX, maxY)

	for y, row := range grid {
		for x, char := range row {
			if string(char) == "#" {
				if x > maxX {
					x = maxX - (x - maxX)
				}
				if y > maxY {
					y = maxY - (y - maxY)
				}
				outGrid[y][x] = "#"
			}
		}
	}
	return outGrid
}

func makeGrid(maxX, maxY int) [][]string {
	grid := [][]string{}
	for j := 0; j <= maxY; j++ {
		row := []string{}
		for i := 0; i <= maxX; i++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}
	return grid
}

func countDots(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, char := range row {
			if string(char) == "#" {
				count++
			}
		}
	}
	return count
}
