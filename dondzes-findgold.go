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

	// Create a dp table to store the maximum gold collected at each position
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// Initialize the first column of dp table with the gold values from mine
	for i := 0; i < rows; i++ {
		dp[i][0] = mine[i][0]
	}

	// Iterate over each column starting from the second column
	for j := 1; j < cols; j++ {
		for i := 0; i < rows; i++ {
			// Consider three possible moves: right, right-up diagonal, right-down diagonal
			moves := []int{dp[i][j-1]}
			if i > 0 {
				moves = append(moves, dp[i-1][j-1])
			}
			if i < rows-1 {
				moves = append(moves, dp[i+1][j-1])
			}
			// Update dp table with the maximum gold collected at the current position
			dp[i][j] = mine[i][j] + max(moves...)
		}
	}

	// Find the maximum gold collected in the last column
	maxGoldCollected := 0
	for _, row := range dp {
		maxGoldCollected = max(maxGoldCollected, row[cols-1])
	}

	return maxGoldCollected
}

func main() {
	//mine1 := [][]int{{1, 3, 3}, {2, 1, 4}, {0, 6, 4}}
	//fmt.Println(maxGold(mine1)) // Output: 12

	//mine2 := [][]int{{1, 3, 1, 5}, {2, 2, 4, 1}, {5, 0, 2, 3}, {0, 6, 1, 2}}
	//fmt.Println(maxGold(mine2)) // Output: 16

	//mine3 := [][]int{{10, 33, 13, 15}, {22, 21, 4, 1}, {5, 0, 2, 3}, {0, 6, 14, 2}}
	//fmt.Println(maxGold(mine3)) // Output: 83

	//mine4 := [][]int{{1}, {6}, {3}}
	//fmt.Println("Expected: 6 =", maxGold(mine4))

	//mine5 := [][]int{{6, 6, 6}}
	//fmt.Println("Expected: 18 =", maxGold(mine5))

	//mine6 := [][]int {{9872,0,0,0,0},{0,0,0,0,0}, {0,0,0,0,0}, {0,0,0,9872,9872}}
	//fmt.Println("Expected: 29616 =", maxGold(mine6))

	mine7 := [][]int {{9872,0,0,0,0},{0,0,0,0,0}, {0,0,0,0,0}, {0,0,0,0,0}, {0,0,0,9872,9872}}
	fmt.Println("Expected: 19744 =", maxGold(mine7))
	
    	//array := generate2DArray(1000, 1000, 9872)
	//fmt.Println("Expected: ALOT =", maxGold(array))
}
