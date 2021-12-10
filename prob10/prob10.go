package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	incompleteLineOpenings := [][]string{}
	corruptedLines := []string{}
	badIndex := []int{}

	for scanner.Scan() {
		var ok, corrupted bool
		nextLine := false
		line := scanner.Text()
		openers := []string{}
		for i, char := range line {
			switch string(char) {
			case "{":
				openers = append(openers, "{")
			case "[":
				openers = append(openers, "[")
			case "(":
				openers = append(openers, "(")
			case "<":
				openers = append(openers, "<")
			case "}":
				if ok, openers = checkOpenersAndRemove(openers, "{"); !ok {
					corruptedLines = append(corruptedLines, line)
					badIndex = append(badIndex, i)
					nextLine = true
					corrupted = true
				}
			case "]":
				if ok, openers = checkOpenersAndRemove(openers, "["); !ok {
					corruptedLines = append(corruptedLines, line)
					badIndex = append(badIndex, i)
					nextLine = true
					corrupted = true
				}
			case ")":
				if ok, openers = checkOpenersAndRemove(openers, "("); !ok {
					corruptedLines = append(corruptedLines, line)
					badIndex = append(badIndex, i)
					nextLine = true
					corrupted = true
				}
			case ">":
				if ok, openers = checkOpenersAndRemove(openers, "<"); !ok {
					corruptedLines = append(corruptedLines, line)
					badIndex = append(badIndex, i)
					nextLine = true
					corrupted = true
				}
			}
			if nextLine {
				break
			}
		}
		if !corrupted {
			incompleteLineOpenings = append(incompleteLineOpenings, openers)
		}
	}
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	total := 0
	for i, line := range corruptedLines {
		total += points[string(line[badIndex[i]])]
	}

	fmt.Println("Part1:", total)

	// Part 2
	reverseSymbol := map[string]string{
		"{": "}",
		"[": "]",
		"<": ">",
		"(": ")",
	}

	incompleteLineClosings := [][]string{}
	for _, incompleteLineOpening := range incompleteLineOpenings {
		incompleteLineClosing := []string{}
		for _, char := range incompleteLineOpening {
			incompleteLineClosing = append(incompleteLineClosing, reverseSymbol[char])
		}

		incompleteLineClosings = append(incompleteLineClosings, reverse(incompleteLineClosing))
	}

	points2 := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scores := []int{}
	for _, line := range incompleteLineClosings {
		total := 0
		for _, str := range line {
			total = total * 5
			total += points2[str]
		}
		scores = append(scores, total)
	}
	sort.Ints(scores)
	middleScore := scores[(len(scores) / 2)]
	fmt.Println("Part2:", middleScore)
}

func checkOpenersAndRemove(openers []string, char string) (bool, []string) {
	if openers[len(openers)-1] != char {
		return false, openers
	}
	openers = openers[:len(openers)-1]
	return true, openers
}

func reverse(input []string) []string {
	output := make([]string, len(input))
	for i, char := range input {
		output[len(input)-i-1] = char
	}
	return output
}
