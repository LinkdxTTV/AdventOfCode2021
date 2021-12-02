package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	horiz := 0
	depth := 0

	// Record for part 2
	type command struct {
		direction string
		value     int
	}

	commandList := []command{}

	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(commands[1])

		commandList = append(commandList, command{commands[0], num})
		if err != nil {
			log.Fatal(err)
		}
		switch commands[0] {
		case "forward":
			horiz += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	fmt.Println(horiz, depth, horiz*depth)

	horiz = 0
	depth = 0
	aim := 0

	// Part 2
	for _, command := range commandList {
		switch command.direction {
		case "forward":
			horiz += command.value
			depth += command.value * aim
		case "down":
			aim += command.value
		case "up":
			aim -= command.value
		}
	}

	fmt.Println(horiz, depth, horiz*depth)
}
