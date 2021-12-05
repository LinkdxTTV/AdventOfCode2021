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

type coordPair struct {
	start point
	end   point
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	coordPairs := []coordPair{}

	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), " -> ")
		coordPairs = append(coordPairs, coordPair{
			start: pointFromString(coordinates[0]),
			end:   pointFromString(coordinates[1]),
		})
	}

	hitSet := map[point]int{}
	for _, pair := range coordPairs {
		if pair.start.x == pair.end.x {
			var ybig, ysmall int
			if pair.start.y < pair.end.y {
				ysmall = pair.start.y
				ybig = pair.end.y
			} else {
				ysmall = pair.end.y
				ybig = pair.start.y
			}
			for i := ysmall; i <= ybig; i++ {
				pointHere := point{
					x: pair.start.x,
					y: i,
				}
				_, ok := hitSet[pointHere]
				if !ok {
					hitSet[pointHere] = 1
				} else {
					hitSet[pointHere]++
				}
			}
		} else if pair.start.y == pair.end.y {
			var xbig, xsmall int
			if pair.start.x < pair.end.x {
				xsmall = pair.start.x
				xbig = pair.end.x
			} else {
				xsmall = pair.end.x
				xbig = pair.start.x
			}
			for i := xsmall; i <= xbig; i++ {
				pointHere := point{
					x: i,
					y: pair.start.y,
				}
				_, ok := hitSet[pointHere]
				if !ok {
					hitSet[pointHere] = 1
				} else {
					hitSet[pointHere]++
				}
			}
		}
	}
	counter := 0
	for _, val := range hitSet {
		if val > 1 {
			counter++
		}
	}
	fmt.Println(counter)

	// Part 2
	hitSet = map[point]int{}
	for _, pair := range coordPairs {
		if pair.start.x == pair.end.x {
			var ybig, ysmall int
			if pair.start.y < pair.end.y {
				ysmall = pair.start.y
				ybig = pair.end.y
			} else {
				ysmall = pair.end.y
				ybig = pair.start.y
			}
			for i := ysmall; i <= ybig; i++ {
				pointHere := point{
					x: pair.start.x,
					y: i,
				}
				_, ok := hitSet[pointHere]
				if !ok {
					hitSet[pointHere] = 1
				} else {
					hitSet[pointHere]++
				}
			}
		} else if pair.start.y == pair.end.y {
			var xbig, xsmall int
			if pair.start.x < pair.end.x {
				xsmall = pair.start.x
				xbig = pair.end.x
			} else {
				xsmall = pair.end.x
				xbig = pair.start.x
			}
			for i := xsmall; i <= xbig; i++ {
				pointHere := point{
					x: i,
					y: pair.start.y,
				}
				_, ok := hitSet[pointHere]
				if !ok {
					hitSet[pointHere] = 1
				} else {
					hitSet[pointHere]++
				}
			}
		} else { // Handle diagonals ASSUME 45 always
			xDiff := pair.start.x - pair.end.x
			yDiff := pair.start.y - pair.end.y
			xVector := []int{}
			yVector := []int{}
			for i := 0; i <= absInt(xDiff); i++ {
				var xNum, yNum int
				if xDiff > 0 {
					xNum = -i
				} else {
					xNum = i
				}
				xVector = append(xVector, xNum)
				if yDiff > 0 {
					yNum = -i
				} else {
					yNum = i
				}
				yVector = append(yVector, yNum)
			}

			if absInt(xDiff) != absInt(yDiff) {
				// This is not a 45 degree angle
				continue
			}
			for i := 0; i <= absInt(xDiff); i++ {
				pointHere := point{
					x: pair.start.x + xVector[i],
					y: pair.start.y + yVector[i],
				}
				_, ok := hitSet[pointHere]
				if !ok {
					hitSet[pointHere] = 1
				} else {
					hitSet[pointHere]++
				}
			}

		}
	}
	counter = 0
	for _, val := range hitSet {
		if val > 1 {
			counter++
		}
	}
	fmt.Println(counter)
}

func pointFromString(in string) point {
	nums := strings.Split(in, ",")
	x, _ := strconv.ParseInt(nums[0], 10, 64)
	y, _ := strconv.ParseInt(nums[1], 10, 64)
	return point{
		x: int(x),
		y: int(y),
	}
}

func absInt(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}
