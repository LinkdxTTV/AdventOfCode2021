package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Part 2
type total struct {
	total int
	lock  sync.Mutex
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()

	numsStr := strings.Split(line, ",")
	fmt.Println(numsStr)

	nums := []int{}

	for _, num := range numsStr {
		numint, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, int(numint))
	}

	part2 := []uint8{}
	for _, num := range nums {
		part2 = append(part2, uint8(num))
	}

	days := 80
	for i := 0; i < days; i++ {
		for i, num := range nums {
			if num == 0 {
				nums[i] = 6
				nums = append(nums, 8)
			} else {
				nums[i]--
			}
		}
	}

	fmt.Println(len(nums))

	fish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, num := range part2 {
		fish[num]++
	}
	fmt.Println(fish)
	days = 256
	for day := 0; day < days; day++ {
		fish[7] += fish[0]
		fish[9] += fish[0]
		temp := fish[1:]
		temp = append(temp, 0)
		fish = temp
	}

	fmt.Println(fish)
	total := 0
	for _, sum := range fish {
		total += sum
	}

	fmt.Println(total)
}
