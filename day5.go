package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func findEmptyRowIndex(rows []string) int {
	for i, row := range rows {
		if row == "" {
			return i
		}
	}
	return -1
}

func getMaxStackNumber(row string) int {
	maxNumber := len(row)/4 + 1
	return maxNumber
}

func getStacksArrayFromRows(rows []string) [][]byte {
	highestStackNumber := getMaxStackNumber(rows[len(rows)-1])
	stackArray := make([][]byte, highestStackNumber)

	for i := len(rows) - 2; i >= 0; i-- {
		row := rows[i][1 : len(rows[i])-1] // removing trailing and leading non-stack symbols
		for j := 0; j < len(row); j += 4 {
			if row[j] != ' ' {
				var stackIndex = j / 4
				stackArray[stackIndex] = append(stackArray[stackIndex], row[j])
			}
		}
	}
	return stackArray
}

func transposeStackArray(stackArray [][]byte) [][]byte {
	maxLength := 0
	for _, stack := range stackArray {
		if len(stack) > maxLength {
			maxLength = len(stack)
		}
	}

	transposedArray := make([][]byte, maxLength)
	for i := range transposedArray {
		transposedArray[i] = make([]byte, len(stackArray))
	}

	for i, stack := range stackArray {
		for j, ch := range stack {
			transposedArray[j][i] = ch
		}
	}

	return transposedArray
}

func printStacksArray(stackArray [][]byte) {
	transposedArray := transposeStackArray(stackArray)

	for i := len(transposedArray) - 1; i >= 0; i-- {
		stack := transposedArray[i]
		asciiStack := make([]string, len(stack))
		for j, b := range stack {
			if b == 0 {
				asciiStack[j] = " 0 "
			} else {
				asciiStack[j] = strconv.QuoteRuneToASCII(rune(b))
			}
		}
		fmt.Println(strings.Join(asciiStack, " "))
	}
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func moveStacksViaInstructionsOneAtTheTime(stackArray [][]byte, instructionRows []string) [][]byte {
	for _, row := range instructionRows {
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := re.FindStringSubmatch(row)
		if len(matches) != 4 {
			log.Fatal("Invalid instruction row: ", row)
		}

		nrToMove := convertStringToInt(matches[1])
		fromStackIndex := convertStringToInt(matches[2]) - 1
		toStackIndex := convertStringToInt(matches[3]) - 1

		fromStack := stackArray[fromStackIndex]
		if len(fromStack) < nrToMove {
			log.Fatal("Trying to move more than there is in stack: ", row)
		}

		toStack := stackArray[toStackIndex]
		for i := 0; i < nrToMove; i++ {
			toStack = append(toStack, fromStack[len(fromStack)-1])
			fromStack = fromStack[:len(fromStack)-1]
		}

		stackArray[fromStackIndex] = fromStack
		stackArray[toStackIndex] = toStack

		// fmt.Println("Moving ", nrToMove, " from ", fromStackIndex, " to ", toStackIndex)
		// fmt.Println("stacksArray state: ")
		// printStacksArray(stackArray)
	}

	return stackArray
}

func day5_part1() {
	rows, err := getRowsFromFile("input5.txt")
	if err != nil {
		log.Fatal(err)
	}

	dividingRow := findEmptyRowIndex(rows)
	stacksArray := getStacksArrayFromRows(rows[0:dividingRow])
	// fmt.Println("stacksArray start state: ")
	// printStacksArray(stacksArray)

	finishedStacksArray := moveStacksViaInstructionsOneAtTheTime(stacksArray, rows[dividingRow+1:len(rows)-1])
	// fmt.Println("stacksArVdray end state: ")
	// printStacksArray(finishedStacksArray)

	solution := ""

	for _, stack := range finishedStacksArray {
		solution += string(stack[len(stack)-1])
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

func moveStacksViaInstructionsMultipleAtTheTime(stackArray [][]byte, instructionRows []string) [][]byte {
	for _, row := range instructionRows {
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := re.FindStringSubmatch(row)
		if len(matches) != 4 {
			log.Fatal("Invalid instruction row: ", row)
		}

		nrToMove := convertStringToInt(matches[1])
		fromStackIndex := convertStringToInt(matches[2]) - 1
		toStackIndex := convertStringToInt(matches[3]) - 1

		fromStack := stackArray[fromStackIndex]
		if len(fromStack) < nrToMove {
			log.Fatal("Trying to move more than there is in stack: ", row)
		}

		toStack := stackArray[toStackIndex]

		toStack = append(toStack, fromStack[len(fromStack)-nrToMove:]...)
		fromStack = fromStack[:len(fromStack)-nrToMove]

		stackArray[fromStackIndex] = fromStack
		stackArray[toStackIndex] = toStack

		// fmt.Println("Moving ", nrToMove, " from ", fromStackIndex, " to ", toStackIndex)
		// fmt.Println("stacksArray state: ")
		// printStacksArray(stackArray)
	}

	return stackArray
}

func day5_part2() {
	rows, err := getRowsFromFile("input5.txt")
	if err != nil {
		log.Fatal(err)
	}

	dividingRow := findEmptyRowIndex(rows)
	stacksArray := getStacksArrayFromRows(rows[0:dividingRow])
	// fmt.Println("stacksArray start state: ")
	// printStacksArray(stacksArray)

	finishedStacksArray := moveStacksViaInstructionsMultipleAtTheTime(stacksArray, rows[dividingRow+1:len(rows)-1])
	// fmt.Println("stacksArVdray end state: ")
	// printStacksArray(finishedStacksArray)

	solution := ""

	for _, stack := range finishedStacksArray {
		solution += string(stack[len(stack)-1])
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}
