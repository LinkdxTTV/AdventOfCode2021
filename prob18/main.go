package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var GlobalList []*int = []*int{}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	nums := []SnailfishNumber{}
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		num := parseSnailfishNumber(line, nil)
		lines = append(lines, line)
		nums = append(nums, num)
	}

	runningSum := nums[0]
	AppendGlobalList(&runningSum, &GlobalList)
	reduceSnailfishNumber(&runningSum)
	for _, num := range nums[1:] {
		runningSum = addSnailfishNumbers(runningSum, num)
		GlobalList = []*int{}
		AppendGlobalList(&runningSum, &GlobalList)
		reduceSnailfishNumber(&runningSum)
	}

	GlobalList = []*int{}
	AppendGlobalList(&runningSum, &GlobalList)
	for _, ptr := range GlobalList {
		fmt.Print(*ptr, " ")
	}
	fmt.Println()
	fmt.Println("Part 1:", calcualateMagnitude(&runningSum))

	// Part 2

	// Oh no pointers fucked me, re parse from strings
	max := 0
	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j {
				continue
			}
			num1 := parseSnailfishNumber(line1, nil)
			num2 := parseSnailfishNumber(line2, nil)
			combination := addSnailfishNumbers(num2, num1)
			GlobalList = []*int{}
			AppendGlobalList(&combination, &GlobalList)
			reduceSnailfishNumber(&combination)
			comboScore := calcualateMagnitude(&combination)
			if comboScore > max {
				max = comboScore
			}
		}
	}

	fmt.Println("Part 2:", max)
}

type SnailfishNumber struct {
	above      *SnailfishNumber
	valueLeft  *int
	valueRight *int
	pairLeft   *SnailfishNumber
	pairRight  *SnailfishNumber
}

func parseSnailfishNumber(input string, cameFrom *SnailfishNumber) SnailfishNumber {
	out := SnailfishNumber{}
	left, right := splitSnailfishString(input)
	if strings.Contains(left, "[") {
		snailLeft := parseSnailfishNumber(left, &out)
		out.pairLeft = &snailLeft
	} else {
		numLeft, err := strconv.ParseInt(left, 10, 64)
		if err != nil {
			log.Fatal("big oops", err)
		}
		numLeftInt := int(numLeft)
		out.valueLeft = &numLeftInt
	}
	if strings.Contains(right, "[") {
		snailRight := parseSnailfishNumber(right, &out)
		out.pairRight = &snailRight
	} else {
		numRight, err := strconv.ParseInt(right, 10, 64)
		if err != nil {
			log.Fatal("big oops", err)
		}
		numRightInt := int(numRight)
		out.valueRight = &numRightInt
	}
	return out
}

func splitSnailfishString(input string) (string, string) {
	commaPos := 0
	markerCount := 0
	for i, char := range input {
		if string(char) == "[" {
			markerCount++
			continue
		}
		if string(char) == "]" {
			markerCount--
			continue
		}
		if string(char) == "," {
			if markerCount == 1 {
				commaPos = i
				break
			}
		}
	}
	return input[1:commaPos], input[commaPos+1 : len(input)-1]
}

func addSnailfishNumbers(first, second SnailfishNumber) SnailfishNumber {
	out := SnailfishNumber{nil, nil, nil, &first, &second}
	return out
}

func reduceSnailfishNumber(num *SnailfishNumber) {
	for {
		GlobalList = []*int{}
		AppendGlobalList(num, &GlobalList)

		if explodeIfPossible(num, 1) {
			continue
		}

		// Look for splits
		if splitIfPossible(num) {
			continue
		}
		// Stop
		break
	}
}

func splitIfPossible(num *SnailfishNumber) bool {
	if num.valueLeft != nil && *num.valueLeft >= 10 {
		left := *num.valueLeft / 2
		right := *num.valueLeft / 2
		if *num.valueLeft%2 != 0 {
			right++
		}

		num.pairLeft = &SnailfishNumber{num, &left, &right, nil, nil}
		num.valueLeft = nil
		return true
	}
	if num.pairLeft != nil {
		if splitIfPossible(num.pairLeft) {
			return true
		}
	}
	if num.pairRight != nil {
		if splitIfPossible(num.pairRight) {
			return true
		}
	}
	if num.valueRight != nil && *num.valueRight >= 10 {
		left := *num.valueRight / 2
		right := *num.valueRight / 2
		if *num.valueRight%2 != 0 {
			right++
		}

		num.pairRight = &SnailfishNumber{num, &left, &right, nil, nil}
		num.valueRight = nil
		return true
	}

	return false
}

func explodeIfPossible(num *SnailfishNumber, depth int) bool {
	if depth > 4 {
		log.Fatal("yikes")
	}
	if depth == 4 && num.pairLeft != nil {
		pair := num.pairLeft
		if pair.valueLeft != nil {
			num.addToFirstLeft(pair.valueLeft)
		}
		if pair.valueRight != nil {
			num.addToFirstRight(pair.valueRight)
		}
		num.pairLeft = nil
		zero := 0
		num.valueLeft = &zero
		return true
	}
	if depth == 4 && num.pairRight != nil {
		pair := num.pairRight
		if pair.valueLeft != nil {
			num.addToFirstLeft(pair.valueLeft)
		}
		if pair.valueRight != nil {
			num.addToFirstRight(pair.valueRight)
		}
		num.pairRight = nil
		zero := 0
		num.valueRight = &zero
		return true
	}
	if num.pairLeft != nil {
		if explodeIfPossible(num.pairLeft, depth+1) {
			return true
		}
	}
	if num.pairRight != nil {
		if explodeIfPossible(num.pairRight, depth+1) {
			return true
		}
	}
	return false
}

func AppendGlobalList(node *SnailfishNumber, list *[]*int) {
	if node.pairLeft != nil {
		AppendGlobalList(node.pairLeft, list)
	}
	if node.valueLeft != nil {
		*list = append(*list, node.valueLeft)
	}
	if node.pairRight != nil {
		AppendGlobalList(node.pairRight, list)
	}
	if node.valueRight != nil {
		*list = append(*list, node.valueRight)
	}
}

func (n *SnailfishNumber) addToFirstLeft(valueToAdd *int) {
	for i, pointer := range GlobalList {
		if pointer == valueToAdd {
			if i == 0 {
				return
			}
			*GlobalList[i-1] += *valueToAdd
			return
		}
	}
}

func (n *SnailfishNumber) addToFirstRight(valueToAdd *int) {
	for i, pointer := range GlobalList {
		if pointer == valueToAdd {
			if i == len(GlobalList)-1 {
				return
			}
			*GlobalList[i+1] += *valueToAdd
		}
	}
}

func calcualateMagnitude(node *SnailfishNumber) int {
	sum := 0
	if node.valueLeft != nil {
		sum += 3 * *node.valueLeft
	}
	if node.valueRight != nil {
		sum += 2 * *node.valueRight
	}
	if node.pairLeft != nil {
		sum += 3 * calcualateMagnitude(node.pairLeft)
	}
	if node.pairRight != nil {
		sum += 2 * calcualateMagnitude(node.pairRight)
	}
	return sum
}
