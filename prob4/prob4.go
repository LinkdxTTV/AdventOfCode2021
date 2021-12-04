package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingo struct {
	nums   [][]int
	marked [][]bool
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := []int{}

	scanner.Scan()
	firstLine := scanner.Text()
	entries := strings.Split(firstLine, ",")
	for _, num := range entries {
		intNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal("parse int error entries", err)
		}
		numbers = append(numbers, int(intNum))
	}

	fmt.Println("Called Number List:", numbers)
	scanner.Scan()

	lineCounter := 0
	board := createNewBingoBoard(5)
	bingoBoards := []bingo{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "\n" || line == "" {
			bingoBoards = append(bingoBoards, board)
			lineCounter = 0
			continue
		}
		if lineCounter == 0 {
			board = createNewBingoBoard(5)
		}
		nums := strings.Split(line, " ")
		for _, num := range nums {
			if num == "" {
				continue
			}
			numInt, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				log.Fatal("parse int error", err, num)
			}
			board.nums[lineCounter] = append(board.nums[lineCounter], int(numInt))
			board.marked[lineCounter] = append(board.marked[lineCounter], false)
		}

		lineCounter++
	}
	// Grab the last board since the input doesnt end in a new line
	bingoBoards = append(bingoBoards, board)

	// ENORMOUS LOOP TIME
	winner := false
	winningNumber := 0
	winningBoard := bingo{}
	for _, calledNumber := range numbers {
		if winner {
			break
		}
		for i, board := range bingoBoards {
			for j, row := range board.nums {
				for k, num := range row {
					if calledNumber == num {
						bingoBoards[i].marked[j][k] = true
					}
				}
			}
			if checkVictory(board) {
				winner = true
				winningBoard = board
				winningNumber = calledNumber
			}
		}

	}

	// Calculate score
	score := calculateScore(winningBoard)

	fmt.Println("Part 1: \n", "sum:", score, "winning number:", winningNumber, "score:", score*winningNumber)

	// Part 2

	// reset bingoBoards
	for _, board := range bingoBoards {
		for _, row := range board.marked {
			for i := range row {
				row[i] = false
			}
		}
	}
	winningNumber = 0
	winningBoard = bingo{}
	boardsWon := 0

	boardTracker := []bool{}
	for i := 0; i < len(bingoBoards); i++ {
		boardTracker = append(boardTracker, false)
	}

	for _, calledNumber := range numbers {
		if boardsWon == len(bingoBoards) {
			break
		}
		for i, board := range bingoBoards {
			for j, row := range board.nums {
				for k, num := range row {
					if calledNumber == num {
						bingoBoards[i].marked[j][k] = true
					}
				}
			}
			if !boardTracker[i] {
				if checkVictory(board) {
					boardTracker[i] = true
					boardsWon++
					winningBoard = board
					winningNumber = calledNumber
				}
			}
		}
	}

	// Calc score
	score = calculateScore(winningBoard)
	fmt.Println("Part 2: \n", "sum:", score, "winning number:", winningNumber, "score:", score*winningNumber)
}

func createNewBingoBoard(size int) bingo {
	nums := [][]int{}
	marked := [][]bool{}
	for i := 0; i < size; i++ {
		nums = append(nums, []int{})
		marked = append(marked, []bool{})
	}
	return bingo{
		nums:   nums,
		marked: marked,
	}
}

func checkVictory(board bingo) bool {
	size := len(board.marked)
	// Check horizontal victory
	for _, row := range board.marked {
		marks := 0
		for _, mark := range row {
			if mark {
				marks++
			}
		}
		if marks == size {
			return true
		}
	}

	// Check vertical victory
	for i := 0; i < size; i++ {
		marks := 0
		for _, row := range board.marked {
			if row[i] {
				marks++
			}
		}
		if marks == size {
			return true
		}
	}

	return false
}

func calculateScore(winningBoard bingo) int {
	score := 0
	for j, row := range winningBoard.marked {
		for k, mark := range row {
			if !mark {
				score += winningBoard.nums[j][k]
			}
		}
	}
	return score
}
