package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1, 5},
		{2, 2, 4, 1},
		{5, 0, 2, 3},
		{0, 6, 1, 2},
	}

	maxVal := goldRun(grid)

	fmt.Println(maxVal)
}

func goldRun(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	maxVal := -1

	for col := 1; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			maxBehind := grid[row][col-1]

			if row > 0 {
				maxBehind = max(maxBehind, grid[row-1][col-1])
			}
			if row < numRows-1 {
				maxBehind = max(maxBehind, grid[row+1][col-1])
			}

			currCellVal := grid[row][col] + maxBehind
			grid[row][col] = currCellVal

			if currCellVal > maxVal {
				maxVal = currCellVal
			}
		}
	}

	return maxVal
}