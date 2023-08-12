package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// noop = 1 cycle, and does nothing
// addx = 2 cycles, adds value to the x register

// cycle 1 - noop    - starts executing noop, and finishes - x register = 1
// cycle 2 - addx 3  - starts executing addx               - x register = 1
// cycle 3 - addx 3  - finishing execution of addx         - x register = 4
// cycle 4 - addx-5  - starts executing addx               - x register = 4
// cycle 5 - addx-5  - finishing execution of addx         - x register = -1

// the signal strength is the cycle number * the x register value
// calculate the signal strength for 20th cycle and after that 40 cycles after that
// (that is, during the 20th, 60th, 100th, 140th, 180th, and 220th cycles)

type Operations int

const (
	NOOP Operations = iota
	ADDX
)

type Instruction struct {
	operation Operations
	value     int
}

func getInstructionFromString(instructionString string) Instruction {
	parsedString := strings.Split(instructionString, " ")
	if parsedString[0] == "noop" {
		return Instruction{NOOP, 0}
	} else {
		value, err := strconv.Atoi(parsedString[1])
		if err != nil {
			log.Fatal(err)
		}
		return Instruction{ADDX, value}
	}
}

func getInstructionsFromStringArray(instructionStrings []string) *[]Instruction {
	instructions := make([]Instruction, len(instructionStrings))
	for idx, instructionString := range instructionStrings {
		instructions[idx] = getInstructionFromString(instructionString)
	}
	return &instructions
}

func getSignalStrength(cycle int, xRegister int) int {
	if cycle == 20 || (cycle > 40 && (cycle-20)%40 == 0) {
		return cycle * xRegister
	} else {
		return 0
	}
}

func getSumOfSignalStrength(instructions *[]Instruction) int {
	sumOfSignalStrength := 0
	xRegister := 1
	cycle := 0
	idx := 0

	for {
		if idx >= len(*instructions) {
			break
		}

		instruction := (*instructions)[idx]
		if instruction.operation == NOOP {
			cycle++
			sumOfSignalStrength += getSignalStrength(cycle, xRegister)
		} else {
			cycle++
			sumOfSignalStrength += getSignalStrength(cycle, xRegister)

			cycle++
			sumOfSignalStrength += getSignalStrength(cycle, xRegister)
			xRegister += instruction.value
		}

		idx++
	}

	return sumOfSignalStrength
}

func day10_part1() {
	rows, err := getRowsFromFile("input10.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions := getInstructionsFromStringArray(rows[:len(rows)-1])
	solution := getSumOfSignalStrength(instructions)
	fmt.Println(getFunctionName(), " solution: ", solution)
}

func printPixels(pixels *[][]byte) {
	for _, pixelRow := range *pixels {
		fmt.Println(string(pixelRow))
	}
}

func drawPixelsFromInstructions(instructions *[]Instruction) *[][]byte {
	pixels := make([][]byte, 6)
	for idx := range pixels {
		pixels[idx] = make([]byte, 40)
		for jdx := range pixels[idx] {
			pixels[idx][jdx] = '.'
		}
	}

	pixelIdx := 0
	pixelJdx := 0

	xRegister := 1
	cycle := 0
	idx := 0

	drawPixels := func() {
		if xRegister-1 <= pixelJdx && pixelJdx <= xRegister+1 {
			pixels[pixelIdx][pixelJdx] = '#'
		}
	}

	cycleTick := func() {
		cycle++
		pixelJdx++
		if pixelJdx%40 == 0 {
			pixelIdx++
			pixelJdx = 0
		}
	}

	for {
		if idx >= len(*instructions) {
			break
		}

		instruction := (*instructions)[idx]
		if instruction.operation == NOOP {
			drawPixels()
			cycleTick()
		} else {
			drawPixels()
			cycleTick()

			drawPixels()
			cycleTick()
			xRegister += instruction.value

		}

		idx++
	}

	return &pixels
}

func day10_part2() {
	rows, err := getRowsFromFile("input10.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions := getInstructionsFromStringArray(rows[:len(rows)-1])
	pixels := drawPixelsFromInstructions(instructions)
	printPixels(pixels)
}
