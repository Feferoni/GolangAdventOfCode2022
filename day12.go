package main

import (
	"fmt"
	"log"
)

func getHeightMapFromRows(rows []string) [][]string {
    heightMap := make([][]string, len(rows))

    for i, row := range rows {
        heightMap[i] = make([]string, len(row))
        for j, char := range row {
            heightMap[i][j] = string(char)
        }
    }

    return heightMap
}

func findHeightPosition(height string, heightMap [][]string) Position {
    for i, row := range heightMap {
        for j, char := range row {
            if char == height {
                return Position{i, j}
            }
        }
    }

    return Position{-1, -1}
}

var directions = []Position{
    {-1, 0},
    {1, 0},
    {0, 1},
    {0, -1},
}

func isValidNextPosition(next Position, rows, cols int) bool {
    return next.x >= 0 && next.x < rows && next.y >= 0 && next.y < cols
}

func distanceBetweenLetters(x, y string) int {
	if x == "S" {
		x = "a"
	}
	if y == "S" {
		y = "a"
	}
	if y == "E" {
		y = "z"
	}
	if x == "E" {
		x = "z"
	}

    asciiX := int(x[0])
	asciiY := int(y[0])

	return asciiY - asciiX
}


type HeightMapPosition struct {
    position Position
    distance int
}

func breadthFirstSearch(heightMap [][]string) int {
    rows := len(heightMap)
    cols := len(heightMap[0])

    visited := map[Position]bool{}

    start := findHeightPosition("S", heightMap)
    finish := findHeightPosition("E", heightMap)

    queue := []HeightMapPosition{{start, 0}}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        currentDistance := current.distance
        currentPosition := current.position
        currentHeight := heightMap[currentPosition.x][currentPosition.y]

        if visited[currentPosition] {
            continue
        }

        visited[currentPosition] = true

        if currentPosition == finish {
            return currentDistance
        }

        for _, dir := range directions {
            nextPosition := Position{currentPosition.x + dir.x, currentPosition.y + dir.y}
            if isValidNextPosition(nextPosition, rows, cols) {
                nextHeight := heightMap[nextPosition.x][nextPosition.y]
                if distanceBetweenLetters(currentHeight, nextHeight) <= 1 {
                    queue = append(queue, HeightMapPosition{nextPosition, currentDistance + 1})
                }
            }
        }
    }

    return -1
}

func day12_part1() {
	rows, err := getRowsFromFile("input12.txt")
	if err != nil {
		log.Fatal(err)
	}

	heightMap := getHeightMapFromRows(rows[:len(rows)-1])
	distance := breadthFirstSearch(heightMap)
	fmt.Println(getFunctionName(), " solution: ", distance)
}

func day12_part2() {
	// rows, err := getRowsFromFile("input11.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(getFunctionName(), " solution: ", solution)
}
