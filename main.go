package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

	solution := 0
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
		solution += currentRoundPoints

		// fmt.Println("Opponent: ", opponent, " You: ", you, " Winner: ", currentWinnerPoints, " Hand: ", currentHandPoints, " Round: ", currentRoundPoints)
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
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

	solution := 0
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
		solution += currentRoundPoints
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

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

		fmt.Println("Moving ", nrToMove, " from ", fromStackIndex, " to ", toStackIndex)
		fmt.Println("stacksArray state: ")
		printStacksArray(stackArray)
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
	fmt.Println("stacksArray start state: ")
	printStacksArray(stacksArray)

	finishedStacksArray := moveStacksViaInstructionsOneAtTheTime(stacksArray, rows[dividingRow+1:len(rows)-1])
	fmt.Println("stacksArVdray end state: ")
	printStacksArray(finishedStacksArray)

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

		fmt.Println("Moving ", nrToMove, " from ", fromStackIndex, " to ", toStackIndex)
		fmt.Println("stacksArray state: ")
		printStacksArray(stackArray)
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
	fmt.Println("stacksArray start state: ")
	printStacksArray(stacksArray)

	finishedStacksArray := moveStacksViaInstructionsMultipleAtTheTime(stacksArray, rows[dividingRow+1:len(rows)-1])
	fmt.Println("stacksArVdray end state: ")
	printStacksArray(finishedStacksArray)

	solution := ""

	for _, stack := range finishedStacksArray {
		solution += string(stack[len(stack)-1])
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

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
		fmt.Println("bufferString: ", row)
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
		fmt.Println("bufferString: ", row)
		solution := getIndexOfFirstUniqueSequence(14, row)

		fmt.Println(getFunctionName(), " solution: ", solution)
	}
}

type File struct {
	name string
	size int
}

type Directory struct {
	name      string
	parent    *Directory
	children  []*Directory
	files     []File
	totalSize int
}

func cdCommand(rootDirectory *Directory, currentDirectory *Directory, cmd []string) (*Directory, error) {
	if cmd[2] == ".." {
		return currentDirectory.parent, nil
	} else if cmd[2] == "/" {
		return rootDirectory, nil
	} else {
		for _, child := range currentDirectory.children {
			if child.name == cmd[2] {
				return child, nil
			}
		}
	}
	return nil, errors.New("Directory not found")
}

func doesDirectoryExist(directory *Directory, name string) bool {
	for _, child := range directory.children {
		if child.name == name {
			return true
		}
	}
	return false
}

