package main

import (
	"fmt"
	"log"
)

func getTreeMatrixFromString(rows []string) *[][]int {
	treeMatrix := make([][]int, len(rows))
	for idx, treeRow := range rows {
		treeMatrix[idx] = make([]int, len(treeRow))
		for jdx, tree := range treeRow {
			treeMatrix[idx][jdx] = convertStringToInt(string(tree))
		}
	}
	return &treeMatrix
}

func printTreeMatrix(treeMatrix *[][]int) {
	for _, treeRow := range *treeMatrix {
		for _, tree := range treeRow {
			fmt.Printf("%d", tree)
		}
		fmt.Println()
	}
}

func getViewableTreesMatrix(treeMatrix *[][]int) *[][]int {
	viewableTreesMatrix := make([][]int, len(*treeMatrix))

	// look at the row tree height, left -> right; right -> left
	for idx, treeRow := range *treeMatrix {
		viewableTreesMatrix[idx] = make([]int, len(treeRow))
		highestTree := -1
		// from left -> right
		for jdx := 0; jdx < len(treeRow); jdx++ {
			tree := treeRow[jdx]
			if jdx == 0 {
				highestTree = tree
				viewableTreesMatrix[idx][jdx] = 1
			} else {
				if tree > highestTree {
					highestTree = tree
					viewableTreesMatrix[idx][jdx] = 1
				}
			}
		}
		highestTree = -1
		// from right -> left
		for jdx := len(treeRow) - 1; jdx >= 0; jdx-- {
			tree := treeRow[jdx]
			if jdx == len(treeRow)-1 {
				highestTree = tree
				viewableTreesMatrix[idx][jdx] = 1
			} else {
				if tree > highestTree {
					highestTree = tree
					viewableTreesMatrix[idx][jdx] = 1
				}
			}
		}
	}

	// look at the column tree height, top -> bottom; bottom -> top
	idx := 0
	for jdx := 0; jdx < len((*treeMatrix)[idx]); jdx++ {
		highestTree := -1
		// from top -> bottom
		for ; idx < len(*treeMatrix)-1; idx++ {
			tree := (*treeMatrix)[idx][jdx]
			if idx == 0 {
				highestTree = tree
				viewableTreesMatrix[idx][jdx] = 1
			} else {
				if tree > highestTree {
					highestTree = tree
					viewableTreesMatrix[idx][jdx] = 1
				}
			}
		}
		highestTree = -1
		idx--
		// from bottom -> top
		for ; idx >= 0; idx-- {
			tree := (*treeMatrix)[idx][jdx]
			if idx == len(*treeMatrix)-1 {
				highestTree = tree
				viewableTreesMatrix[idx][jdx] = 1
			} else {
				if tree > highestTree {
					highestTree = tree
					viewableTreesMatrix[idx][jdx] = 1
				}
			}
		}
		idx = 0
	}

	return &viewableTreesMatrix
}

func countNrOfVisibleTrees(viewableTreesMatrix *[][]int) int {
	nrOfVisibleTrees := 0
	for _, viewableTreesRow := range *viewableTreesMatrix {
		for _, viewableTree := range viewableTreesRow {
			nrOfVisibleTrees += viewableTree
		}
	}
	return nrOfVisibleTrees
}

func day8_part1() {
	rows, err := getRowsFromFile("input8.txt")

	if err != nil {
		log.Fatal(err)
	}

	treeMatrix := getTreeMatrixFromString(rows)
	// printTreeMatrix(treeMatrix)
	viewableTreesMatrix := getViewableTreesMatrix(treeMatrix)
	// printTreeMatrix(viewableTreesMatrix)
	solution := countNrOfVisibleTrees(viewableTreesMatrix)
	fmt.Println(getFunctionName(), " solution: ", solution)
}

func getScenicScoreForePosition(treeMatrix *[][]int, x int, y int) int {
	currentTreeHeight := (*treeMatrix)[x][y]
	rightScore := 0
	for jdx := y + 1; jdx < len((*treeMatrix)[x]); jdx++ {
		viewedTreeHeight := (*treeMatrix)[x][jdx]
		if viewedTreeHeight < currentTreeHeight {
			rightScore++
		} else {
			rightScore++
			break
		}
	}

	leftScore := 0
	for jdx := y - 1; jdx >= 0; jdx-- {
		viewedTreeHeight := (*treeMatrix)[x][jdx]
		if viewedTreeHeight < currentTreeHeight {
			leftScore++
		} else {
			leftScore++
			break
		}
	}

	topScore := 0
	for idx := x - 1; idx >= 0; idx-- {
		viewedTreeHeight := (*treeMatrix)[idx][y]
		if viewedTreeHeight < currentTreeHeight {
			topScore++
		} else {
			topScore++
			break
		}
	}

	bottomScore := 0
	for idx := x + 1; idx < len(*treeMatrix)-1; idx++ {
		viewedTreeHeight := (*treeMatrix)[idx][y]
		if viewedTreeHeight < currentTreeHeight {
			bottomScore++
		} else {
			bottomScore++
			break
		}
	}

	return leftScore * rightScore * topScore * bottomScore
}

// To measure the viewing distance from a given tree, look up, down, left, and right from that tree;
// stop if you reach an edge or at the first tree that is the same height or taller than the tree under consideration.
// (If a tree is right on the edge, at least one of its viewing distances will be zero.)
func getScenicScoreMatrix(treeMatrix *[][]int) *[][]int {
	scenicScoreMatrix := make([][]int, len(*treeMatrix))
	for idx := range scenicScoreMatrix {
		scenicScoreMatrix[idx] = make([]int, len((*treeMatrix)[idx]))
	}

	for idx, treeRow := range *treeMatrix {
		for jdx := range treeRow {
			if jdx != 0 && jdx != len((*treeMatrix)[idx])-1 && idx != 0 && idx != len(*treeMatrix)-1 {
				scenicScoreMatrix[idx][jdx] = getScenicScoreForePosition(treeMatrix, idx, jdx)
			} else {
				scenicScoreMatrix[idx][jdx] = 0
			}
		}
	}
	return &scenicScoreMatrix
}

func findHighestScenicScore(scenicScoreMatrix *[][]int) int {
	highestScenicScore := -1
	for _, scenicScoreRow := range *scenicScoreMatrix {
		for _, scenicScore := range scenicScoreRow {
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return highestScenicScore
}

func day8_part2() {
	rows, err := getRowsFromFile("input8.txt")

	if err != nil {
		log.Fatal(err)
	}

	treeMatrix := getTreeMatrixFromString(rows)
	// printTreeMatrix(treeMatrix)
	viewableTreesMatrix := getScenicScoreMatrix(treeMatrix)
	// printTreeMatrix(viewableTreesMatrix)
	solution := findHighestScenicScore(viewableTreesMatrix)
	fmt.Println(getFunctionName(), " solution: ", solution)
}
