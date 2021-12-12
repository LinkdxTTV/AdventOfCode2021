package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	pointConnections := map[string]map[string]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "-")
		start, end := points[0], points[1]
		_, ok := pointConnections[start]
		if !ok {
			pointConnections[start] = map[string]bool{}
		}
		pointConnections[start][end] = true

		_, ok = pointConnections[end]
		if !ok {
			pointConnections[end] = map[string]bool{}
		}
		pointConnections[end][start] = true
	}

	pathSet := map[string]bool{}
	pathSoFar := []string{}
	recursivelyMove(pathSet, pointConnections, "start", pathSoFar)

	count := 0
	for key := range pathSet {
		if key[len(key)-3:] == "end" {
			count++
		}
	}
	fmt.Println("Part1:", count)

	// Part 2
	pathSet = map[string]bool{}
	pathSoFar = []string{}
	recursivelyMovePart2(pathSet, pointConnections, "start", pathSoFar)
	count = 0
	for key := range pathSet {
		if key[len(key)-3:] == "end" {
			count++
		}
	}
	fmt.Println("Part2:", count)
}

func isBigCave(input string) bool {
	for _, char := range input {
		if unicode.IsLower(char) {
			return false
		}
	}
	return true
}

func recursivelyMove(pathSet map[string]bool, pointConnections map[string]map[string]bool, point string, pathSoFar []string) {
	pathSoFar = append(pathSoFar, point)
	_, ok := pathSet[reducePathToString(pathSoFar)]
	if ok {
		return
	}
	pathSet[reducePathToString(pathSoFar)] = true

	if point == "end" {
		return
	}

	for nextPoint := range pointConnections[point] {
		// You can only retraverse big caves
		if !isBigCave(nextPoint) && haveTraversedBefore(nextPoint, pathSoFar) {
			continue
		}
		recursivelyMove(pathSet, pointConnections, nextPoint, pathSoFar)
	}
}

func recursivelyMovePart2(pathSet map[string]bool, pointConnections map[string]map[string]bool, point string, pathSoFar []string) {
	pathSoFar = append(pathSoFar, point)
	_, ok := pathSet[reducePathToString(pathSoFar)]
	if ok {
		return
	}
	pathSet[reducePathToString(pathSoFar)] = true

	if point == "end" {
		return
	}

	for nextPoint := range pointConnections[point] {
		// Small caves can only be retraversed once
		if !isBigCave(nextPoint) {
			// Been here before AND we already visited a cave twice
			if haveTraversedBefore(nextPoint, pathSoFar) && anySmallCaveVisitedTwice(pathSoFar) {
				continue
			}
		}
		if nextPoint == "start" {
			continue
		}
		recursivelyMovePart2(pathSet, pointConnections, nextPoint, pathSoFar)
	}
}

func reducePathToString(input []string) string {
	path := ""
	for i, point := range input {
		path += point
		if i != len(input)-1 {
			path += "-"
		}
	}
	return path // Gets rid of last dash
}

func haveTraversedBefore(point string, path []string) bool {
	for _, test := range path {
		if point == test {
			return true
		}
	}
	return false
}

func anySmallCaveVisitedTwice(path []string) bool {
	visits := map[string]bool{}
	for _, place := range path {
		_, ok := visits[place]
		if ok {
			return true
		}
		if !isBigCave(place) {
			visits[place] = true
		}
	}
	return false
}