func parseDirectoryFromStrings(fileLines []string) *Directory {
	root := Directory{name: "/", parent: nil, children: []*Directory{}, files: []File{}}
	currentDirectory := &root
	for _, row := range fileLines {
		if row == "" {
			continue
		}

		parsedRow := strings.Split(row, " ")
		if parsedRow[0] == "$" { // checks for command
			if parsedRow[1] == "cd" {
				var err error
				currentDirectory, err = cdCommand(&root, currentDirectory, parsedRow)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if parsedRow[0] == "dir" {
				if !doesDirectoryExist(currentDirectory, parsedRow[1]) {
					currentDirectory.children = append(currentDirectory.children, &Directory{name: parsedRow[1], parent: currentDirectory, children: []*Directory{}, files: []File{}})
				} else {
					fmt.Println("Directory ", parsedRow[1], " already exists")
				}
			} else {
				currentDirectory.files = append(currentDirectory.files, File{name: parsedRow[1], size: getFileSizeFromString(parsedRow[0])})
			}
		}
	}

	return &root
}

func getFileSizeFromString(fileString string) int {
	size, err := strconv.Atoi(fileString)
	if err != nil {
		log.Fatal(err)
	}
	return size
}

func hasVisitedDirectory(current *Directory, visited []*Directory) bool {
	for _, directory := range visited {
		if directory.name == current.name && directory.parent.name == current.parent.name {
			return true
		}
	}
	return false
}

func isAllChildrenVisited(directory *Directory, visited []*Directory) bool {
	for _, child := range directory.children {
		if !hasVisitedDirectory(child, visited) {
			return false
		}
	}
	return true
}

func getFileSizeOfDirectory(directory *Directory) int {
	size := 0
	for _, file := range directory.files {
		size += file.size
	}
	return size
}

func setTotalSizeToAllDirectories(directory *Directory) int {
	size := 0
	for _, child := range directory.children {
		size += setTotalSizeToAllDirectories(child)
	}
	directory.totalSize = size + getFileSizeOfDirectory(directory)
	return directory.totalSize
}

func printDirectoryTree(directory *Directory) {
	visitedDirectories := []*Directory{}
	depth := 0
	depthToPrint := 0
	printDirectoryRecursive(directory, visitedDirectories, depth, depthToPrint)
}

func printDirectoryRecursive(directory *Directory, visitedDirectories []*Directory, depth int, depthToPrint int) {
	indentation := strings.Repeat("  ", depth)

	if depthToPrint == 0 || depth <= depthToPrint {
		fmt.Printf("%sdir: %s size: [%d] depth: %d\n", indentation, directory.name, directory.totalSize, depth)
		visitedDirectories = append(visitedDirectories, directory)
	}
	for _, child := range directory.children {
		if !hasVisitedDirectory(child, visitedDirectories) {
			printDirectoryRecursive(child, visitedDirectories, depth+1, depthToPrint)
		}
	}
}

func day7_part1() {
	rows, err := getRowsFromFile("input7.txt")
	if err != nil {
		log.Fatal(err)
	}

	root := parseDirectoryFromStrings(rows)
	setTotalSizeToAllDirectories(root)

	maxSize := 100000
	solution := 0

	visitedDirectories := []*Directory{}
	currentDirectory := root
	for {
		if len(currentDirectory.children) == 0 {
			if currentDirectory.totalSize <= maxSize {
				solution += currentDirectory.totalSize
			}
			visitedDirectories = append(visitedDirectories, currentDirectory)
			if currentDirectory.parent == nil {
				break
			}
			currentDirectory = currentDirectory.parent
		} else {
			if isAllChildrenVisited(currentDirectory, visitedDirectories) {
				if currentDirectory.totalSize <= maxSize {
					solution += currentDirectory.totalSize
				}
				visitedDirectories = append(visitedDirectories, currentDirectory)
				if currentDirectory.parent == nil {
					break
				}
				currentDirectory = currentDirectory.parent
			} else {
				for _, child := range currentDirectory.children {
					if !hasVisitedDirectory(child, visitedDirectories) {
						currentDirectory = child
						break
					}
				}
			}
		}
	}

	fmt.Println(getFunctionName(), " solution: ", solution)
}

func day7_part2() {
	rows, err := getRowsFromFile("input7.txt")

	if err != nil {
		log.Fatal(err)
	}

	root := parseDirectoryFromStrings(rows)
	setTotalSizeToAllDirectories(root)

	fileSystemMaxSize := 70000000
	neededSpace := 30000000
	usedSpace := root.totalSize
	totSize := neededSpace + usedSpace

	needToFree := totSize - fileSystemMaxSize


	deleteCandidate := root
	visitedDirectories := []*Directory{}
	currentDirectory := root
	for {
		if len(currentDirectory.children) == 0 {
			if currentDirectory.totalSize >= needToFree {
				if currentDirectory.totalSize <= deleteCandidate.totalSize {
					deleteCandidate = currentDirectory
				}
			}
			visitedDirectories = append(visitedDirectories, currentDirectory)
			if currentDirectory.parent == nil {
				break
			}
			currentDirectory = currentDirectory.parent
		} else {
			if isAllChildrenVisited(currentDirectory, visitedDirectories) {
				if currentDirectory.totalSize >= needToFree {
					if currentDirectory.totalSize <= deleteCandidate.totalSize {
						deleteCandidate = currentDirectory
					}
				}
				visitedDirectories = append(visitedDirectories, currentDirectory)
				if currentDirectory.parent == nil {
					break
				}
				currentDirectory = currentDirectory.parent
			} else {
				for _, child := range currentDirectory.children {
					if !hasVisitedDirectory(child, visitedDirectories) {
						currentDirectory = child
						break
					}
				}
			}
		}
	}

	solution := deleteCandidate.totalSize
	fmt.Println(getFunctionName(), " solution: ", solution)
}

func main() {
	// day1_part1()
	// day1_part2()
	// day2_part1()
	// day2_part2()
	// day3_part1()
	// day3_part2()
	// day4_part1()
	// day4_part2()
	// day5_part1()
	// day5_part2()
	// day6_part1()
	// day6_part2()
	// day7_part1()
	day7_part2()
}
