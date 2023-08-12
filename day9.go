package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

type SnakeInstruction struct {
	direction Direction
	steps     int
}

type Position struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInstructionsFromStrings(rows []string) *[]SnakeInstruction {
	instructions := make([]SnakeInstruction, 0)

	for _, row := range rows {
		parsedRow := strings.Split(row, " ")

		var direction Direction
		switch parsedRow[0] {
		case "U":
			direction = Up
		case "D":
			direction = Down
		case "L":
			direction = Left
		case "R":
			direction = Right
		default:
			log.Fatalf("Not a valid direction while parsing: %s", parsedRow[0])
		}

		steps, err := strconv.Atoi(parsedRow[1])
		if err != nil {
			log.Fatalf("Not a valid number while parsing: %s", parsedRow[1])
		}

		instruction := SnakeInstruction{direction: direction, steps: steps}
		instructions = append(instructions, instruction)
	}

	return &instructions
}

func isPosAdjecent(pos1 Position, pos2 Position) bool {
    x_diff := abs(pos1.x - pos2.x)
    y_diff := abs(pos1.y - pos2.y)

    return 0 <= x_diff && x_diff <= 1 && 0 <= y_diff && y_diff <= 1
}

func getNewPos(headPos Position, tailPos Position) Position {
    if headPos.x == tailPos.x && (headPos.y - tailPos.y) >= 1 {
        // north
        return Position{tailPos.x, tailPos.y + 1}
    } else if headPos.x == tailPos.x && (headPos.y - tailPos.y) <= 1 {
        // south
        return Position{tailPos.x, tailPos.y - 1}
    } else if headPos.y == tailPos.y && (headPos.x - tailPos.x) >= 1 {
        // east
        return Position{tailPos.x + 1, tailPos.y}
    } else if headPos.y == tailPos.y && (headPos.x - tailPos.x) <= 1 {
        // west
        return Position{tailPos.x - 1, tailPos.y}
    } else if (headPos.x - tailPos.x) >= 1 && (headPos.y - tailPos.y) >= 1 {
        // north-east
        return Position{tailPos.x + 1, tailPos.y + 1}
    } else if (headPos.x - tailPos.x) >= 1 && (headPos.y - tailPos.y) <= 1 {
        // south-east
        return Position{tailPos.x + 1, tailPos.y - 1}
    } else if (headPos.x - tailPos.x) <= 1 && (headPos.y - tailPos.y) >= 1 {
        // north-west
        return Position{tailPos.x - 1, tailPos.y + 1}
    } else if (headPos.x - tailPos.x) <= 1 && (headPos.y - tailPos.y) <= 1 {
        // south-west
        return Position{tailPos.x - 1, tailPos.y - 1}
    } else {
        log.Fatalf("Not adjecent positions: %v, %v", headPos, tailPos)
    }

    return Position{0, 0}
}

func day9_part1() {
	rows, err := getRowsFromFile("input9.txt")

	if err != nil {
		log.Fatal(err)
	}

    instructions := getInstructionsFromStrings(rows[:len(rows)-1])

    numberOfSections := 2 
    visitedPositions := make([]map[Position]bool, numberOfSections) 
    snake := make([]Position, numberOfSections)

    for i := 0; i < numberOfSections; i++ {
        visitedPositions[i] = make(map[Position]bool) 
        snake[i] = Position{0, 0}
        visitedPositions[i][snake[i]] = true
    }

    for _, instruction := range *instructions {
        head := &snake[0]
        for step := 0; step < instruction.steps; step++ {
            switch instruction.direction {
            case Up:
                head.y += 1
            case Down:
                head.y -= 1
            case Left:
                head.x -= 1
            case Right:
                head.x += 1
            default:
                log.Fatalf("Not a valid direction: %s", instruction.direction)
            }

            visitedPositions[0][*head] = true

            currentHead := snake[0]
            for i := 1; i < numberOfSections; i++ {
                if !isPosAdjecent(currentHead, snake[i]) {
                    snake[i] = getNewPos(currentHead, snake[i])
                    visitedPositions[i][snake[i]] = true
                }
                currentHead = snake[i]
            }
        }
    }
    solution := len(visitedPositions[numberOfSections - 1])
    fmt.Println(getFunctionName(), solution)
}

func day9_part2() {
	rows, err := getRowsFromFile("input9.txt")

	if err != nil {
		log.Fatal(err)
	}

    instructions := getInstructionsFromStrings(rows[:len(rows)-1])

    numberOfSections := 10 
    visitedPositions := make([]map[Position]bool, numberOfSections) 
    snake := make([]Position, numberOfSections)

    for i := 0; i < numberOfSections; i++ {
        visitedPositions[i] = make(map[Position]bool) 
        snake[i] = Position{0, 0}
        visitedPositions[i][snake[i]] = true
    }

    for _, instruction := range *instructions {
        head := &snake[0]
        for step := 0; step < instruction.steps; step++ {
            switch instruction.direction {
            case Up:
                head.y += 1
            case Down:
                head.y -= 1
            case Left:
                head.x -= 1
            case Right:
                head.x += 1
            default:
                log.Fatalf("Not a valid direction: %s", instruction.direction)
            }

            visitedPositions[0][*head] = true

            currentHead := snake[0]
            for i := 1; i < numberOfSections; i++ {
                if !isPosAdjecent(currentHead, snake[i]) {
                    snake[i] = getNewPos(currentHead, snake[i])
                    visitedPositions[i][snake[i]] = true
                }
                currentHead = snake[i]
            }
        }
    }
    solution := len(visitedPositions[numberOfSections - 1])
    fmt.Println(getFunctionName(), solution)
}
