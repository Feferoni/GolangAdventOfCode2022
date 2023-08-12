
package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

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
