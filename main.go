package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	stdTime "time"
)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

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

func main() {
	time := stdTime.Now()
	day1_part1()
	day1_part2()
	day2_part1()
	day2_part2()
	day3_part1()
	day3_part2()
	day4_part1()
	day4_part2()
	day5_part1()
	day5_part2()
	day6_part1()
	day6_part2()
	day7_part1()
	day7_part2()
	day8_part1()
	day8_part2()
	day9_part1()
    day9_part2()
	day10_part1()
	day10_part2()
	day11_part1()
	day11_part2()

	duration := stdTime.Since(time)
	fmt.Println("Duration: ", duration)
}
