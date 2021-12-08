package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
	inputs := [][]string{}
	outputs := [][]string{}

	for scanner.Scan() {
		inout := strings.Split(scanner.Text(), "|")
		in := strings.Split(inout[0], " ")
		inputs = append(inputs, in)
		out := strings.Split(inout[1], " ")
		outputs = append(outputs, out)
	}
	total := 0

	for _, nums := range outputs {
		for _, num := range nums {
			if len(num) == 2 || len(num) == 4 || len(num) == 3 || len(num) == 7 {
				total++
			}
		}
	}

	fmt.Println("Part 1", total)

	/*
		number | line segments
		0		6
		1		2		unique
		2		5
		3		5
		4		4		unique
		5		5
		6		6
		7		3		unique
		8		7		unique
		9		6
	*/

	// Part 2 lol
	outputNums := []int{}

	for i, input := range inputs {
		mapping := map[int]number{}
		filteredList := []string{}
		filteredList = append(filteredList, input...)
		// Get 1, 4, 7, and 8
		for _, str := range input {
			if len(str) == 2 {
				num := createNumber(1, str)
				mapping[1] = num
				filteredList = removeStrFromList(filteredList, str)
				continue
			}
			if len(str) == 4 {
				num := createNumber(4, str)
				mapping[4] = num
				filteredList = removeStrFromList(filteredList, str)
				continue
			}
			if len(str) == 3 {
				num := createNumber(7, str)
				mapping[7] = num
				filteredList = removeStrFromList(filteredList, str)
				continue
			}

			if len(str) == 7 {
				num := createNumber(8, str)
				mapping[8] = num
				filteredList = removeStrFromList(filteredList, str)
				continue
			}
		}
		// Figure out 5s
		for _, str := range filteredList {
			if len(str) != 5 {
				continue
			}
			strMap := map[string]bool{}
			for _, letter := range str {
				strMap[string(letter)] = true
			}
			// If it shares 2 letters with 1, and has 5 letters, it must be 3
			if checkSimilarLetters(mapping[1].letters, strMap) == 2 {
				mapping[3] = createNumber(3, str)
				filteredList = removeStrFromList(filteredList, str)
			}
		}

		for _, str := range filteredList {
			if len(str) != 5 {
				continue
			}
			strMap := map[string]bool{}
			for _, letter := range str {
				strMap[string(letter)] = true
			}
			// Shares 3 letters with 4, must be 5
			if checkSimilarLetters(mapping[4].letters, strMap) == 3 {
				mapping[5] = createNumber(5, str)
				filteredList = removeStrFromList(filteredList, str)
			}
		}

		// Last number must be 2
		for _, str := range filteredList {
			if len(str) != 5 {
				continue
			}
			mapping[2] = createNumber(2, str)
			filteredList = removeStrFromList(filteredList, str)
		}

		// Need to find 0, 6, 9
		for _, str := range filteredList {
			if len(str) != 6 {
				continue
			}
			strMap := map[string]bool{}
			for _, letter := range str {
				strMap[string(letter)] = true
			}
			// 9 shares 4 with 4
			if checkSimilarLetters(mapping[4].letters, strMap) == 4 {
				mapping[9] = createNumber(9, str)
				filteredList = removeStrFromList(filteredList, str)
			}
		}

		// 6 vs 0
		for _, str := range filteredList {
			if len(str) != 6 {
				continue
			}
			strMap := map[string]bool{}
			for _, letter := range str {
				strMap[string(letter)] = true
			}
			//
			if checkSimilarLetters(mapping[7].letters, strMap) == 3 {
				mapping[0] = createNumber(0, str)
				filteredList = removeStrFromList(filteredList, str)
			}
		}

		// Last number must be 6
		for _, str := range filteredList {
			if len(str) != 6 {
				continue
			}
			mapping[6] = createNumber(6, str)
			filteredList = removeStrFromList(filteredList, str)
		}

		if len(mapping) != 10 {
			log.Fatal("wtf", mapping)
		}

		outputStr := ""
		for _, out := range outputs[i] {
			if len(out) == 0 {
				continue
			}
			num := strMap2Number(mapping, out)
			if num == 10 {
				log.Fatal("oh no 10")
			} // shitty error handle
			outputStr += fmt.Sprintf("%d", num)
		}
		num, _ := strconv.ParseInt(outputStr, 10, 64)
		outputNums = append(outputNums, int(num))
	}

	fmt.Println("Outputs \n", outputNums)
	total = 0
	for _, num := range outputNums {
		total += num
	}

	fmt.Println("Sum of outputs:", total)
}

type number struct {
	number  int
	letters map[string]bool
}

func strMap2Number(mapping map[int]number, letters string) int {
	strMap := map[string]bool{}
	for _, letter := range letters {
		strMap[string(letter)] = true
	}

	for key, value := range mapping {
		if reflect.DeepEqual(value.letters, strMap) {
			return key
		}
	}
	fmt.Println(mapping)
	fmt.Println(letters)
	return 10
}

func checkSimilarLetters(a, b map[string]bool) int {
	total := 0
	for key, _ := range a {
		_, ok := b[key]
		if ok {
			total++
		}
	}
	return total
}

func createNumber(num int, letters string) number {
	number := number{
		number:  num,
		letters: map[string]bool{},
	}

	for _, letter := range letters {
		number.letters[string(letter)] = true
	}

	return number
}

func removeStrFromList(list []string, str string) []string {
	output := []string{}
	for _, entry := range list {
		if entry != str {
			output = append(output, entry)
		}
	}

	return output
}
