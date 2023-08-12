package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func day1_part1() {
	rows, err := getRowsFromFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution := 0
	current_elfs_calories := 0
	for _, row := range rows {
		if row == "" {
			if current_elfs_calories > solution {
				solution = current_elfs_calories
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

	fmt.Println(getFunctionName(), " solution: ", solution)
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

	solution := 0
	for _, value := range elfs_calories[0:3] {
		solution += value
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

