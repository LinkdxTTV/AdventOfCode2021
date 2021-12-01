package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input1")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	inc := 0

	for i, num := range numbers {
		if i == len(numbers)-1 {
			break
		}

		if numbers[i+1] > num {
			inc++
		}
	}

	fmt.Println(inc)
	windowed := []int{}

	// Part 2
	for i := range numbers {
		if i < 1 || i > len(numbers)-2 {
			continue
		}
		sum := numbers[i-1] + numbers[i] + numbers[i+1]
		windowed = append(windowed, sum)
	}

	windowsinc := 0

	for i, num := range windowed {
		if i == len(windowed)-1 {
			break
		}

		if windowed[i+1] > num {
			windowsinc++
		}
	}

	fmt.Println(windowsinc)
}
