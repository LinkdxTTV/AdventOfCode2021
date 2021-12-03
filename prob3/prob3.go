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

	numbers := []string{}

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	// Part 1
	length := len(numbers[0])

	mostCommon := ""

	for i := 0; i < length; i++ {
		ones := 0
		zeros := 0
		for _, num := range numbers {
			if string(num[i]) == "0" {
				zeros++
			} else if string(num[i]) == "1" {
				ones++
			}
		}

		if ones > zeros {
			mostCommon += "1"
		} else {
			mostCommon += "0"
		}
	}

	gamma, _ := strconv.ParseInt(mostCommon, 2, 64)
	leastCommon := bitFlip(mostCommon)
	eps, _ := strconv.ParseInt(leastCommon, 2, 64)
	fmt.Println("gamma:", gamma, "eps:", eps, "power:", gamma*eps)

	// Part 2

	numbersO2 := numbers

	for i := 0; i < length; i++ {
		if len(numbersO2) == 1 {
			break
		}
		ones := 0
		zeros := 0
		var keep string
		for _, num := range numbersO2 {
			if string(num[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if ones >= zeros {
			keep = "1"
		} else if zeros > ones {
			keep = "0"
		}

		// Filter the list
		filtered := []string{}
		for _, num := range numbersO2 {
			if string(num[i]) == keep {
				filtered = append(filtered, num)
			}
		}

		// Restart loop with new filtered list
		numbersO2 = filtered
	}

	// Do the same for CO2 / least common
	numbersCO2 := numbers

	for i := 0; i < length; i++ {
		if len(numbersCO2) == 1 {
			break
		}
		ones := 0
		zeros := 0
		var keep string
		for _, num := range numbersCO2 {
			if string(num[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if ones >= zeros {
			keep = "0"
		} else if zeros > ones {
			keep = "1"
		}

		filtered := []string{}
		for _, num := range numbersCO2 {
			if string(num[i]) == keep {
				filtered = append(filtered, num)
			}
		}

		numbersCO2 = filtered
	}

	oxygen, _ := strconv.ParseInt(numbersO2[0], 2, 64)
	co2, _ := strconv.ParseInt(numbersCO2[0], 2, 64)

	fmt.Println("o2:", oxygen, "co2:", co2, "multiple:", oxygen*co2)
}

func bitFlip(in string) string {
	out := ""
	for _, char := range in {
		if string(char) == "0" {
			out += "1"
		} else if string(char) == "1" {
			out += "0"
		}
	}
	return out
}
