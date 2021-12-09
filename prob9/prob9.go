package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type xypoint struct {
	x int
	y int
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	rows := [][]int{}

	for scanner.Scan() {
		columns := []int{}
		row := scanner.Text()
		for _, letter := range row {
			num, err := strconv.ParseInt(string(letter), 10, 64)
			if err != nil {
				log.Fatal(err, num)
			}
			columns = append(columns, int(num))
		}
		rows = append(rows, columns)
	}

	lowPoints := []xypoint{}

	risks := []int{}
	for i, row := range rows {
		for j, num := range row {
			if isLocalMinimum(i, j, rows) {
				risks = append(risks, num+1)
				lowPoints = append(lowPoints, xypoint{
					x: i,
					y: j,
				})
			}
		}
	}

	total := 0
	for _, num := range risks {
		total += num
	}

	fmt.Println("Part 1:", total)

	// Part 2 recursion omegalul
	basinSizes := []int{}
	for _, point := range lowPoints {
		basinSize := recursiveClimb(point, rows, &map[xypoint]bool{})
		basinSizes = append(basinSizes, basinSize)
	}

	// golang sort OMEGALUL

	length := len(basinSizes)
	fmt.Println("Part 2: Largest Basins:", basinSizes[length-1], basinSizes[length-2], basinSizes[length-3])
	fmt.Println("Part 2: Product of Largest Basins:", basinSizes[length-1]*basinSizes[length-2]*basinSizes[length-3])
}

func isLocalMinimum(row, column int, mapping [][]int) bool {
	maxColumn := len(mapping[0]) - 1
	maxRow := len(mapping) - 1
	currentHeight := mapping[row][column]

	// Up
	if row-1 >= 0 {
		if currentHeight >= mapping[row-1][column] {
			return false
		}
	}
	// Down
	if row+1 <= maxRow {
		if currentHeight >= mapping[row+1][column] {
			return false
		}
	}

	// Left
	if column-1 >= 0 {
		if currentHeight >= mapping[row][column-1] {
			return false
		}
	}

	// Right
	if column+1 <= maxColumn {
		if currentHeight >= mapping[row][column+1] {
			return false
		}
	}

	return true
}

// recursiveClimb should take a point and return how many non-9 points are
// greater than itself, and adjacent
func recursiveClimb(point xypoint, mapping [][]int, dedupe *map[xypoint]bool) int {
	maxColumn := len(mapping[0]) - 1
	maxRow := len(mapping) - 1
	currentHeight := mapping[point.x][point.y]
	if currentHeight == 9 {
		return 0
	}
	dedupeMap := *dedupe
	_, ok := dedupeMap[point]
	if ok {
		return 0
	}

	totalHigherAdjacents := 0
	// Up
	if point.x-1 >= 0 {
		if currentHeight < mapping[point.x-1][point.y] {
			totalHigherAdjacents += recursiveClimb(xypoint{point.x - 1, point.y}, mapping, dedupe)
		}
	}
	// Down
	if point.x+1 <= maxRow {
		if currentHeight < mapping[point.x+1][point.y] {
			totalHigherAdjacents += recursiveClimb(xypoint{point.x + 1, point.y}, mapping, dedupe)
		}
	}

	// Left
	if point.y-1 >= 0 {
		if currentHeight < mapping[point.x][point.y-1] {
			totalHigherAdjacents += recursiveClimb(xypoint{point.x, point.y - 1}, mapping, dedupe)
		}
	}

	// Right
	if point.y+1 <= maxColumn {
		if currentHeight < mapping[point.x][point.y+1] {
			totalHigherAdjacents += recursiveClimb(xypoint{point.x, point.y + 1}, mapping, dedupe)
		}
	}

	dedupeMap[point] = true
	dedupe = &dedupeMap
	// Count self
	return totalHigherAdjacents + 1
}
