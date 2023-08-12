
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Ranges struct {
	min uint
	max uint
}

func (lhs *Ranges) isSubRangeOf(rhs *Ranges) bool {
	return rhs.min <= lhs.min && lhs.max <= rhs.max
}

func (lhs *Ranges) isIntersecting(rhs *Ranges) bool {
	return lhs.min <= rhs.min && rhs.min <= lhs.max ||
		lhs.min <= rhs.max && rhs.max <= lhs.max ||
		rhs.min <= lhs.min && lhs.min <= rhs.max ||
		rhs.min <= lhs.max && lhs.max <= rhs.max
}

func getRangeFromString(rangeString string) *Ranges {
	ranges := strings.Split(rangeString, "-")
	min, _ := strconv.ParseUint(ranges[0], 10, 32)
	max, _ := strconv.ParseUint(ranges[1], 10, 32)
	return &Ranges{uint(min), uint(max)}
}

func day4_part1() {
	rows, err := getRowsFromFile("input4.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution := 0

	for _, row := range rows[0 : len(rows)-1] {
		ranges := strings.Split(row, ",")
		firstRange := getRangeFromString(ranges[0])
		secondRange := getRangeFromString(ranges[1])

		if firstRange.isSubRangeOf(secondRange) || secondRange.isSubRangeOf(firstRange) {
			solution += 1
		}
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

func day4_part2() {
	rows, err := getRowsFromFile("input4.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution := 0

	for _, row := range rows[0 : len(rows)-1] {
		ranges := strings.Split(row, ",")
		firstRange := getRangeFromString(ranges[0])
		secondRange := getRangeFromString(ranges[1])

		if firstRange.isIntersecting(secondRange) {
			solution += 1
		}
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}
