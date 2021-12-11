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

	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, char := range line {
			num, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, int(num))
		}
		grid = append(grid, row)
	}

	part2grid := deepCopyGrid(grid)

	flashes := 0

	// Loop 100 times
	for loop := 0; loop < 100; loop++ {
		for i, row := range grid {
			for j := range row {
				increaseAndRecursivelyFlash(i, j, grid, &flashes)
			}
		}
		cleanGrid(grid)
	}

	fmt.Println("Part 1:", flashes)

	// Part 2
	loopNum := 0
	grid = part2grid
	loop := 0
	for {
		for i, row := range grid {
			for j := range row {
				increaseAndRecursivelyFlash(i, j, grid, &flashes)
			}
		}
		cleanGrid(grid)

		if isGridAllZeros(grid) {
			loopNum = loop + 1
			break
		}
		loop++
	}
	fmt.Println("Part 2:", loopNum)
}

func deepCopyGrid(grid [][]int) [][]int {
	output := [][]int{}
	for _, row := range grid {
		outrow := []int{}
		for _, num := range row {
			outrow = append(outrow, num)
		}
		output = append(output, outrow)
	}
	return output
}

func increaseAndRecursivelyFlash(x, y int, grid [][]int, flashCounter *int) {
	rowMax := len(grid) - 1
	columnMax := len(grid[0]) - 1

	grid[x][y] += 1
	if grid[x][y] == 10 {
		*flashCounter += 1
		for xvar := -1; xvar <= 1; xvar++ {
			for yvar := -1; yvar <= 1; yvar++ {
				newX := x + xvar
				newY := y + yvar
				if newX < 0 || newX > rowMax || newY < 0 || newY > columnMax || (xvar == 0 && yvar == 0) {
					continue
				}
				increaseAndRecursivelyFlash(newX, newY, grid, flashCounter)
			}
		}
	}
}

func cleanGrid(grid [][]int) {
	for i, row := range grid {
		for j, num := range row {
			if num >= 10 {
				grid[i][j] = 0
			}
		}
	}
}

func isGridAllZeros(grid [][]int) bool {
	for _, row := range grid {
		for _, num := range row {
			if num != 0 {
				return false
			}
		}
	}
	return true
}
