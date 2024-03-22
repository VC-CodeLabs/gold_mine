package main

import (
	"fmt"
)

func max(nums ...int) int {
    if len(nums) == 0 {
        return 0 // or any appropriate default value
    }
    maxVal := nums[0]
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }
    return maxVal
}

func maxGold(mine [][]int) int {
	rows := len(mine)
	cols := len(mine[0])

	// Create table to store the max gold collected at each position
	goldvein := make([][]int, rows)
	for i := range goldvein {
		goldvein[i] = make([]int, cols)
	}

	// Initialize the first column of table with the gold values from mine
	for i := 0; i < rows; i++ {
		goldvein[i][0] = mine[i][0]
	}

	// loop over each column starting from the second column
	for j := 1; j < cols; j++ {
		for i := 0; i < rows; i++ {

			// move east
			moves := []int{goldvein[i][j-1]}
			if i > 0 {
				// move northeast
				moves = append(moves, goldvein[i-1][j-1])
			}
			if i < rows-1 {
				// move southeast
				moves = append(moves, goldvein[i+1][j-1])
			}
			// Update table with the maximum gold collected at the current position
			goldvein[i][j] = mine[i][j] + max(moves...)
		}
	}

	// Find the maximum gold collected in last column
	maxGoldCollected := 0
	for _, row := range goldvein {
		maxGoldCollected = max(maxGoldCollected, row[cols-1])
	}

	return maxGoldCollected
}

func main() {
	mine1 := [][]int{{1, 3, 3}, {2, 1, 4}, {0, 6, 4}}
	fmt.Println("Expected: 12 =", maxGold(mine1))

	mine2 := [][]int{{1, 3, 1, 5}, {2, 2, 4, 1}, {5, 0, 2, 3}, {0, 6, 1, 2}}
	fmt.Println("Expected: 16 =", maxGold(mine2))

	mine3 := [][]int{{10, 33, 13, 15}, {22, 21, 4, 1}, {5, 0, 2, 3}, {0, 6, 14, 2}}
	fmt.Println("Expected: 83 =", maxGold(mine3))
}
