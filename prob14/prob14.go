package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	steps := false
	template := ""
	stepMapping := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			steps = true
			continue
		}
		if !steps {
			template = line
			continue
		}
		if steps {
			stepSplit := strings.Split(line, " -> ")
			stepMapping[stepSplit[0]] = stepSplit[1]
		}
	}
	templateCopy := template

	totalSteps := 10

	for steps := 0; steps < totalSteps; steps++ {
		tempTemplate := ""
		for i := range template {
			tempTemplate += string(template[i])
			if i+1 == len(template) {
				break
			}
			insert, ok := stepMapping[template[i:i+2]]
			if !ok {
				continue
			}
			tempTemplate += insert
		}
		template = tempTemplate
	}

	count := map[string]int{}
	for _, char := range template {
		_, ok := count[string(char)]
		if !ok {
			count[string(char)] = 1
		} else {
			count[string(char)]++
		}
	}

	var max, min = 0, 1000000
	var maxLetter, minLetter string
	for key, val := range count {
		if val > max {
			max = val
			maxLetter = key
		}
		if val < min {
			min = val
			minLetter = key
		}
	}

	fmt.Println("Part 1:")
	fmt.Println("Max Letter:", maxLetter, max, "Min Letter:", minLetter, min, "Diff:", max-min)

	// Part 2
	// Sequencing is pretty damn deterministic. i.e. CP -> C always makes CC and CP. We can probably iterate from there
	// Later on we will need to do some quickMATH in order to find the literal amount of each letter

	// On second thought maybe we can keep track as we go
	sequences := map[string]int{}
	for key := range stepMapping {
		sequences[key] = 0
	}

	letterCount := map[string]int{}

	// Initial sequences
	for i := range templateCopy {
		incrementMap(letterCount, templateCopy[i:i+1], 1)
		if i+1 == len(templateCopy) {
			continue
		}
		incrementMap(sequences, templateCopy[i:i+2], 1)
	}

	// Steppin time
	for step := 0; step < 40; step++ {
		tempSequences := map[string]int{}
		for key, value := range sequences {
			spawn, ok := stepMapping[key]
			if !ok {
				continue
			}
			incrementMap(letterCount, spawn, value)
			leftHalf := key[:1] + spawn
			rightHalf := spawn + key[1:]
			incrementMap(tempSequences, leftHalf, value)
			incrementMap(tempSequences, rightHalf, value)
		}
		sequences = tempSequences
	}

	max, min = 0, math.MaxInt64 // wtf all the numbers are actually so high
	maxLetter, minLetter = "", ""
	for key, val := range letterCount {
		if val > max {
			max = val
			maxLetter = key
		}
		if val < min {
			min = val
			minLetter = key
		}
	}

	fmt.Println("Part 2:")
	fmt.Println("Max Letter:", maxLetter, max, "Min Letter:", minLetter, min, "Diff:", max-min)
}

func incrementMap(letterCount map[string]int, letter string, value int) {
	_, ok := letterCount[letter]
	if !ok {
		letterCount[letter] = value
	} else {
		letterCount[letter] += value
	}
}
