package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var globalRiskMin int = math.MaxInt32

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
		lineInt := []int{}
		for _, char := range line {
			num, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			lineInt = append(lineInt, int(num))
		}
		grid = append(grid, lineInt)
	}

	// Start at the top
	fromStart := map[xypoint]int{}
	move(pointRisk{xypoint{0, 0}, -grid[0][0]}, grid, fromStart)
	fmt.Println("Part 1:", globalRiskMin)

	// Part 2
	fromStart = map[xypoint]int{}
	globalRiskMin = math.MaxInt32

	// Make the map bigger
	outputGrid := [][]int{}
	for j := 0; j < 5; j++ {
		for _, row := range grid {
			outputRow := []int{}
			for i := 0; i < 5; i++ {
				for _, num := range row {
					outputRow = append(outputRow, num+i+j)
				}
			}
			for i, num := range outputRow {
				if num > 9 {
					outputRow[i] = num - 9
				}
			}
			outputGrid = append(outputGrid, outputRow)
		}
	}
	move(pointRisk{xypoint{0, 0}, -outputGrid[0][0]}, outputGrid, fromStart)
	fmt.Println("Part 2:", globalRiskMin)
}

type xypoint struct {
	x int
	y int
}

type pointRisk struct {
	point xypoint
	risk  int
}

func move(start pointRisk, grid [][]int, lowestRiskPerPoint map[xypoint]int) {
	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1
	var nextPoints = []pointRisk{start}
	for {
		if len(nextPoints) == 0 {
			break
		}
		point := nextPoints[0]
		nextPoints = nextPoints[1:]

		currentRisk := point.risk + grid[point.point.y][point.point.x]
		risk, ok := lowestRiskPerPoint[point.point]
		if ok && currentRisk >= risk {
			continue
		}
		if currentRisk > globalRiskMin {
			continue
		}
		lowestRiskPerPoint[point.point] = currentRisk
		if point.point.x == xMax && point.point.y == yMax {
			if currentRisk < globalRiskMin {
				globalRiskMin = currentRisk
			}
			continue
		}

		if point.point.x < xMax {
			nextPoints = append(nextPoints, pointRisk{xypoint{point.point.x + 1, point.point.y}, currentRisk})
		}
		if point.point.y < yMax {
			nextPoints = append(nextPoints, pointRisk{xypoint{point.point.x, point.point.y + 1}, currentRisk})
		}
		if point.point.x > 0 {
			nextPoints = append(nextPoints, pointRisk{xypoint{point.point.x - 1, point.point.y}, currentRisk})
		}
		if point.point.y > 0 {
			nextPoints = append(nextPoints, pointRisk{xypoint{point.point.x, point.point.y - 1}, currentRisk})
		}
	}
}
