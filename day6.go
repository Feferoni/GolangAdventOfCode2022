
package main

import (
	"fmt"
	"log"
)

func getIndexOfFirstUniqueSequence(nrOfUnique int, sequence string) int {
	i := nrOfUnique - 1
	for {
		if i >= len(sequence) {
			break
		}
		minIndex := i - nrOfUnique + 1
		charBits := 1 << (uint32(sequence[i]) % 32)
		for j := i - 1; j >= minIndex; j-- {
			currCharBit := 1 << (uint32(sequence[j]) % 32)
			if charBits&currCharBit != 0 {
				i = j + nrOfUnique
				break
			}

			charBits = charBits | currCharBit

			if j == minIndex {
				return i + 1
			}
		}
	}

	return -1
}

func day6_part1() {
	rows, err := getRowsFromFile("input6.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows[:len(rows)-1] {
		solution := getIndexOfFirstUniqueSequence(4, row)

		fmt.Println(getFunctionName(), " solution: ", solution)
	}
}

func day6_part2() {
	rows, err := getRowsFromFile("input6.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows[:len(rows)-1] {
		solution := getIndexOfFirstUniqueSequence(14, row)

		fmt.Println(getFunctionName(), " solution: ", solution)
	}
}
