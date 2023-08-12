
package main

import (
	"fmt"
	"log"
	"unicode"
)

func getSymbolsBits(row string) uint64 {
	var symbolsBit uint64 = 0
	for _, value := range row {
		var currentSymbolsBit uint64 = 0
		currentSymbolsBit = 1 << (uint64(value) % 64)
		symbolsBit = symbolsBit | currentSymbolsBit
	}
	return symbolsBit
}

func day3_part1() {
	rows, err := getRowsFromFile("input3.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution := 0
	for _, row := range rows {
		var middle int = len(row) / 2
		firstHalfOfRow := row[0:middle]
		secondHalfOfRow := row[middle:]

		var firstHalvesSymbols uint64 = getSymbolsBits(firstHalfOfRow)
		var secondHalvesSymbols uint64 = 0

		for _, value := range secondHalfOfRow {
			var currentSymbolsBit uint64 = 0
			currentSymbolsBit = 1 << (uint64(value) % 64)
			secondHalvesSymbols = secondHalvesSymbols | currentSymbolsBit
			if firstHalvesSymbols&currentSymbolsBit != 0 {
				var currentValue int = 0
				if unicode.IsLower(value) {
					currentValue = int(value) - 'a' + 1
				} else {
					currentValue = int(value) - 'A' + 27
				}
				solution += currentValue
				break
			}
		}
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

func day3_part2() {
	rows, err := getRowsFromFile("input3.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution := 0

	for i := 2; i < len(rows); i += 3 {
		firstElf := rows[i-2]
		secondElf := rows[i-1]
		thirdElf := rows[i]

		var firstAndSecondSymbolsBit uint64 = getSymbolsBits(firstElf) & getSymbolsBits(secondElf)

		for _, value := range thirdElf {
			var currentSymbolsBit uint64 = 0
			currentSymbolsBit = 1 << (uint64(value) % 64)
			if firstAndSecondSymbolsBit&currentSymbolsBit != 0 {
				var currentValue int = 0
				if unicode.IsLower(value) {
					currentValue = int(value) - 'a' + 1
				} else {
					currentValue = int(value) - 'A' + 27
				}
				solution += currentValue
				break
			}
		}
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}
