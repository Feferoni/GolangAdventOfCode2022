package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readFromFile(filename string) (string, error) {
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(read), nil
}

func getRowsFromFile(filename string) ([]string, error) {
	file, err := readFromFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(file, "\n"), nil
}

func day1_part1() {
	rows, err := getRowsFromFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	highest_calories := 0
	current_elfs_calories := 0
	for _, row := range rows {
		if row == "" {
			if current_elfs_calories > highest_calories {
				highest_calories = current_elfs_calories
			}
			current_elfs_calories = 0
		} else {
			current_elfs_row_calories, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}
			current_elfs_calories += current_elfs_row_calories
		}
	}
	fmt.Printf("Day1, part1 solution: %d\n", highest_calories)
}

func day1_part2() {
	rows, err := getRowsFromFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	elfs_calories := make([]int, 0)
	current_elfs_calories := 0
	for _, row := range rows {
		if row == "" {
			elfs_calories = append(elfs_calories, current_elfs_calories)
			current_elfs_calories = 0
		} else {
			current_elfs_row_calories, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}
			current_elfs_calories += current_elfs_row_calories
		}
	}
	sort.Slice(elfs_calories, func(i, j int) bool {
		return elfs_calories[i] > elfs_calories[j]
	})

	top_three := 0
	for _, value := range elfs_calories[0:3] {
		top_three += value
	}

	fmt.Printf("Day1, part2 solution: %d\n", top_three)
}

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

func (h Hand) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[h]
}

func convertByteToHand(handByte byte) Hand {
	if handByte == 'A' || handByte == 'X' {
		return Rock
	} else if handByte == 'B' || handByte == 'Y' {
		return Paper
	} else {
		return Scissors
	}
}

func getWinnerPoints(handByte Hand, you Hand) int {
	if handByte == you {
		return 3
	} else if handByte == Rock && you == Paper {
		return 6
	} else if handByte == Paper && you == Scissors {
		return 6
	} else if handByte == Scissors && you == Rock {
		return 6
	} else {
		return 0
	}
}

func getHandPoints(hand Hand) int {
	if hand == Rock {
		return 1
	} else if hand == Paper {
		return 2
	} else {
		return 3
	}
}

// A for Rock, B for Paper, and C for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func day2_part1() {
	rows, err := getRowsFromFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, row := range rows {
		if row == "" {
			continue
		}
		if len(row) != 3 {
			log.Fatal("Invalid input: ", row)
		}
		var opponent Hand = convertByteToHand(row[0])
		var you Hand = convertByteToHand(row[2])

		currentWinnerPoints := getWinnerPoints(opponent, you)
		currentHandPoints := getHandPoints(you)
		currentRoundPoints := currentWinnerPoints + currentHandPoints
		total += currentRoundPoints

		// fmt.Println("Opponent: ", opponent, " You: ", you, " Winner: ", currentWinnerPoints, " Hand: ", currentHandPoints, " Round: ", currentRoundPoints)
	}

	fmt.Printf("Day2, part1 solution: %d\n", total)
}

type ExpectedResult int

const (
	Lost ExpectedResult = iota
	Draw
	Won
)

func (e ExpectedResult) String() string {
	return [...]string{"Lost", "Draw", "Won"}[e]
}

func convertByteToExpectedResult(expectedResultByte byte) ExpectedResult {
	if expectedResultByte == 'X' {
		return Lost
	} else if expectedResultByte == 'Y' {
		return Draw
	} else {
		return Won
	}
}

func getYourHand(expectedResult ExpectedResult, opponent Hand) Hand {
	if expectedResult == Lost {
		if opponent == Rock {
			return Scissors
		} else if opponent == Paper {
			return Rock
		} else {
			return Paper
		}
	} else if expectedResult == Draw {
		return opponent
	} else {
		if opponent == Rock {
			return Paper
		} else if opponent == Paper {
			return Scissors
		} else {
			return Rock
		}
	}
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func day2_part2() {
	rows, err := getRowsFromFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, row := range rows {
		if row == "" {
			continue
		}
		if len(row) != 3 {
			log.Fatal("Invalid input: ", row)
		}
		var opponent Hand = convertByteToHand(row[0])
		var expectedResult ExpectedResult = convertByteToExpectedResult(row[2])
		var you Hand = getYourHand(expectedResult, opponent)

		currentWinnerPoints := getWinnerPoints(opponent, you)
		currentHandPoints := getHandPoints(you)
		currentRoundPoints := currentWinnerPoints + currentHandPoints
		total += currentRoundPoints
	}

	fmt.Printf("Day2, part1 solution: %d\n", total)
}

func main() {
	// day1_part1()
	// day1_part2()
	// day2_part1()
	// day2_part2()
}
