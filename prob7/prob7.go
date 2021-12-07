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

	scanner.Scan()
	line := scanner.Text()

	numsStr := strings.Split(line, ",")

	nums := []int{}

	for _, num := range numsStr {
		numint, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, int(numint))
	}

	var minFuel int = 1000000000
	var point int
	min, max := minmax(nums)
	for i := min; i <= max; i++ {
		fuel := calculateFuel(nums, i)
		if fuel < minFuel {
			point = i
			minFuel = fuel
		}
	}

	fmt.Println(point, minFuel)

	minFuel = 1000000000
	for i := min; i <= max; i++ {
		fuel := calculateFuel2(nums, i)
		if fuel < minFuel {
			point = i
			minFuel = fuel
		}
	}

	fmt.Println(point, minFuel)
}

func calculateFuel(nums []int, point int) int {
	totalFuel := 0
	for _, num := range nums {
		fuel := abs(num - point)
		totalFuel += fuel
	}

	return totalFuel
}

func calculateFuel2(nums []int, point int) int {
	totalFuel := 0
	for _, num := range nums {
		n := abs(num - point)
		fuel := n * (n + 1) / 2
		totalFuel += fuel
	}

	return totalFuel
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minmax(nums []int) (int, int) {
	var min, max int
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}
