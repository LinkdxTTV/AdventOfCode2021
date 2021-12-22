package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 444356092776315
// 341960390180808

type dice struct {
	roll int
}

type player struct {
	space int
	score int
}

type scorecard struct {
	player1 int
	player2 int
}

type PlayerState struct {
	player1space int
	player1score int
	player2space int
	player2score int
	whoseTurn    int
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal("we lost")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	split := strings.Split(scanner.Text(), ": ")
	player1Start, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	split = strings.Split(scanner.Text(), ": ")
	player2Start, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start:", int(player1Start), int(player2Start))

	player1, player2 := player{int(player1Start), 0}, player{int(player2Start), 0}
	dice := dice{1}
	for {

		player1Move := dice.Roll() + dice.Roll() + dice.Roll()
		// Hacky modulo but uses 10 and not 0
		player1.space += player1Move
		for player1.space > 10 {
			player1.space -= 10
		}
		player1.score += player1.space
		if player1.score >= 1000 {
			break
		}
		player2Move := dice.Roll() + dice.Roll() + dice.Roll()
		// Hacky modulo but uses 10 and not 0
		player2.space += player2Move
		for player2.space > 10 {
			player2.space -= 10
		}
		player2.score += player2.space
		if player2.score >= 1000 {
			break
		}
	}

	fmt.Println("Part 1:", less(player1.score, player2.score)*(dice.roll-1))

	// Part 2 giga recursion
	playerState := playerStateArrays{player1: scoreWithSpaces{}, player2: scoreWithSpaces{}, playerTurn: 1}
	scorecard := scorecard{0, 0}
	// Initialize
	playerState.player1[0][int(player1Start)] = 1
	playerState.player2[0][int(player2Start)] = 1

	for playerState.shouldContinue() {
		playerState.playRound(&scorecard)
		// playerState.print()
	}
	fmt.Println("Part 2:", more(scorecard.player1, scorecard.player2))
}

func (s playerStateArrays) print() {
	fmt.Println("player 1")
	for score, spaces := range s.player1 {
		fmt.Println(score, ":", spaces)
	}
	fmt.Println("player 2")
	for score, spaces := range s.player2 {
		fmt.Println(score, ":", spaces)
	}
}

func less(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func more(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func (d *dice) Roll() int {
	d.roll++
	return d.roll - 1
}

var diceDistribution map[int]int = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type scoreWithSpaces [21][11]int

type playerStateArrays struct {
	player1    scoreWithSpaces
	player2    scoreWithSpaces
	playerTurn int
}

func (state *playerStateArrays) playRound(scorecard *scorecard) {
	var currentScoreWithSpaces scoreWithSpaces
	if state.playerTurn == 1 {
		currentScoreWithSpaces = state.player1
	} else {
		currentScoreWithSpaces = state.player2
	}

	newScoreWithSpaces := scoreWithSpaces{}
	for score, spaces := range currentScoreWithSpaces {
		endUp := [11]int{}
		for space, num := range spaces {
			for diceRoll, times := range diceDistribution {
				newSpace := space + diceRoll
				for newSpace > 10 {
					newSpace -= 10
				}
				endUp[newSpace] += (times * num)
			}
		}
		for space, players := range endUp {
			if space+score >= 21 {
				if state.playerTurn == 1 {
					scorecard.player1 += (players * countTotalUniverses(state.player2))
				} else {
					scorecard.player2 += (players * countTotalUniverses(state.player1))
				}
				continue
			}
			newScoreWithSpaces[score+space][space] += players
		}
	}

	if state.playerTurn == 1 {
		state.player1 = newScoreWithSpaces
		state.playerTurn = 2
	} else {
		state.player2 = newScoreWithSpaces
		state.playerTurn = 1
	}
}

func (s *playerStateArrays) shouldContinue() bool {
	for _, states := range s.player1 {
		for _, num := range states {
			if num != 0 {
				return true
			}
		}
	}
	for _, states := range s.player2 {
		for _, num := range states {
			if num != 0 {
				return true
			}
		}
	}
	return false
}

func countTotalUniverses(input scoreWithSpaces) int {
	total := 0
	for _, row := range input {
		for _, num := range row {
			total += num
		}
	}
	return total
}
